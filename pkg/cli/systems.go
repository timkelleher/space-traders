package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var systemsCommands = cli.Command{
	Name: "systems",
	Subcommands: []*cli.Command{
		{
			Name: "asteroids",
			Subcommands: []*cli.Command{
				{
					Name:  "list",
					Usage: "list shipyards in a system",
					Action: func(cCtx *cli.Context) error {
						res, err := apiClient.ListWaypointsByType("X1-FH96", "ENGINEERED_ASTEROID")
						if err != nil {
							return err
						}

						headers := []string{
							"System Symbol",
							"Symbol",
							"Coordinates",
							"Traits",
							"Faction Symbol",
							"Orbits",
						}

						var data [][]string
						for i, waypoint := range res.Data {
							data = append(data, []string{
								waypoint.SystemSymbol,
								waypoint.Symbol,
								fmt.Sprintf("%d,%d", waypoint.X, waypoint.Y),
								strings.Join(res.Traits(i), "\n"),
								waypoint.Faction.Symbol,
								waypoint.Orbits,
							})
						}

						printTable(headers, data)

						return nil
					},
				},
			},
		},
		{
			Name: "shipyards",
			Subcommands: []*cli.Command{
				{
					Name:  "list",
					Usage: "list shipyards in a system",
					Action: func(cCtx *cli.Context) error {
						res, err := apiClient.ListWaypointsByTrait("X1-FH96", "SHIPYARD")
						if err != nil {
							return err
						}

						headers := []string{
							"System Symbol",
							"Symbol",
							"Coordinates",
							"Traits",
							"Faction Symbol",
							"Orbits",
						}

						var data [][]string
						for i, waypoint := range res.Data {
							data = append(data, []string{
								waypoint.SystemSymbol,
								waypoint.Symbol,
								fmt.Sprintf("%d,%d", waypoint.X, waypoint.Y),
								strings.Join(res.Traits(i), "\n"),
								waypoint.Faction.Symbol,
								waypoint.Orbits,
							})
						}

						printTable(headers, data)

						return nil
					},
				},
				{
					Name:  "ships",
					Usage: "list ships in a shipyard",
					Action: func(cCtx *cli.Context) error {
						res, err := apiClient.ListShipyardShips("X1-FH96", "X1-FH96-A2")
						if err != nil {
							return err
						}

						headers := []string{
							"Type",
							"Name",
							"Supply",
							"Activity",
							"Purchase Price",
							"Frame",
							"Reactor",
							"Engine",
							"Modules",
						}

						var data [][]string
						for i, ship := range res.Data.Ships {
							data = append(data, []string{
								ship.Type,
								ship.Name,
								ship.Supply,
								ship.Activity,
								strconv.Itoa(ship.PurchasePrice),
								ship.Frame.Name,
								ship.Reactor.Name,
								ship.Engine.Name,
								strings.Join(res.ShipModules(i), "\n"),
							})
						}

						printTable(headers, data)

						return nil
					},
				},
			},
		},
	},
}
