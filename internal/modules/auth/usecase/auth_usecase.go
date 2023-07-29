package usecase

import (
	"github.com/advengulo/go-clean-arch-test/internal/models"
	"github.com/advengulo/go-clean-arch-test/internal/modules/user/usecase"
	"github.com/advengulo/go-clean-arch-test/pkg/utils"
	"net/http"
)

type AuthUseCase interface {
	Login(pl *models.UserPayload) models.Response
	Validate(token string) models.Response
}

type auth struct {
	ucUser usecase.UserUseCase
}

func NewAuthUseCase(ucUser usecase.UserUseCase) AuthUseCase {
	return &auth{ucUser: ucUser}
}

func (a *auth) Login(pl *models.UserPayload) models.Response {
	user := a.ucUser.GetByUsername(pl.Username)
	if user.Error != nil {
		return utils.Response("Error", nil, "Username or password invalid", http.StatusUnauthorized)
	}

	userData := user.Data.(*models.User)

	if !utils.CheckPasswordHash(pl.Password, userData.Password) {
		return utils.Response("Error", nil, "Username or password invalid", http.StatusUnauthorized)
	}

	token, err := utils.CreateToken(*pl)
	if err != nil {
		return utils.Response("Error", nil, "Something went wrong", http.StatusInternalServerError)

	}

	return utils.Response("OK", token, nil, http.StatusOK)
}

func (a *auth) Validate(token string) models.Response {
	dataToken, err := utils.GetDataToken(token)
	if err != nil {
		return utils.Response("Error", nil, err, http.StatusUnauthorized)
	}

	return utils.Response("OK", dataToken, nil, http.StatusOK)
}
