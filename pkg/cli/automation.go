package cli

import (
	"errors"

	"github.com/timkelleher/space-traders/pkg/automation"
	"github.com/urfave/cli/v2"
)

var automationCommands = cli.Command{
	Name: "automation",
	Subcommands: []*cli.Command{
		{
			Name:  "cargoFill",
			Usage: "fill a ship's cargo with one resource and jettison everything else",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				return automation.FillShipCargo(apiClient, shipSymbol)
			},
		},
		{
			Name:  "contractFulfill",
			Usage: "accept a contract and deliver all cargo to destination until fulfilled",
			Action: func(cCtx *cli.Context) error {
				contractId := cCtx.Args().Get(0)
				if contractId == "" {
					return errors.New("invalid contract id")
				}

				shipSymbol := cCtx.Args().Get(1)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				return automation.FulfillContract(apiClient, contractId, shipSymbol)
			},
		},
	},
}
