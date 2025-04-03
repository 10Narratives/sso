package sqliteapp

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/10Narratives/sso/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	log  *slog.Logger
	conn *sql.DB
}

func New(log *slog.Logger, cfg config.StorageConfig) (*App, error) {
	const op = "sqliteapp.New"

	conn, err := sql.Open(cfg.Driver, cfg.DSN)
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
	const op = "sqliteapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping working with database")

	a.conn.Close()
}
