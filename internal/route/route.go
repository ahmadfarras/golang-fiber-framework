package route

import (
	"ahmadfarras/fiberframework/internal/pkg"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitHttpRoute(app *fiber.App, db *gorm.DB) {
	app.Get("/", func(context *fiber.Ctx) error {
		return context.JSON("Hello, World!")
	})

	userController := pkg.NewUserController(db)
	usersGroup := app.Group("/users")
	usersGroup.Post("/create", userController.CreateNewUser)
	usersGroup.Get("", userController.GetAllUser)
	usersGroup.Get("/:id", userController.GetById)
	usersGroup.Put("/:id", userController.Update)
	usersGroup.Delete("/:id", userController.Delete)
}
