package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pyKis/test/internal/config"
	"github.com/pyKis/test/internal/handler"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	r := chi.NewRouter()
	handler.RegisterRoutes(r, handler.Dependencies{
		AssetsFS: http.Dir(cfg.AssetsFS),
	})

	s := http.Server{
		Addr: cfg.ServerAddr,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")
		s.Shutdown(ctx)
	}()

	slog.Info("starting server",slog.String("addr",cfg.ServerAddr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}