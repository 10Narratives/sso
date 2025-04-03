package app

import (
	"log/slog"

	sqliteapp "github.com/10Narratives/sso/internal/app/database/sqlite"
	grpcapp "github.com/10Narratives/sso/internal/app/grpc"
	"github.com/10Narratives/sso/internal/config"
)

type App struct {
	GRPCApp *grpcapp.App
	DBApp   *sqliteapp.App
}

func New(log *slog.Logger, cfg *config.Config) *App {
	grpcApp := grpcapp.New(log, cfg.GRPC.Port)

	databaseApp, err := sqliteapp.New(log, cfg.Storage)
	if err != nil {
		panic(err)
	}

	return &App{grpcApp, databaseApp}
}
