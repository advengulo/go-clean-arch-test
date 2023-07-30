package usecase

import (
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/advengulo/go-clean-arch-test/internal/modules/user/usecase"
	"github.com/advengulo/go-clean-arch-test/internal/utils"
	"net/http"
)

type AuthUseCase interface {
	Login(pl *domains.UserPayload) domains.Response
	Validate(token string) domains.Response
}

type auth struct {
	userUC usecase.UserUseCase
}

func NewAuthUseCase(ucUser usecase.UserUseCase) AuthUseCase {
	return &auth{userUC: ucUser}
}

func (a *auth) Login(pl *domains.UserPayload) domains.Response {
	user := a.userUC.GetByUsername(pl.Username)
	if user.Error != nil {
		return utils.Response(nil, "Username or password invalid", http.StatusUnauthorized)
	}

	userData := user.Data.(*domains.User)

	if !utils.CheckPasswordHash(pl.Password, userData.Password) {
		return utils.Response(nil, "Username or password invalid", http.StatusUnauthorized)
	}

	token, err := utils.CreateToken(*pl)
	if err != nil {
		return utils.Response(nil, "Something went wrong", http.StatusInternalServerError)

	}

	return utils.Response(token, nil, http.StatusOK)
}

func (a *auth) Validate(token string) domains.Response {
	dataToken, err := utils.GetDataToken(token)
	if err != nil {
		return utils.Response(nil, err, http.StatusUnauthorized)
	}

	return utils.Response(dataToken, nil, http.StatusOK)
}
