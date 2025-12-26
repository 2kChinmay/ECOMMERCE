package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string
	Storage_path string
	Http_server  HttpServer
}

type HttpServer struct {
	Address string
}

func LoadAndSerializeConfig() *Config {
	var configPath string
	configPath = os.Getenv("CFG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "for setting config file path")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path not set")
		}
	}

	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatalf("Config file not found at config path: %s", configPath)
	}

	var cfg Config
	e := cleanenv.ReadConfig(configPath, &cfg)
	if e != nil {
		log.Fatalf("Failed to read config: %s", e.Error())
	}

	SerializedConfig := &cfg
	return SerializedConfig
}
