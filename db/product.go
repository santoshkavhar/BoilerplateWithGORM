package db

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (s *pgStore) ListProduct(ctx context.Context) (product Product) {
	s.db.First(&product, 1)
	return
}

func (s *pgStore) CreateProduct(ctx context.Context) {
	s.db.Create(&Product{Code: "D42", Price: 100})
	return
}
