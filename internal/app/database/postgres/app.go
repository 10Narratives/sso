package postgresapp

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/10Narratives/sso/internal/config"
	"github.com/jackc/pgx/v5"
)

type App struct {
	log  *slog.Logger
	conn *pgx.Conn
}

func New(log *slog.Logger, cfg *config.StorageConfig) (*App, error) {
	const op = "postgresapp.New"

	conn, err := pgx.Connect(context.Background(), cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &App{log: log, conn: conn}, nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop() {
	const op = "postgresapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping working with database")

	a.conn.Close(context.Background())
}
