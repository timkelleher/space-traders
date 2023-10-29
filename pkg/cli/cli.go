package cli

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/timkelleher/space-traders/pkg/api"
	"github.com/urfave/cli/v2"
)

func Run() {
	client, err := api.New()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		HideHelp: true,
		Commands: []*cli.Command{
			{
				Name:    "register",
				Aliases: []string{"r"},
				Usage:   "register a new account",
				Action: func(cCtx *cli.Context) error {
					callSign := cCtx.Args().Get(0)
					if callSign == "" {
						return errors.New("invalid callsign")
					}
					faction := cCtx.Args().Get(1)
					if faction == "" {
						return errors.New("invalid faction")
					}

					fmt.Printf("call sign: %s, faction: %s\n", callSign, faction)
					return client.Register(callSign, faction)
				},
			},
			{
				Name:    "agent",
				Aliases: []string{"a"},
				Usage:   "view my agent data",
				Action: func(cCtx *cli.Context) error {
					res, err := client.MyAgent()
					if err != nil {
						return err
					}

					headers := []string{"Account ID", "Symbol", "HQ", "Starting Faction", "Credits", "Ship Count"}
					data := [][]string{
						{
							res.Data.AccountId,
							res.Data.Symbol,
							res.Data.Headquarters,
							res.Data.StartingFaction,
							strconv.Itoa(res.Data.Credits),
							strconv.Itoa(res.Data.ShipCount),
						},
					}
					printTable(headers, data)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
