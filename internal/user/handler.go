package user

import (
	"net/http"

	"github.com/SifatU360/GoTix-.git/internal/httpresponse"
	"github.com/SifatU360/GoTix-.git/internal/user/dto"
	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateUser(c *echo.Context) error {
	var user dto.CreateRequest  // user input

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid Request Payload",
			Details: err.Error(),
		})
	}


	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Validation Failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.CreateUser(user)
	if err !=nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Failed to Create User",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response)
}