package controller

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model/request"
	"ahmadfarras/fiberframework/internal/pkg/domain/model/response"
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

	userId, err := u.userUsecase.Create(*request)
	if err != nil {
		return err
	}

	location, err := context.GetRouteURL("users.show", fiber.Map{"id": userId})
	if err != nil {
		return err
	}

	context.Location(location)
	context.SendStatus(fiber.StatusCreated)
	return context.JSON(response.Response{
		StatusCode: 201,
		Message:    "Created",
	})
}

func (u *UserController) GetAllUser(context *fiber.Ctx) error {

	response, err := u.userUsecase.GetAll()
	if err != nil {
		return err
	}

	context.SendStatus(fiber.StatusOK)
	return context.JSON(response)
}

func (u *UserController) GetById(context *fiber.Ctx) error {

	userId := context.Params("id")
	response, err := u.userUsecase.GetById(userId)
	if err != nil {
		return err
	}

	context.SendStatus(fiber.StatusOK)
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

	context.SendStatus(fiber.StatusOK)
	return context.JSON(response.Response{
		StatusCode: 200,
		Message:    "Success",
	})
}

func (u *UserController) Delete(context *fiber.Ctx) error {

	userId := context.Params("id")

	err := u.userUsecase.DeleteById(userId)
	if err != nil {
		return err
	}

	context.SendStatus(fiber.StatusOK)
	return context.JSON(response.Response{
		StatusCode: 200,
		Message:    "Success",
	})
}
