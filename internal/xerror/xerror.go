package xerror

import "errors"

type Kind string

// see: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
const (
	InvalidArgument Kind = "INVALID_ARGUMENT"
	NotFound        Kind = "NOT_FOUND"
	Internal        Kind = "INTERNAL"
)

func (k Kind) Error() string {
	return string(k)
}

type Option func(e *Error)

func WithCause(cause error) Option {
	return func(e *Error) {
		e.cause = cause
	}
}

type Error struct {
	kind    Kind
	cause   error
	message string
}

// Is エラーの比較をする
//   1. エラーが完全一致するか
//   2. Kindが一致しているか
func (e *Error) Is(err error) bool {
	if e == nil {
		return err == nil
	}
	var er *Error
	if errors.As(err, &er) {
		return e == er
	}
	return errors.Is(e.kind, err)
}

func (e *Error) Unwrap() error {
	return e.cause
}

func New(kind Kind, message string, opts ...Option) error {
	e := &Error{
		kind:    kind,
		message: message,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

func (e *Error) Error() string {
	return e.message
}
