package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
	"url_shortener/libs"
)

var (
	localEnv = libs.GetWorkPath() + "/config/url_shortener/local.yaml"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" envDefault:"local"`
	HttpServer `yaml:"http_server" env:"HTTP_SERVER" env-required:"true"`
	Database   `yaml:"database" env:"DATABASE" env-required:"true"`
}

type HttpServer struct {
	Host        string        `yaml:"host" env:"HOST" envDefault:"localhost"`
	Port        string        `yaml:"port" env:"PORT" envDefault:"8080"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" envDefault:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" envDefault:"60s"`
}

type Database struct {
	Host     string `yaml:"host" env:"HOST" envDefault:"localhost"`
	Port     string `yaml:"port" env:"PORT" envDefault:"3306"`
	User     string `yaml:"user" env:"USER" envDefault:"root"`
	Password string `yaml:"password" env:"PASSWORD" envDefault:"root"`
	DB       string `yaml:"db" env:"DB" envDefault:"test"`
}

func MustLoad() Config {
	configPath := localEnv
	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		configPath = os.Getenv("CONFIG_PATH")
	}
	if configPath == "" {
		panic("CONFIG_PATH environment variable not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("CONFIG_PATH does not exist")
	}
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		panic("Error reading config: " + err.Error())
	}
	return config
}
