package renderer

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ryutah/oapi-codegen-sample/presentation/rest/oapi"
	"github.com/ryutah/oapi-codegen-sample/usecase"
)

type rendrer struct {
	w http.ResponseWriter
	r *http.Request
}

var _ usecase.ErrorOutputPort = (*rendrer)(nil)

func newRendrer(w http.ResponseWriter, r *http.Request) *rendrer {
	return &rendrer{
		w: w,
		r: r,
	}
}

func (r *rendrer) InternalError(ctx context.Context, err error) {
	render.Status(r.r, http.StatusInternalServerError)
	render.JSON(r.w, r.r, oapi.Error{
		Message: err.Error(),
	})
}

func (r *rendrer) InvalidArgument(ctx context.Context, err error) {
	render.Status(r.r, http.StatusBadRequest)
	render.JSON(r.w, r.r, oapi.Error{
		Message: err.Error(),
	})
}

func (r *rendrer) NotFound(ctx context.Context, err error) {
	render.Status(r.r, http.StatusNotFound)
	render.JSON(r.w, r.r, oapi.Error{
		Message: err.Error(),
	})
}

type GetHello struct {
	*rendrer
}

func NewGetHelloRendrer(w http.ResponseWriter, r *http.Request) *GetHello {
	return &GetHello{
		rendrer: newRendrer(w, r),
	}
}

var _ usecase.HelloDetailOutputPort = (*GetHello)(nil)

func (g *GetHello) OK(ctx context.Context, resp usecase.HelloDetailResponse) {
	render.Status(g.r, http.StatusOK)
	id := int(resp.ID)
	render.JSON(g.rendrer.w, g.rendrer.r, oapi.Hello{
		Id:      &id,
		Message: resp.Message,
	})
}

type PostHelloRendrer struct {
	*rendrer
}

func NewPostHelloRendrer(w http.ResponseWriter, r *http.Request) *PostHelloRendrer {
	return &PostHelloRendrer{
		rendrer: newRendrer(w, r),
	}
}

var _ usecase.HelloCreateOutputPort = (*PostHelloRendrer)(nil)

func (g *PostHelloRendrer) Created(ctx context.Context, resp usecase.HelloCreateResponse) {
	render.Status(g.r, http.StatusCreated)
	id := int(resp.ID)
	render.JSON(g.rendrer.w, g.rendrer.r, oapi.Hello{
		Id:      &id,
		Message: resp.Message,
	})
}
