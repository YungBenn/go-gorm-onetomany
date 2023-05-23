package user

import (
	"github.com/YungBenn/go-gorm-fiber/internal/domain/entity"
	"github.com/YungBenn/go-gorm-fiber/internal/domain/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetAllUsers() ([]entity.User, error)
	GetUser(id string) (entity.User, error)
	Create(c *fiber.Ctx) (*entity.User, error)
	Update(c *fiber.Ctx) (*entity.User, error)
	Delete(c *fiber.Ctx) (*entity.User, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo}
}

func (us *UserServiceImpl) GetAllUsers() ([]entity.User, error) {
	return us.userRepo.FindAllUsers()
}

func (us *UserServiceImpl) GetUser(id string) (entity.User, error) {
	return us.userRepo.FindUserByID(id)
}

func (us *UserServiceImpl) Create(c *fiber.Ctx) (*entity.User, error) {
	var input CreateUserInput

	if err := c.BodyParser(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	result, err := us.userRepo.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *UserServiceImpl) Update(c *fiber.Ctx) (*entity.User, error) {
	id := c.Params("id")

	var input UpdateUserInput

	if err := c.BodyParser(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		ID:       string(id),
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	result, err := us.userRepo.UpdateUserByID(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *UserServiceImpl) Delete(c *fiber.Ctx) (*entity.User, error) {
	id := c.Params("id")

	user := entity.User{
		ID: string(id),
	}

	result, err := us.userRepo.DeleteUserByID(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
