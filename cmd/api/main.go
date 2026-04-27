package main

import (
	"api-skeleton/internal/config"
	"api-skeleton/internal/db"
	"api-skeleton/internal/notes"
	"api-skeleton/internal/server"
	"fmt"
	"log"
)

// config -> db -> router -> run server

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	pool, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	defer func() {
		pool.Close()
	}()

	repo := notes.NewPostgresRepository(pool)
	handler := notes.NewHandler(repo)
	r := server.NewRouter(handler)
	addr := fmt.Sprintf(":%s", cfg.Server_Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}

}
