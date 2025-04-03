package app

import (
	"log/slog"

	grpcapp "github.com/10Narratives/sso/internal/app/grpc"
	"github.com/10Narratives/sso/internal/config"
)

type App struct {
	GRPCApp *grpcapp.App
}

func New(log *slog.Logger, cfg *config.Config) *App {
	// TODO: initialize storage

	// TODO: initialize auth service

	grpcApp := grpcapp.New(log, cfg.GRPC.Port)
	return &App{grpcApp}
}
