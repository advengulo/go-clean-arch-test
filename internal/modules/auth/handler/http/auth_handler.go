package handler

import (
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/advengulo/go-clean-arch-test/internal/modules/auth/usecase"
	"github.com/advengulo/go-clean-arch-test/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	ucAuth    usecase.AuthUseCase
	Validator *validator.Validate
}

func NewAuthHandler(e *echo.Echo, v *validator.Validate, ucAuth usecase.AuthUseCase) {
	h := &AuthHandler{
		ucAuth:    ucAuth,
		Validator: v,
	}

	e.POST("/auth/login", h.Login)
	e.GET("/auth/validation", h.Validate)
}

func (a *AuthHandler) Login(c echo.Context) error {
	var payload *domains.UserPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := a.Validator.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.Response("Error Validation", nil, utils.ErrorValidation(errors), http.StatusBadRequest))
	}

	resp := a.ucAuth.Login(payload)

	return c.JSON(resp.HttpCode(), resp)
}

func (a *AuthHandler) Validate(c echo.Context) error {
	token := utils.GetHeaderToken(c)

	resp := a.ucAuth.Validate(token)

	return c.JSON(resp.HttpCode(), resp)
}
