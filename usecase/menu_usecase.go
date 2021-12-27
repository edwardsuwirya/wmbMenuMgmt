package usecase

import (
	"github.com/edwardsuwirya/wmbMenuMgmt/entity"
	"github.com/edwardsuwirya/wmbMenuMgmt/repository"
)

type IMenuUseCase interface {
	GetAllMenu() ([]entity.Menu, error)
	SearchMenuByName(menuName string) ([]entity.Menu, error)
	SearchMenuById(menuId string) (*entity.Menu, error)
}

type MenuUseCase struct {
	menuRepo repository.IMenuRepository
}

func NewMenuUseCase(menuRepo repository.IMenuRepository) IMenuUseCase {
	return &MenuUseCase{menuRepo: menuRepo}
}

func (m *MenuUseCase) GetAllMenu() ([]entity.Menu, error) {
	return m.menuRepo.GetAllMenu()
}

func (m *MenuUseCase) SearchMenuByName(menuName string) ([]entity.Menu, error) {
	return m.menuRepo.GetMenuByName(menuName)
}

func (m *MenuUseCase) SearchMenuById(menuId string) (*entity.Menu, error) {
	return m.menuRepo.GetMenuById(menuId)
}
