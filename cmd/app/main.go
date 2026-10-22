package main

import (
	"log"

	"github.com/dogukanttopcuoglu/go-clean-arch/config"
	"github.com/dogukanttopcuoglu/go-clean-arch/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalf("Application error: %s", err)
	}

}
