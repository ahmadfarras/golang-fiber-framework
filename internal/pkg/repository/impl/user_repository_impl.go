package repository

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model"
	"ahmadfarras/fiberframework/internal/pkg/repository"

	"gorm.io/gorm"
)

type UserRepositoryGormImpl struct {
	db *gorm.DB
}

func NewUserRepositoryGormImpl(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryGormImpl{
		db: db,
	}
}

func (u *UserRepositoryGormImpl) Create(user model.User) error {
	u.db.Create(user)

	return nil
}
