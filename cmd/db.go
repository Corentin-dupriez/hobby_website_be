package main

import (
	"database/sql"
	"log"
	"log/slog"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func ConnectToDb() {
	envFile, _ := godotenv.Read(".env")
	envHost := envFile["DB_HOST"]
	envPort, err := strconv.Atoi(envFile["DB_PORT"])
	if err != nil {
		slog.Error("Error parsing the port from the environment", "Port", envPort)
	}
	envUser := envFile["DB_USER"]
	cfg := pq.Config{
		Host: envHost,
		Port: uint16(envPort),
		User: envUser,
	}

	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	db := sql.OpenDB(c)
	slog.Info("Connected to database", "DB_HOST", envHost)
	defer db.Close()
}
