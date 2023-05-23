package item

import (
	"github.com/YungBenn/go-gorm-fiber/internal/domain/entity"
	"github.com/YungBenn/go-gorm-fiber/internal/domain/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ItemService interface {
	GetAllItems() ([]entity.Item, error)
	GetItem(id string) (entity.Item, error)
	Create(c *fiber.Ctx) (*entity.Item, error)
	Update(c *fiber.Ctx) (*entity.Item, error)
	Delete(c *fiber.Ctx) (*entity.Item, error)
}

type ItemServiceImpl struct {
	itemRepo repository.ItemRepository
}

func NewItemService(itemRepo repository.ItemRepository) ItemService {
	return &ItemServiceImpl{itemRepo}
}

func (is *ItemServiceImpl) GetAllItems() ([]entity.Item, error) {
	return is.itemRepo.FindAllItems()
}

func (is *ItemServiceImpl) GetItem(id string) (entity.Item, error) {
	return is.itemRepo.FindItemByID(id)
}

func (is *ItemServiceImpl) Create(c *fiber.Ctx) (*entity.Item, error) {
	var input CreateItemInput

	if err := c.BodyParser(&input); err != nil {
		return nil, err
	}

	userID := input.UserID

	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	item := entity.Item{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Sold:        input.Sold,
		UserID:      input.UserID,
	}

	result, err := is.itemRepo.SaveItem(item, userID)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (is *ItemServiceImpl) Update(c *fiber.Ctx) (*entity.Item, error) {
	id := c.Params("id")

	var input UpdateItemInput

	if err := c.BodyParser(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	item := entity.Item{
		ID:          string(id),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Sold:        input.Sold,
	}

	result, err := is.itemRepo.UpdateItemByID(item)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (is *ItemServiceImpl) Delete(c *fiber.Ctx) (*entity.Item, error) {
	id := c.Params("id")

	item := entity.Item{
		ID: string(id),
	}

	result, err := is.itemRepo.DeleteItemByID(item)
	if err != nil {
		return nil, err
	}

	return result, err
}
