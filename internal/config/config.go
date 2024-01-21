package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	ServiceURL  string `env:"SERVICE_URL"`
	ServicePort string `env:"SERVICE_PORT"`
}

// ReadConfig функция для чтения конфигурации из файла и переменных окружения
func ReadConfig(cfg *Config, path string) error {
	return cleanenv.ReadConfig(path, cfg)
}
