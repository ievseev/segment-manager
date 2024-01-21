package main

import (
	"log"
	"segment-manager/internal/config"
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
	var serviceConfig config.Config
	err := config.ReadConfig(&serviceConfig, "config/.env")
	if err != nil {
		log.Fatal("Failed to read log: %v", err)
	}

	//TODO: init logger - slog (log/slog начиная с 1.21)

	//TODO: init storage - postgres

	//TODO: init router - chi (совместим с net/http)

	//TODO: run server
}
