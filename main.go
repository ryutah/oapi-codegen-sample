package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ryutah/oapi-codegen-sample/interactor"
	"github.com/ryutah/oapi-codegen-sample/presentation/rest"
	"github.com/ryutah/oapi-codegen-sample/presentation/rest/oapi"
)

func main() {
	swagger, err := oapi.GetSwagger()
	if err != nil {
		log.Fatalf("failed to get swagger spec: %v", err)
	}
	swagger.Servers = nil

	r := chi.NewRouter()
	// NOTE(ryutah): エラーメッセージのフォーマットを変更できないため、利用しなくてもいいかも
	r.Use(oapimiddleware.OapiRequestValidator(swagger))
	r.Use(middleware.StripSlashes)

	// NOTE(ryutah): wireのような DIライブラリ使うのが望ましい
	if err := serve(oapi.HandlerWithOptions(interactor.InjectServer(), oapi.ChiServerOptions{
		BaseRouter:       r,
		ErrorHandlerFunc: rest.ErrorHandlerFunc,
	})); err != nil {
		log.Fatal(err.Error())
	}
}

// serve サーバ起動
// Graceulシャットダウン対応
func serve(handler http.Handler) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%s", port),
	}

	errChan := make(chan error)
	go func() {
		log.Printf("start server on port %v", port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errChan <- fmt.Errorf("failed to start server: %w", err)
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
	}

	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown gracefully: %w", err)
	}
	log.Println("Server shutdown")
	return nil
}
