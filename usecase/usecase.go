package usecase

import (
	"context"
	"errors"

	"github.com/ryutah/oapi-codegen-sample/internal/xerror"
)

// Input ports
type (
	HelloInputPort interface {
		Detail(context.Context, HelloDetailOutputPort, HelloDetailRequest)
		Create(context.Context, HelloCreateOutputPort, HelloCreateRequest)
	}
)

// Output Ports
type (
	ErrorOutputPort interface {
		InternalError(context.Context, error)
		InvalidArgument(context.Context, error)
		NotFound(context.Context, error)
	}

	HelloDetailOutputPort interface {
		ErrorOutputPort
		OK(context.Context, HelloDetailResponse)
	}

	HelloCreateOutputPort interface {
		ErrorOutputPort
		Created(context.Context, HelloCreateResponse)
	}
)

func handleError(ctx context.Context, out ErrorOutputPort, err error) {
	switch {
	case errors.Is(err, xerror.Internal):
		out.InternalError(ctx, err)
	case errors.Is(err, xerror.InvalidArgument):
		out.InvalidArgument(ctx, err)
	case errors.Is(err, xerror.NotFound):
		out.NotFound(ctx, err)
	default:
		out.InternalError(ctx, err)
	}
}
