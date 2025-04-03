package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/10Narratives/sso/internal/app"
	"github.com/10Narratives/sso/internal/config"
	"github.com/10Narratives/sso/internal/lib/logger/sl"
)

func main() {
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	application := app.New(log, cfg)

	log.Info("starting application")
	go application.GRPCApp.MustRun()
	go application.DBApp.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("stopping application")

	application.GRPCApp.Stop()
	log.Info("grpc server stopped")

	application.DBApp.Stop()
	log.Info("working with database stopped")
}
