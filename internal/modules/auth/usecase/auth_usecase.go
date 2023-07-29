package usecase

import (
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/advengulo/go-clean-arch-test/internal/modules/user/usecase"
	utils2 "github.com/advengulo/go-clean-arch-test/internal/utils"
	"net/http"
)

type AuthUseCase interface {
	Login(pl *domains.UserPayload) domains.Response
	Validate(token string) domains.Response
}

type auth struct {
	ucUser usecase.UserUseCase
}

func NewAuthUseCase(ucUser usecase.UserUseCase) AuthUseCase {
	return &auth{ucUser: ucUser}
}

func (a *auth) Login(pl *domains.UserPayload) domains.Response {
	user := a.ucUser.GetByUsername(pl.Username)
	if user.Error != nil {
		return utils2.Response("Error", nil, "Username or password invalid", http.StatusUnauthorized)
	}

	userData := user.Data.(*domains.User)

	if !utils2.CheckPasswordHash(pl.Password, userData.Password) {
		return utils2.Response("Error", nil, "Username or password invalid", http.StatusUnauthorized)
	}

	token, err := utils2.CreateToken(*pl)
	if err != nil {
		return utils2.Response("Error", nil, "Something went wrong", http.StatusInternalServerError)

	}

	return utils2.Response("OK", token, nil, http.StatusOK)
}

func (a *auth) Validate(token string) domains.Response {
	dataToken, err := utils2.GetDataToken(token)
	if err != nil {
		return utils2.Response("Error", nil, err, http.StatusUnauthorized)
	}

	return utils2.Response("OK", dataToken, nil, http.StatusOK)
}
