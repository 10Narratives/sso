package main

import (
	"github.com/10Narratives/sso/internal/config"
)

func main() {
	// TODO: Make reading of configuration
	_ = config.MustLoad()

	// TODO: Initialize logger

	// TODO: Initialize application

	// TODO: run gRPC-server
}
