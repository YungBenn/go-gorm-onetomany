package repository

import (
	"log"

	"github.com/YungBenn/go-gorm-fiber/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUsers() ([]entity.User, error)
	FindUserByID(id string) (entity.User, error)
	SaveUser(user entity.User) (*entity.User, error)
	UpdateUserByID(user entity.User) (*entity.User, error)
	DeleteUserByID(user entity.User) (*entity.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) FindAllUsers() ([]entity.User, error) {
	var users []entity.User

	err := u.db.Model([]entity.User{}).Preload("Items").Find(&users).Error
	return users, err
}

func (u *UserRepositoryImpl) FindUserByID(id string) (entity.User, error) {
	var user = entity.User{ID: id}

	result := u.db.Model([]entity.User{}).Preload("Items").First(&user)
	if result.RowsAffected == 0 {
		log.Println("Error")
	}

	return user, nil
}

func (u *UserRepositoryImpl) SaveUser(user entity.User) (*entity.User, error) {
	result := u.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserRepositoryImpl) UpdateUserByID(user entity.User) (*entity.User, error) {
	result := u.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserRepositoryImpl) DeleteUserByID(user entity.User) (*entity.User, error) {
	result := u.db.Delete(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
