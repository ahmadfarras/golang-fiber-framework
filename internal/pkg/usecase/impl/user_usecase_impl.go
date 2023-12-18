package usecase

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model"
	"ahmadfarras/fiberframework/internal/pkg/domain/model/request"
	"ahmadfarras/fiberframework/internal/pkg/domain/model/response"
	"ahmadfarras/fiberframework/internal/pkg/repository"
	"ahmadfarras/fiberframework/internal/pkg/usecase"

	"github.com/google/uuid"
)

type UserUsecaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &UserUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *UserUsecaseImpl) Create(request request.CreateUser) (string, error) {
	var userId string

	newUser := model.CreateNewUser(request.FullName, request.Password, request.Email)

	err := u.userRepository.Create(newUser)
	if err != nil {
		return userId, err
	}

	userId = newUser.ID.String()
	return userId, nil
}

func (u *UserUsecaseImpl) GetAll() ([]response.UserDetailResponse, error) {
	var users []model.User
	err := u.userRepository.GetAll(&users)
	if err != nil {
		return []response.UserDetailResponse{}, err
	}

	return response.BuildUserListResponse(users), nil
}

func (u *UserUsecaseImpl) GetById(id string) (response.UserDetailResponse, error) {
	userId := uuid.MustParse(id)
	user, err := u.userRepository.GetById(userId)
	if err != nil {
		return response.UserDetailResponse{}, err
	}

	return response.BuildUserDetailResponse(user), nil

}

func (u *UserUsecaseImpl) Update(id string, request request.UpdateUser) error {
	userId := uuid.MustParse(id)
	user, err := u.userRepository.GetById(userId)
	if err != nil {
		return err
	}

	user.UpdateUser(request.FullName, request.Password, request.Email)
	u.userRepository.Update(user)

	return nil
}

func (u *UserUsecaseImpl) DeleteById(id string) error {
	userId := uuid.MustParse(id)
	user, err := u.userRepository.GetById(userId)
	if err != nil {
		return err
	}

	u.userRepository.Delete(user)

	return nil
}
