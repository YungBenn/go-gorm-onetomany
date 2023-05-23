package handlers

import (
	"net/http"

	"github.com/YungBenn/go-gorm-fiber/internal/services/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user.UserService
	c           *fiber.Ctx
}

func NewUserHandler(userService user.UserService, c *fiber.Ctx) UserHandler {
	return UserHandler{userService, c}
}

func (uc *UserHandler) Index(c *fiber.Ctx) error {
	data, err := uc.userService.GetAllUsers()
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": 200,
		"data":   data,
	})

}

func (uc *UserHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := uc.userService.GetUser(id)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": 200,
		"data":   data,
	})
}

func (uc *UserHandler) Create(c *fiber.Ctx) error {
	data, err := uc.userService.Create(c)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": 200,
		"data":   data,
	})
}

func (uc *UserHandler) Update(c *fiber.Ctx) error {
	data, err := uc.userService.Update(c)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": 200,
		"data":   data,
	})
}

func (uc *UserHandler) Delete(c *fiber.Ctx) error {
	data, err := uc.userService.Delete(c)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     200,
		"message":    "Delete Success",
		"deleted_id": data.ID,
	})
}
