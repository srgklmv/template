package app

import (
	"fmt"
	"os"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

func Startup() error {
	// init config the way you prefer
	var cfg Config
	p, _ := os.Executable()
	err := cleanenv.ReadConfig(path.Join(path.Dir(p), "config.json"), &cfg)
	if err != nil {
		panic(fmt.Sprintf("Config initialisation error: %s", err))
	}

	// init logger

	// init adapters

	// init repos with adapters

	// init services with repos

	// init controllers with services

	// create web-app, like fiber.App

	// init app's api with controllers

	// run app

	// graceful shutdown

	panic("not implemented")
}
