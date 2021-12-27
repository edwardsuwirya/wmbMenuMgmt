package manager

import "github.com/edwardsuwirya/wmbMenuMgmt/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.IMenuUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) MenuUseCase() usecase.IMenuUseCase {
	return usecase.NewMenuUseCase(uc.repo.MenuRepo())
}
func NewUseCaseManger(manager RepoManager) UseCaseManager {
	return &useCaseManager{repo: manager}
}
