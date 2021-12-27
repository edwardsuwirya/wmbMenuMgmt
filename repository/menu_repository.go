package repository

import (
	"fmt"
	"github.com/edwardsuwirya/wmbMenuMgmt/entity"
	"gorm.io/gorm"
	"log"
)

type IMenuRepository interface {
	GetAllMenu() ([]entity.Menu, error)
	GetMenuByName(name string) ([]entity.Menu, error)
	GetMenuById(id string) (*entity.Menu, error)
}

type MenuRepository struct {
	db *gorm.DB
}

func (m *MenuRepository) GetAllMenu() ([]entity.Menu, error) {
	var list []entity.Menu
	err := m.db.Find(&list).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return list, nil
}

func (m *MenuRepository) GetMenuByName(name string) ([]entity.Menu, error) {
	var list []entity.Menu
	err := m.db.Where("menu_name ilike ?", fmt.Sprintf("%%%s%%", name)).Find(&list).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return list, nil
}

func (m *MenuRepository) GetMenuById(id string) (*entity.Menu, error) {
	var menu entity.Menu
	err := m.db.Where("id = ?", id).Find(&menu).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &menu, nil
}

func NewMenuRepository(resource *gorm.DB) IMenuRepository {
	menuRepo := &MenuRepository{db: resource}
	return menuRepo
}
