package pkg

import (
	"ahmadfarras/fiberframework/internal/pkg/controller"
	repository "ahmadfarras/fiberframework/internal/pkg/repository/impl"
	usecase "ahmadfarras/fiberframework/internal/pkg/usecase/impl"

	"gorm.io/gorm"
)

func NewUserController(db *gorm.DB) *controller.UserController {
	ur := repository.NewUserRepositoryGormImpl(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	return uc
}
