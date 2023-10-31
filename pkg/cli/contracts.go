package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var contractCommands = cli.Command{
	Name: "contracts",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list my contracts",
			Action: func(cCtx *cli.Context) error {
				res, err := apiClient.ListContracts()
				if err != nil {
					return err
				}

				headers := []string{
					"ID",
					"Fulfilled",
					"Accepted",
					"Deadline",
					"Faction Symbol",
					"Deliveries",
					"Payment",
				}

				var data [][]string
				for i, contract := range res.Data {
					data = append(data, []string{
						contract.ID,
						strconv.FormatBool(contract.Fulfilled),
						strconv.FormatBool(contract.Accepted),
						contract.Terms.Deadline,
						contract.FactionSymbol,
						strings.Join(res.Delieveries(i), "\n"),
						fmt.Sprintf("%d, %d", contract.Terms.Payment.OnAccepted, contract.Terms.Payment.OnFulfilled),
					})
				}

				printTable(headers, data)

				return nil
			},
		},
		{
			Name:  "accept",
			Usage: "accept a contract",
			Action: func(cCtx *cli.Context) error {
				id := cCtx.Args().Get(0)
				if id == "" {
					return errors.New("invalid id")
				}

				return apiClient.AcceptContract(id)
			},
		},
		{
			Name:  "deliver",
			Usage: "deliver resources for a contract",
			Action: func(cCtx *cli.Context) error {
				contractId := cCtx.Args().Get(0)
				if contractId == "" {
					return errors.New("invalid contract id")
				}

				shipSymbol := cCtx.Args().Get(1)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				cargoSymbol := cCtx.Args().Get(2)
				if cargoSymbol == "" {
					return errors.New("invalid cargo symbol")
				}

				cargoUnitsStr := cCtx.Args().Get(3)
				cargoUnits, err := strconv.Atoi(cargoUnitsStr)
				if err != nil || cargoUnits <= 0 {
					return errors.New("invalid cargo units")
				}

				return apiClient.DeliverContract(contractId, shipSymbol, cargoSymbol, cargoUnits)
			},
		},
	},
}
