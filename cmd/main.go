package main

import (
	"fmt"
	"github.com/janghanbin/go-storage-manager/configs"
	"github.com/janghanbin/go-storage-manager/internal/storage"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	cApp = &cli.App{}
)

func init() {
	cApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"},
			Usage:    "Load configuration from `FILE`",
			Required: true,
		},
	}

	cApp.Action = func(c *cli.Context) error {
		configs.Cfg = configs.ReadConfiguration("configs/config.json")
		az := storage.NewClient("azure", configs.Cfg)
		az.GetList()
		return nil
	}
}

func main() {
	if err := cApp.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
