package cli

import (
	"log"
	"os"

	"github.com/timkelleher/space-traders/pkg/api"
	"github.com/urfave/cli/v2"
)

var apiClient *api.Client

func Run() {
	client, err := api.New()
	if err != nil {
		log.Fatal(err)
	}
	apiClient = client

	app := &cli.App{
		HideHelp: true,
		Commands: []*cli.Command{
			&agentCommands,
			&contractCommands,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
