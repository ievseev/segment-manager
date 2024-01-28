package main

import (
	"fmt"
	"log/slog"
	"os"
	"segment-manager/internal/config"
	"segment-manager/internal/storage/postgres"
)

//func CreateSegmentHandler(resp http.ResponseWriter, req *http.Request) {
//
//}
//
//func GetSegmentHandler(resp http.ResponseWriter, req *http.Request) {
//
//}

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

	fmt.Println(storage)

	//TODO: init router - chi (совместим с net/http)

	//TODO: run server
}
