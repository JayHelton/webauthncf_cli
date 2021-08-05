package cli

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Commands: []*cli.Command{
			serverCmd(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
