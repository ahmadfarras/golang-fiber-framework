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

	request := new(request.CreateUser)
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

func (u *UserController) GetAllUser(context *fiber.Ctx) error {

	response, err := u.userUsecase.GetAll()
	if err != nil {
		return err
	}

	return context.JSON(response)
}

func (u *UserController) GetById(context *fiber.Ctx) error {

	userId := context.Params("id")
	response, err := u.userUsecase.GetById(userId)
	if err != nil {
		return err
	}

	return context.JSON(response)
}

func (u *UserController) Update(context *fiber.Ctx) error {
	if !context.Is("json") {
		return context.SendStatus(fiber.StatusNotAcceptable)
	}

	userId := context.Params("id")

	request := new(request.UpdateUser)
	err := context.BodyParser(request)
	if err != nil {
		return err
	}

	err = u.userUsecase.Update(userId, *request)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserController) Delete(context *fiber.Ctx) error {

	userId := context.Params("id")

	err := u.userUsecase.DeleteById(userId)
	if err != nil {
		return err
	}

	return nil
}
