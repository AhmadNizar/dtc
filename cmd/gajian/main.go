package main

import (
	"log"
	"os"

	gajian "github.com/AhmadNizar/dtc/cmd/gajian/http"
	"github.com/AhmadNizar/dtc/internal/config"
	clihandler "github.com/AhmadNizar/dtc/internal/delivery/cli"
	"github.com/urfave/cli/v2"
)

func main() {
	app := clihandler.New()

	app.Name = "backpack-platform"
	app.Usage = "Supply all data needed"
	app.Version = "1.0.0"

	app.Action = func(c *cli.Context) error {
		cfg := config.LoadConfig()

		log.Println("Starting gajian platform...")
		gajian.Start(cfg)

		return nil
	}

	log.Println("Starting server on :8080")
	err := app.Run(os.Args)

	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
