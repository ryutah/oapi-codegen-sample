package model

import (
	"fmt"

	"github.com/ryutah/oapi-codegen-sample/internal/xerror"
)

type intID int64

func (i intID) Int() int64 {
	return int64(i)
}

func newIntID(i int64) (intID, error) {
	if i <= 0 {
		return 0, xerror.New(xerror.InvalidArgument, fmt.Sprintf("id must be grater than 0, actual value is %v", i))
	}
	return intID(i), nil
}

type HelloID struct {
	intID
}

func NewHelloID(i int64) (*HelloID, error) {
	id, err := newIntID(i)
	if err != nil {
		return nil, err
	}
	return &HelloID{
		intID: id,
	}, nil
}

type Hello struct {
	ID      HelloID
	Message string
}

func NewHello(message string) (*Hello, error) {
	if message == "" {
		return nil, xerror.New(xerror.InvalidArgument, "message must not be blank")
	}
	return &Hello{
		Message: message,
	}, nil
}
