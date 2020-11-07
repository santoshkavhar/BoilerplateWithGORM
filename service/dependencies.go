package service

import "santoshkavhar/BoilerplateWithGORM/db"

type Dependencies struct {
	Store db.Storer
	// define other service dependencies
}
