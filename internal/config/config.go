package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string
}
type Config struct {
	Env          string
	Storage_path string
	Http_server  HttpServer
}

func LoadConfigAndSerialize() *Config{
	var configFilePath string = "config.yaml"
	_, err := os.Stat(configFilePath)
	if err != nil {
		flags := flag.String("config", "", "for setting config file path")
		flag.Parse()
		configFilePath = *flags
	}
	if configFilePath == "" {
		log.Fatal("Config path not set")
	}

	var cnf Config
	err = cleanenv.ReadConfig(configFilePath, &cnf)
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	return &cnf
}
