package repository

import "ahmadfarras/fiberframework/internal/pkg/domain/model"

type UserRepository interface {
	Create(user model.User) error
}
