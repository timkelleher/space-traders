package cli

import (
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var factionsCommands = cli.Command{
	Name: "factions",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list my factions",
			Action: func(cCtx *cli.Context) error {
				res, err := apiClient.ListFactions()
				if err != nil {
					return err
				}

				headers := []string{
					"Symbol",
					"Name",
					"Headquarters",
					"Traits",
					"Is Recruiting",
				}

				var data [][]string
				for i, faction := range res.Data {
					data = append(data, []string{
						faction.Symbol,
						faction.Name,
						faction.Headquarters,
						strings.Join(res.Traits(i), "\n"),
						strconv.FormatBool(faction.IsRecruiting),
					})
				}

				printTable(headers, data)

				return nil
			},
		},
	},
}
