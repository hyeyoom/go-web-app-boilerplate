package engine

import "github.com/hyeyoom/go-web-app-boilerplate/domain"

type (
	StorageFactory interface {
		NewMemberRepository() domain.MemberRepository
		NewProductRepository() domain.ProductRepository
	}
)
