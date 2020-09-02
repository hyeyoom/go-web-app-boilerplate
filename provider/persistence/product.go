package persistence

import (
	"github.com/hyeyoom/go-web-app-boilerplate/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	repository
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	var pr productRepository
	pr.db = db
	return &pr
}

func (mr *productRepository) Create(product *domain.Product) {
	mr.db.Create(product)
}
