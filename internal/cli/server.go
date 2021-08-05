package cli

import (
	"fmt"
	"github.com/jayhelton/webauthncf/internal/server"
	"github.com/urfave/cli/v2"
)

func serverCmd() *cli.Command {
	var port string

	cmd := &cli.Command{
		Name:  "server",
		Usage: "Run a local server with endpoints for test assertions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Value:       "9001",
				Usage:       "port for the server",
				Destination: &port,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Starting Server...")
			server.Start(port)
			return nil
		},
	}
	return cmd
}
