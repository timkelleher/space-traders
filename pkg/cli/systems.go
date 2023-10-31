package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var systemsCommands = cli.Command{
	Name: "systems",
	Subcommands: []*cli.Command{
		{
			Name:  "get",
			Usage: "get system data",
			Action: func(cCtx *cli.Context) error {
				systemSymbol := cCtx.Args().Get(0)
				if systemSymbol == "" {
					return errors.New("invalid system symbol")
				}

				res, err := apiClient.GetSystem(systemSymbol)
				if err != nil {
					return err
				}

				headers := []string{
					"Sector Symbol",
					"System Symbol",
					"Type",
					"Factions",
					"Waypoints",
				}

				data := [][]string{
					{
						res.Data.SectorSymbol,
						res.Data.Symbol,
						res.Data.Type,
						strings.Join(res.Data.GetFactions(), "\n"),
						strings.Join(res.Data.GetWaypoints(), "\n"),
					},
				}

				printTable(headers, data)

				return nil
			},
		},
		{
			Name:  "list",
			Usage: "list systems",
			Action: func(cCtx *cli.Context) error {
				page := cCtx.Args().Get(0)
				if page == "" {
					page = "1"
				}

				res, err := apiClient.ListSystems(page)
				if err != nil {
					return err
				}

				headers := []string{
					"Sector Symbol",
					"System Symbol",
					"Type",
					"Factions",
					"Waypoints",
				}

				var data [][]string
				for i, system := range res.Data {
					data = append(data, []string{
						system.SectorSymbol,
						system.Symbol,
						system.Type,
						strings.Join(res.Data[i].GetFactions(), "\n"),
						strings.Join(res.Data[i].GetWaypoints(), "\n"),
					})
				}

				printTable(headers, data)

				log.Info("Result", "page", res.Meta.Page, "total", res.Meta.Total)

				return nil
			},
		},
		{
			Name: "waypoints",
			Subcommands: []*cli.Command{
				{
					Name:  "search",
					Usage: "search for waypoints by type in a system",
					Action: func(cCtx *cli.Context) error {
						systemSymbol := cCtx.Args().Get(0)
						if systemSymbol == "" {
							return errors.New("invalid system symbol")
						}
						waypointType := cCtx.Args().Get(1)
						if waypointType == "" {
							return errors.New("invalid waypoint type")
						}

						res, err := apiClient.ListWaypointsByType(systemSymbol, strings.ToUpper(waypointType))
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
		},
	},
}
