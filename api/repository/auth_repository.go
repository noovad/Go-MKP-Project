package repository

import (
	"go-gin-project/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindByUsername(username string) (model.User, error)
	FindById(id uuid.UUID) (model.User, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{
		db: db,
	}
}

func (r *authRepositoryImpl) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *authRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *authRepositoryImpl) FindById(id uuid.UUID) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}
