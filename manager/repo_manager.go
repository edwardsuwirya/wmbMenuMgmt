package manager

import "github.com/edwardsuwirya/wmbMenuMgmt/repository"

type RepoManager interface {
	MenuRepo() repository.IMenuRepository
}

type repoManager struct {
	infra Infra
}

func (rm *repoManager) MenuRepo() repository.IMenuRepository {
	return repository.NewMenuRepository(rm.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{infra}
}
