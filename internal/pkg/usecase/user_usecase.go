package usecase

import "ahmadfarras/fiberframework/internal/pkg/domain/model/request"

type UserUsecase interface {
	Create(request request.CreateNewUser) error
}
