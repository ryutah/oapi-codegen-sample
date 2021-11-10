package usecase

import (
	"context"

	"github.com/ryutah/oapi-codegen-sample/domain/model"
	"github.com/ryutah/oapi-codegen-sample/domain/repository"
)

type HelloDetailRequest struct {
	ID int64
}

type HelloDetailResponse struct {
	ID      int64
	Message string
}

func newHelloDetailResponse(h model.Hello) *HelloDetailResponse {
	return &HelloDetailResponse{
		ID:      h.ID.Int(),
		Message: h.Message,
	}
}

type HelloCreateRequest struct {
	Message string
}

type HelloCreateResponse struct {
	ID      int64
	Message string
}

func newHelloCreateResponse(h model.Hello) *HelloCreateResponse {
	return &HelloCreateResponse{
		ID:      h.ID.Int(),
		Message: h.Message,
	}
}

type Hello struct {
	repository struct {
		hello repository.Hello
	}
}

var _ HelloInputPort = (*Hello)(nil)

func NewHello(helloRepo repository.Hello) *Hello {
	return &Hello{
		repository: struct {
			hello repository.Hello
		}{
			hello: helloRepo,
		},
	}
}

func (h *Hello) Detail(ctx context.Context, out HelloDetailOutputPort, req HelloDetailRequest) {
	id, err := model.NewHelloID(req.ID)
	if err != nil {
		handleError(ctx, out, err)
		return
	}
	hello, err := h.repository.hello.Get(ctx, *id)
	if err != nil {
		handleError(ctx, out, err)
		return
	}

	out.OK(ctx, *newHelloDetailResponse(*hello))
}

func (h *Hello) Create(ctx context.Context, out HelloCreateOutputPort, req HelloCreateRequest) {
	hello, err := model.NewHello(req.Message)
	if err != nil {
		handleError(ctx, out, err)
		return
	}
	newHello, err := h.repository.hello.Create(ctx, *hello)
	if err != nil {
		handleError(ctx, out, err)
		return
	}
	out.Created(ctx, *newHelloCreateResponse(*newHello))
}
