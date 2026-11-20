package main

import (
	"fmt"
	"net/http"

	"github.com/SifatU360/GoTix-.git/internal/config"
	"github.com/SifatU360/GoTix-.git/internal/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
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

	cfg := config.LoadEnv()

	db := config.ConnectDatabase(cfg)
	db.AutoMigrate(&User{})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	// e.POST("/users", func(c *echo.Context) error {
	// 	u := new(User)

	// 	// Binding the user data
	// 	if err := c.Bind(u); err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	// 	}

	// 	// Validating the user data
	// 	if err := c.Validate(u); err != nil {
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	// 	}

	// 	// save to database
	// 	if err := db.Create(u).Error; err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	// 	}

	// 	return c.JSON(http.StatusCreated, u)
	// })

	user.RegisterRoutes(e, db)

	port := fmt.Sprintf(":%s", cfg.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
