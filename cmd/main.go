package main

import (
	"github.com/janghanbin/go-storage-manager/configs"
	"github.com/janghanbin/go-storage-manager/internal/storage"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "Load configuration from `FILE`",
				Required: true,
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		configs.Cfg = configs.ReadConfiguration("configs/config.json")
		az := storage.NewClient("azure", configs.Cfg)
		az.GetList()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
