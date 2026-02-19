package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/layer-3/nitewatch/config"
	"github.com/layer-3/nitewatch/service"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "worker" {
		fmt.Fprintln(os.Stderr, "usage: nitewatch worker")
		os.Exit(1)
	}

	configPath := os.Getenv("NITEWATCH_CONFIG")
	if configPath == "" {
		configPath = "config.yaml"
	}

	conf, err := config.Load(configPath)
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	svc, err := service.New(*conf)
	if err != nil {
		slog.Error("Failed to create service", "error", err)
		os.Exit(1)
	}

	if err := svc.RunWorker(); err != nil {
		slog.Error("Worker failed", "error", err)
		os.Exit(1)
	}
}
