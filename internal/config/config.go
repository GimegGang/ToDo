package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env" required:"true"`
	Address     string        `yaml:"address" required:"true"`
	Timeout     time.Duration `yaml:"timeout" required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" required:"true"`
	StoragePath string        `yaml:"storage_path" required:"true"`
}

func MustLoad(path string) Config {
	if _, err := os.Stat(path); err != nil {
		log.Fatalln("config file not exists:", path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalln(err)
	}
	return cfg
}
