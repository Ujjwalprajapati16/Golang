package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

// env-default:"production" will be used if ENV is not set

// env-required:"true" will panic if ENV is not set
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	// Load the config from file, env, etc.
	// If any error occurs, panic
	var configPath string

	// Check if the config path is set in the environment variable
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			// If the config path is not set, panic
			log.Fatal("config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var cfg Config
	// Read the config from the file
	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("failed to load config: %v", err.Error())
	}

	return &cfg
}
