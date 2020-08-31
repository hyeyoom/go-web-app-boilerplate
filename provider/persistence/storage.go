package persistence

import (
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
	"github.com/hyeyoom/go-web-app-boilerplate/engine"
	"gorm.io/gorm"
	"sync"
)

type (
	storageFactory struct {
		db *gorm.DB
	}

	repository struct {
		db *gorm.DB
	}
)

var (
	memberRepositoryInstance  domain.MemberRepository
	memberRepositoryOnce      sync.Once
	productRepositoryInstance domain.ProductRepository
	productRepositoryOnce     sync.Once
)

func NewStorage(database *gorm.DB) engine.StorageFactory {
	return &storageFactory{database}
}

func (sf *storageFactory) NewMemberRepository() domain.MemberRepository {
	memberRepositoryOnce.Do(func() {
		memberRepositoryInstance = NewMemberRepository(sf.db)
	})
	return memberRepositoryInstance
}

func (sf *storageFactory) NewProductRepository() domain.ProductRepository {
	return nil
}
