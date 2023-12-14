package usecase

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model"
	"ahmadfarras/fiberframework/internal/pkg/domain/model/request"
	"ahmadfarras/fiberframework/internal/pkg/repository"
	"ahmadfarras/fiberframework/internal/pkg/usecase"
)

type UserUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &UserUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *UserUsecaseImpl) Create(request request.CreateNewUser) error {
	newUser := model.CreateNewUser(request.FullName, request.Password, request.Email)

	err := u.userRepository.Create(newUser)
	if err != nil {
		return err
	}

	return nil
}
