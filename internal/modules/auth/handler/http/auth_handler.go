package handler

import (
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/advengulo/go-clean-arch-test/internal/modules/auth/usecase"
	"github.com/advengulo/go-clean-arch-test/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	authUC    usecase.AuthUseCase
	validator *validator.Validate
}

func NewAuthHandler(e *echo.Echo, v *validator.Validate, ucAuth usecase.AuthUseCase) {
	h := &AuthHandler{
		authUC:    ucAuth,
		validator: v,
	}

	e.POST("/auth/login", h.Login)
	e.GET("/auth/validation", h.Validate)
}

func (a *AuthHandler) Login(c echo.Context) error {
	var payload *domains.UserPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := a.validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.Response("Error Validation", nil, utils.ErrorValidation(errors), http.StatusBadRequest))
	}

	resp := a.authUC.Login(payload)

	return c.JSON(resp.HttpCode(), resp)
}

func (a *AuthHandler) Validate(c echo.Context) error {
	token := utils.GetHeaderToken(c)

	resp := a.authUC.Validate(token)

	return c.JSON(resp.HttpCode(), resp)
}
