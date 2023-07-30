package usecase

import (
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/advengulo/go-clean-arch-test/internal/modules/user/repository"
	"github.com/advengulo/go-clean-arch-test/internal/utils"
	"net/http"
)

// UserUseCase is the interface for the user service
type UserUseCase interface {
	GetAllUser() domains.Response
	GetUser(id uint) domains.Response
	Create(pl *domains.User) domains.Response
	Delete(id uint) domains.Response
	GetByUsername(username string) domains.Response
}

// UserUseCaseImpl implements the UserUseCase interface
type UserUseCaseImpl struct {
	userRepository repository.UserRepository
}

// NewUserUseCase creates a new UserUseCaseImpl instance
func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{userRepository: repository}
}

func (s *UserUseCaseImpl) GetAllUser() domains.Response {
	data, err := s.userRepository.GetAllUser()
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusInternalServerError)
	}

	return utils.Response("OK", data, nil, http.StatusOK)
}

// GetUser returns a user with the given ID
func (s *UserUseCaseImpl) GetUser(id uint) domains.Response {
	data, err := s.userRepository.GetByID(id)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusNotFound)
	}

	return utils.Response("OK", data, nil, http.StatusOK)
}

func (s *UserUseCaseImpl) Create(pl *domains.User) domains.Response {
	hashPassword, err := utils.HashPassword(pl.Password)
	pl.Password = hashPassword

	err = s.userRepository.Create(pl)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusInternalServerError)
	}

	return utils.Response("OK", pl, nil, http.StatusOK)
}

func (s *UserUseCaseImpl) Delete(id uint) domains.Response {
	user, err := s.userRepository.GetByID(id)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusNotFound)
	}

	err = s.userRepository.Delete(user)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusInternalServerError)
	}

	return utils.Response("OK", user, nil, http.StatusOK)

}

// GetByUsername returns a user with the given username
func (s *UserUseCaseImpl) GetByUsername(username string) domains.Response {
	data, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusNotFound)
	}

	return utils.Response("OK", data, nil, http.StatusOK)
}
