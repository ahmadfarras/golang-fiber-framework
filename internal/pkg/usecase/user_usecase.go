package usecase

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model/request"
	"ahmadfarras/fiberframework/internal/pkg/domain/model/response"
)

type UserUsecase interface {
	Create(request request.CreateUser) error
	GetAll() ([]response.UserDetailResponse, error)
	GetById(id string) (response.UserDetailResponse, error)
	Update(id string, request request.UpdateUser) error
	DeleteById(id string) error
}
