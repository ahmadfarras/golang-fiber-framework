package repository

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model"
	"ahmadfarras/fiberframework/internal/pkg/repository"

	"github.com/google/uuid"
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

func (u *UserRepositoryGormImpl) GetAll(users *[]model.User) error {
	u.db.Find(users)

	return nil
}

func (u *UserRepositoryGormImpl) GetById(id uuid.UUID) (model.User, error) {
	user := model.User{ID: id}
	d := u.db.First(&user)

	return user, d.Error
}

func (u *UserRepositoryGormImpl) Update(updatedUser model.User) error {
	u.db.Save(&updatedUser)

	return nil
}

func (u *UserRepositoryGormImpl) Delete(user model.User) error {
	u.db.Delete(&user)

	return nil
}
