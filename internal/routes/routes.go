package routes

import (
	"github.com/YungBenn/go-gorm-fiber/internal/domain/repository"
	"github.com/YungBenn/go-gorm-fiber/internal/handlers"
	"github.com/YungBenn/go-gorm-fiber/internal/services/item"
	"github.com/YungBenn/go-gorm-fiber/internal/services/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var c *fiber.Ctx

func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService, c)

	userRouter := app.Group("api/v1/users")

	userRouter.Post("/", userHandler.Create)
	userRouter.Get("/", userHandler.Index)
	userRouter.Get("/:id", userHandler.GetById)
	userRouter.Put("/:id", userHandler.Update)
	userRouter.Delete("/:id", userHandler.Delete)
}

func SetupItemRoutes(app *fiber.App, db *gorm.DB) {
	itemRepository := repository.NewItemRepository(db)
	itemService := item.NewItemService(itemRepository)
	itemHandler := handlers.NewItemHandler(itemService, c)

	itemRouter := app.Group("api/v1/items")

	itemRouter.Post("/", itemHandler.Create)
	itemRouter.Get("/", itemHandler.Index)
	itemRouter.Get("/:id", itemHandler.GetById)
	itemRouter.Put("/:id", itemHandler.Update)
	itemRouter.Post("/:id", itemHandler.Delete)
}
