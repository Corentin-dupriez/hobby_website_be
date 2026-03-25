package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8000",
		db:   dbConfig{},
	}
	api := api{
		config: cfg,
	}

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "Error", err)
		os.Exit(1)
	}
}
