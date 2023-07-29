package usecase

import (
	"github.com/advengulo/go-clean-arch-test/internal/models"
	"github.com/advengulo/go-clean-arch-test/internal/modules/user/repository"
	"github.com/advengulo/go-clean-arch-test/pkg/utils"
	"net/http"
)

// UserUseCase is the interface for the user service
type UserUseCase interface {
	GetAllUser() models.Response
	GetUser(id uint) models.Response
	Create(pl *models.User) models.Response
	Delete(id uint) models.Response
	GetByUsername(username string) models.Response
}

// UserUseCaseImpl implements the UserUseCase interface
type UserUseCaseImpl struct {
	UserRepository repository.UserRepository
}

// NewUserUseCase creates a new UserUseCaseImpl instance
func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{UserRepository: repository}
}

func (s *UserUseCaseImpl) GetAllUser() models.Response {
	data, err := s.UserRepository.GetAllUser()
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusInternalServerError)
	}

	return utils.Response("OK", data, nil, http.StatusOK)
}

// GetUser returns a user with the given ID
func (s *UserUseCaseImpl) GetUser(id uint) models.Response {
	data, err := s.UserRepository.GetByID(id)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusNotFound)
	}

	return utils.Response("OK", data, nil, http.StatusOK)
}

func (s *UserUseCaseImpl) Create(pl *models.User) models.Response {
	hashPassword, err := utils.HashPassword(pl.Password)
	pl.Password = hashPassword

	err = s.UserRepository.Create(pl)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusInternalServerError)
	}

	return utils.Response("OK", pl, nil, http.StatusOK)
}

func (s *UserUseCaseImpl) Delete(id uint) models.Response {
	user, err := s.UserRepository.GetByID(id)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusNotFound)
	}

	err = s.UserRepository.Delete(user)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusInternalServerError)
	}

	return utils.Response("OK", user, nil, http.StatusOK)

}

// GetByUsername returns a user with the given username
func (s *UserUseCaseImpl) GetByUsername(username string) models.Response {
	data, err := s.UserRepository.GetByUsername(username)
	if err != nil {
		return utils.Response("ERROR", nil, err.Error(), http.StatusNotFound)
	}

	return utils.Response("OK", data, nil, http.StatusOK)
}
