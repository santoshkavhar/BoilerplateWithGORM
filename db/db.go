package db

import (
	"context"
)

type Storer interface {
	ListProduct(context.Context) Product
	CreateProduct(context.Context)
	//GetUser(context.Context) (User, error)
	//Delete(context.Context, string) error
}
