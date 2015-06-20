package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	BindAddress string `toml:"bind_address"`
	Respond     string `toml:"always_respond"`
}

var config *Config

func init() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() error {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return err
	}

	return nil
}
