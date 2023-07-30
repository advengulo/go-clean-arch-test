package handler

import (
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/advengulo/go-clean-arch-test/internal/modules/user/usecase"
	"github.com/advengulo/go-clean-arch-test/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	userUC    usecase.UserUseCase
	validator *validator.Validate
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(e *echo.Echo, v *validator.Validate, ucUser usecase.UserUseCase) {
	h := &UserHandler{
		userUC:    ucUser,
		validator: v,
	}

	e.GET("/users", h.GetAllUser)
	e.GET("/users/:id", h.GetUser)
	e.POST("/users", h.Create)
	e.DELETE("users/:id", h.Delete)
}

// GetAllUser returns all users
func (h *UserHandler) GetAllUser(c echo.Context) error {
	resp := h.userUC.GetAllUser()

	return c.JSON(resp.HttpCode(), resp)
}

// GetUser returns a user with the given ID
func (h *UserHandler) GetUser(c echo.Context) error {
	// Get the ID from the URL parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Handle the error
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	// Call the user service to get the user
	resp := h.userUC.GetUser(uint(id))

	return c.JSON(resp.HttpCode(), resp)
}

func (h *UserHandler) Create(c echo.Context) error {
	var payload domains.User

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := h.validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.Response("Error Validation", nil, utils.ErrorValidation(errors, payload), http.StatusBadRequest))
	}

	resp := h.userUC.Create(&payload)

	return c.JSON(resp.HttpCode(), resp)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Handle the error
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}

	// Call the user service to get the user
	resp := h.userUC.Delete(uint(id))

	return c.JSON(resp.HttpCode(), resp)
}
