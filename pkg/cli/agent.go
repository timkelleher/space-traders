package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

var agentCommands = cli.Command{
	Name: "agent",
	Subcommands: []*cli.Command{
		{
			Name:  "register",
			Usage: "register a new agent",
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
				return apiClient.RegisterAgent(callSign, faction)
			},
		},
		{
			Name:  "data",
			Usage: "view my agent data",
			Action: func(cCtx *cli.Context) error {
				res, err := apiClient.AgentData()
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
