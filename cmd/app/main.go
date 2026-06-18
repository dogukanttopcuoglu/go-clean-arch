package main

import (
	"log"

	"github.com/dogukanttopcuoglu/clean-lab/config"
	"github.com/dogukanttopcuoglu/clean-lab/internal/app"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
