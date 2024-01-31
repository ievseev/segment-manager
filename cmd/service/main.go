package main

import (
	"fmt"
	"log/slog"
	"os"
	"segment-manager/internal/config"
	"segment-manager/internal/storage/postgres"
	"segment-manager/internal/store"
)

func main() {
	//TODO: init config - cleanenv
	cfg := config.MustLoad("config/.env")

	//TODO: init logger - slog
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	//TODO: init storage - postgres
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage") // TODO сделать ошибку детальнее
		os.Exit(1)                          // идти дальше смысла нет, выходим
	}

	segmentDB := store.New(storage)
	err = segmentDB.SaveSegment("segment")
	if err != nil {
		return
	}

	fmt.Println(storage)

	//TODO: init router - chi (совместим с net/http)

	//TODO: run server
}
