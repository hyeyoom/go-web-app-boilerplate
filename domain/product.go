package domain

import "github.com/hyeyoom/go-web-app-boilerplate/domain/base"

type Product struct {
	base.DefaultModel
	Name string
}
