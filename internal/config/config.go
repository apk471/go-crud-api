package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type HttpServer struct{
	Addr string `yaml:"address"`
}

type Config struct{
	Env string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
	HttpServer  `yaml:"http_server"`
}


func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if(configPath == ""){
		flags := flag.String("config" , "" , "path to config file")
		flag.Parse()
		configPath = *flags

		if(configPath == ""){
			log.Fatal("config path is required")
		}
	}

	_, err := os.Stat(configPath)
	if(os.IsNotExist(err)){
		log.Fatalf("config file not found: %v", err)
	}

	var cfg Config
	err = cleanenv.ReadConfig(configPath, &cfg)

	if(err != nil){
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
	
}

