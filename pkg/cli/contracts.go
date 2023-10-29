package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var contractCommands = cli.Command{
	Name:    "contracts",
	Aliases: []string{"c"},
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
						strconv.FormatBool(contract.Accepted),
						contract.Terms.Deadline,
						contract.FactionSymbol,
						strings.Join(res.TermsDelieveries(i), "\n"),
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
	},
}
