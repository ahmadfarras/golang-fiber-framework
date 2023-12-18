package controller

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model/request"
	"ahmadfarras/fiberframework/internal/pkg/domain/model/response"
	"ahmadfarras/fiberframework/internal/pkg/handler"
	"ahmadfarras/fiberframework/internal/pkg/usecase"
	"net/http"

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
		r := handler.HandleError(err)
		context.SendStatus(r.StatusCode)
		return context.JSON(r)
	}

	userId, err := u.userUsecase.Create(*request)
	if err != nil {
		r := handler.HandleError(err)
		context.SendStatus(r.StatusCode)
		return context.JSON(r)
	}

	location, err := context.GetRouteURL("users.show", fiber.Map{"id": userId})
	if err != nil {
		r := handler.HandleError(err)
		context.SendStatus(r.StatusCode)
		return context.JSON(r)
	}

	context.Location(location)
	context.SendStatus(fiber.StatusCreated)
	return context.JSON(response.BuildSuccessResponse(http.StatusCreated, "Created", nil))
}

func (u *UserController) GetAllUser(context *fiber.Ctx) error {

	resp, err := u.userUsecase.GetAll()
	if err != nil {
		return err
	}

	context.SendStatus(http.StatusOK)
	return context.JSON(response.BuildSuccessResponse(http.StatusOK, "Success", resp))
}

func (u *UserController) GetById(context *fiber.Ctx) error {

	userId := context.Params("id")
	resp, err := u.userUsecase.GetById(userId)
	if err != nil {
		r := handler.HandleError(err)
		context.SendStatus(r.StatusCode)
		return context.JSON(r)
	}

	context.SendStatus(http.StatusOK)
	return context.JSON(response.BuildSuccessResponse(http.StatusOK, "Success", resp))
}

func (u *UserController) Update(context *fiber.Ctx) error {
	if !context.Is("json") {
		return context.SendStatus(fiber.StatusNotAcceptable)
	}

	userId := context.Params("id")

	request := new(request.UpdateUser)
	err := context.BodyParser(request)
	if err != nil {
		r := handler.HandleError(err)
		context.SendStatus(r.StatusCode)
		return context.JSON(r)
	}

	err = u.userUsecase.Update(userId, *request)
	if err != nil {
		r := handler.HandleError(err)
		context.SendStatus(r.StatusCode)
		return context.JSON(r)
	}

	context.SendStatus(http.StatusOK)
	return context.JSON(response.BuildSuccessResponse(http.StatusOK, "Success", nil))
}

func (u *UserController) Delete(context *fiber.Ctx) error {

	userId := context.Params("id")

	err := u.userUsecase.DeleteById(userId)
	if err != nil {
		r := handler.HandleError(err)
		context.SendStatus(r.StatusCode)
		return context.JSON(r)
	}

	context.SendStatus(http.StatusOK)
	return context.JSON(response.BuildSuccessResponse(http.StatusOK, "Success", nil))
}
