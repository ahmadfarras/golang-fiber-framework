package repository

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user model.User) error
	GetAll(user *[]model.User) error
	GetById(id uuid.UUID) (model.User, error)
	Update(user model.User) error
	Delete(user model.User) error
}
