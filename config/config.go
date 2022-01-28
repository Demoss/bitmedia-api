package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	MongoDB struct {
		Host       string `json:"host"`
		Port       string `json:"port"`
		Database   string `json:"database"`
		AuthDB     string `json:"auth_db"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Collection string `json:"collection"`
	} `json:"mongodb"`
	URL    string `json:"URL"`
	APIKey string `json:"APIKey"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
