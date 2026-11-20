package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally return the error to let each route control the status code.
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

type User struct {
	gorm.Model
	Name     string `json:"name" xml:"name" form:"name" query:"name" validate:"required" gorm:"type:varchar(100); not null"`
	Email    string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email" gorm:"type:varchar(255); uniqueIndex; not null"`
	Password string `json:"password" xml:"password" form:"password" query:"password" validate:"required,min=6" gorm:"type:varchar(100); not null"`
}

func main() {

	dsn := "host=localhost user=postgres password=1234 dbname=gotix port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to connect database")
	} else {
		println("Database connected successfully")
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	e.POST("/users", func(c *echo.Context) error {
		u := new(User)

		// Binding the user data
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		// Validating the user data
		if err := c.Validate(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		// save to database
		if err := db.Create(u).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, u)
	})

	if err := e.Start(":5000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
