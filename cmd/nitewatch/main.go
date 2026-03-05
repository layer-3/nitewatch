package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	"github.com/layer-3/nitewatch/config"
	"github.com/layer-3/nitewatch/service"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "worker" {
		fmt.Fprintln(os.Stderr, "usage: nitewatch worker")
		os.Exit(1)
	}

	conf, err := loadConfig()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	if conf.Blockchain.PrivateKey == "" {
		fmt.Print("Enter private key: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			slog.Error("Failed to read private key", "error", err)
			os.Exit(1)
		}
		fmt.Println() // Print newline after input
		conf.Blockchain.PrivateKey = strings.TrimSpace(string(bytePassword))
		if conf.Blockchain.PrivateKey == "" {
			slog.Error("Private key cannot be empty")
			os.Exit(1)
		}
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

func loadConfig() (*config.Config, error) {
	if raw := os.Getenv("NITEWATCH_CONFIG"); raw != "" {
		return config.LoadFromEnv(raw)
	}
	configPath := os.Getenv("NITEWATCH_CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}
	return config.Load(configPath)
}
