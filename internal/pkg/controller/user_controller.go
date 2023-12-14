package controller

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model/request"
	"ahmadfarras/fiberframework/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (u *UserController) CreateNewUser(context *fiber.Ctx) error {
	if !context.Is("json") {
		return context.SendStatus(fiber.StatusNotAcceptable)
	}

	request := new(request.CreateNewUser)
	err := context.BodyParser(request)
	if err != nil {
		return err
	}

	err = u.userUsecase.Create(*request)
	if err != nil {
		return err
	}

	return nil
}
