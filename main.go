package main

import (
	"Video-Streaming-Platform/config"
	"Video-Streaming-Platform/database"
	"Video-Streaming-Platform/server"
	"fmt"
)

func main() {
	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Envs.DB_HOST, config.Envs.DB_PORT, config.Envs.DB_USERNAME,
		config.Envs.DB_PASSWORD, config.Envs.DB_NAME)
	db := database.NewDBConnection(connStr)
	defer db.Close()

	server := server.NewServer(":8080", nil)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
