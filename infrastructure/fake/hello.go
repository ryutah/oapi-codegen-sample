package fake

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/ryutah/oapi-codegen-sample/domain/model"
	"github.com/ryutah/oapi-codegen-sample/domain/repository"
	"github.com/ryutah/oapi-codegen-sample/internal/xerror"
)

type HelloRepository struct {
	rand  *rand.Rand
	store map[int64]*model.Hello
}

var _ repository.Hello = (*HelloRepository)(nil)

// nolint:gosec
func NewHelloRepository() *HelloRepository {
	return &HelloRepository{
		rand:  rand.New(rand.NewSource(1)),
		store: make(map[int64]*model.Hello),
	}
}

// Get returns hello data
func (h *HelloRepository) Get(ctx context.Context, id model.HelloID) (*model.Hello, error) {
	hello, ok := h.store[id.Int()]
	if !ok {
		return nil, xerror.New(xerror.NotFound, fmt.Sprintf("id(%v) record is not exists", id.Int()))
	}
	return hello, nil
}

// Create inserts new hello object and return inserted object.
// return object should be set id parameter.
func (h *HelloRepository) Create(ctx context.Context, hello model.Hello) (*model.Hello, error) {
	// it should be declare Recreate function to HelloID struct
	// like below
	//
	//  func RecreateHelloID(i int64) HelloID {
	//		return HelloID{intID: int64(i)}
	//  }
	newID, err := model.NewHelloID(h.rand.Int63())
	if err != nil {
		return nil, err
	}
	hello.ID = *newID
	h.store[newID.Int()] = &hello

	return &hello, nil
}
