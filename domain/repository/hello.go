package repository

import (
	"context"

	"github.com/ryutah/oapi-codegen-sample/domain/model"
)

type Hello interface {
	// Get returns hello data
	Get(context.Context, model.HelloID) (*model.Hello, error)

	// Create inserts new hello object and return inserted object.
	// return object should be set id parameter.
	Create(context.Context, model.Hello) (*model.Hello, error)
}
