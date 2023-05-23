package repository

import (
	"log"

	"github.com/YungBenn/go-gorm-fiber/internal/domain/entity"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAllItems() ([]entity.Item, error)
	FindItemByID(id string) (entity.Item, error)
	SaveItem(item entity.Item, userID string) (*entity.Item, error)
	UpdateItemByID(item entity.Item) (*entity.Item, error)
	DeleteItemByID(item entity.Item) (*entity.Item, error)
}

type ItemRepositoryImpl struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &ItemRepositoryImpl{db}
}

func (i *ItemRepositoryImpl) FindAllItems() ([]entity.Item, error) {
	var items []entity.Item

	err := i.db.Model([]entity.Item{}).Find(&items).Error

	return items, err
}

func (i *ItemRepositoryImpl) FindItemByID(id string) (entity.Item, error) {
	var item = entity.Item{ID: id}

	result := i.db.Model([]entity.Item{}).First(&item)
	if result.RowsAffected == 0 {
		log.Println("Wrong ID")
	}

	return item, nil
}

func (i *ItemRepositoryImpl) SaveItem(item entity.Item, userID string) (*entity.Item, error) {
	var user entity.User
	if err := i.db.Where("id = ?", userID).First(&user).Error; err != nil {
		panic("Error searching user id")
	}

	user.Items = append(user.Items, item)

	if err := i.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *ItemRepositoryImpl) UpdateItemByID(item entity.Item) (*entity.Item, error) {
	result := i.db.Save(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (i *ItemRepositoryImpl) DeleteItemByID(item entity.Item) (*entity.Item, error) {
	result := i.db.Delete(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}
