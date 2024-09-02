package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env               string `yaml:"env" env-default:"local"`
	SqliteStoragePath string `yaml:"sqlite_storage_path" env-required:"true"`
	HTTPServer        `yaml:"server"`
}

type HTTPServer struct {
	Address     string `yarm:"address"`
	Timeout     string `yarm:"timeout"`
	IdleTimeout string `yarm:"idle_timeout"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist, path: %v\n", configPath)
	}

	var config Config

	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatal("cannot read config")
	}

	return &config
}
