package handlers

import (
	"net/http"

	"github.com/YungBenn/go-gorm-fiber/internal/services/item"
	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	itemService item.ItemService
	c           *fiber.Ctx
}

func NewItemHandler(itemService item.ItemService, c *fiber.Ctx) ItemHandler {
	return ItemHandler{itemService, c}
}

func (ic *ItemHandler) Index(c *fiber.Ctx) error {
	data, err := ic.itemService.GetAllItems()
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": 200,
		"item":   data,
	})
}

func (ic *ItemHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := ic.itemService.GetItem(id)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": 200,
		"item":   data,
	})
}

func (ic *ItemHandler) Create(c *fiber.Ctx) error {
	data, err := ic.itemService.Create(c)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status": 201,
		"item":   data,
	})
}

func (ic *ItemHandler) Update(c *fiber.Ctx) error {
	data, err := ic.itemService.Update(c)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": 200,
		"item":   data,
	})
}

func (ic *ItemHandler) Delete(c *fiber.Ctx) error {
	data, err := ic.itemService.Delete(c)
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
