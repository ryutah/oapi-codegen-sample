package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/ryutah/oapi-codegen-sample/internal/xerror"
	"github.com/ryutah/oapi-codegen-sample/presentation/rest/internal/renderer"
	"github.com/ryutah/oapi-codegen-sample/presentation/rest/oapi"
	"github.com/ryutah/oapi-codegen-sample/usecase"
)

type Server struct {
	usecases struct {
		hello usecase.HelloInputPort
	}
}

var _ oapi.ServerInterface = (*Server)(nil)

func NewServer(hello usecase.HelloInputPort) *Server {
	return &Server{
		usecases: struct {
			hello usecase.HelloInputPort
		}{
			hello: hello,
		},
	}
}

func (s *Server) PostHello(w http.ResponseWriter, r *http.Request) {
	out := renderer.NewPostHelloRendrer(w, r)

	payload, err := parsePostHelloRequest(r)
	if err != nil {
		out.InvalidArgument(r.Context(), err)
		return
	}

	s.usecases.hello.Create(r.Context(), out, usecase.HelloCreateRequest{
		Message: payload.Message,
	})
}

func (s *Server) GetHello(w http.ResponseWriter, r *http.Request, helloId int) {
	s.usecases.hello.Detail(r.Context(), renderer.NewGetHelloRendrer(w, r), usecase.HelloDetailRequest{
		ID: int64(helloId),
	})
}

func parsePostHelloRequest(r *http.Request) (*oapi.PostHelloJSONRequestBody, error) {
	var payload oapi.PostHelloJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, xerror.New(
			xerror.InvalidArgument,
			fmt.Sprintf("failed to parse request body: %v", err),
			xerror.WithCause(err),
		)
	}
	// NOTE(ryutah) https://github.com/go-playground/validator を使うかどうか要検討
	if err := validator.New().Struct(payload); err != nil {
		return nil, xerror.New(
			xerror.InvalidArgument,
			err.Error(),
			xerror.WithCause(err),
		)
	}
	return &payload, nil
}

func ErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, oapi.Error{
		Message: err.Error(),
	})
}
