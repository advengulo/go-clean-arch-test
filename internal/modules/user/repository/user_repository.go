package repository

import (
	"fmt"
	"github.com/advengulo/go-clean-arch-test/domains"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]*domains.User, error)
	Create(user *domains.User) error
	GetByID(id uint) (*domains.User, error)
	GetByUsername(username string) (*domains.User, error)
	Update(user *domains.User) error
	Delete(user *domains.User) error
}

// UserRepositoryImpl implements the UserRepository interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepositoryImpl instance
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetAllUser() (users []*domains.User, err error) {
	if err = r.db.Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("users data is not exist")
		}
		return nil, fmt.Errorf("failed to get users data %w", err)
	}

	return
}

func (r *UserRepositoryImpl) Create(user *domains.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepositoryImpl) GetByID(id uint) (*domains.User, error) {
	user := &domains.User{}
	if err := r.db.First(user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found with id %d", id)
		}
		return nil, fmt.Errorf("failed to get user with id %d: %w", id, err)
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetByUsername(username string) (*domains.User, error) {
	user := &domains.User{}
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("failed to get user")
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(user *domains.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(user *domains.User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
