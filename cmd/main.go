package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
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
	Name     string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Email    string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" form:"password" query:"password" validate:"required,min=6"`
}

func main() {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	e.POST("/users", func(c *echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		if err := c.Validate(u); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, u)
	})

	if err := e.Start(":5000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
