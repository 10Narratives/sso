package main

import (
	"github.com/10Narratives/sso/internal/app"
	"github.com/10Narratives/sso/internal/config"
	"github.com/10Narratives/sso/internal/lib/logger/sl"
)

func main() {
	// TODO: Make reading of configuration
	cfg := config.MustLoad()

	// TODO: Initialize logger
	log := sl.SetupLogger(cfg.Env)

	log.Info("Starting application")

	// TODO: Initialize application

	application := app.New(log, cfg)

	application.GRPCApp.MustRun()

	// TODO: run gRPC-server
	log.Info("Stopping application")
}
