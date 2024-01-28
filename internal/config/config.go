package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	//ServicePath string `env:"SERVICE_PATH"`
	StoragePath string `env:"STORAGE_PATH"`
}

// MustLoad функция для чтения конфигурации из файла и переменных окружения
func MustLoad(path string) *Config {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Fatalf("Read config error: %v", err)
	}

	return cfg
}
