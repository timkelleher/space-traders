package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var shipsCommands = cli.Command{
	Name: "ships",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list my ships",
			Action: func(cCtx *cli.Context) error {
				res, err := apiClient.ListMyShips()
				if err != nil {
					return err
				}

				headers := []string{
					"Symbol",
					"Faction",
					"Role",
					"Waypoint",
					"Nav Status",
					"Nav Flight Mode",
					"Fuel",
					"Cargo",
					"Arrival Time",
				}

				var data [][]string
				for _, ship := range res.Data {
					data = append(data, []string{
						ship.Symbol,
						ship.Registration.FactionSymbol,
						ship.Registration.Role,
						ship.Nav.WaypointSymbol,
						ship.Nav.Status,
						ship.Nav.FlightMode,
						fmt.Sprintf("%d/%d", ship.Fuel.Current, ship.Fuel.Capacity),
						fmt.Sprintf("%d/%d", ship.Cargo.Units, ship.Cargo.Capacity),
						ship.Nav.Route.Arrival.String(),
					})
				}

				printTable(headers, data)

				return nil
			},
		},
		{
			Name:  "get",
			Usage: "get ship data",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				res, err := apiClient.GetShip(shipSymbol)
				if err != nil {
					return err
				}

				headers := []string{
					"Symbol",
					"Faction",
					"Role",
					"Waypoint",
					"Nav Status",
					"Nav Flight Mode",
					"Fuel",
					"Cargo",
					"Arrival Time",
				}

				data := [][]string{
					{
						res.Data.Symbol,
						res.Data.Registration.FactionSymbol,
						res.Data.Registration.Role,
						res.Data.Nav.WaypointSymbol,
						res.Data.Nav.Status,
						res.Data.Nav.FlightMode,
						fmt.Sprintf("%d/%d", res.Data.Fuel.Current, res.Data.Fuel.Capacity),
						fmt.Sprintf("%d/%d", res.Data.Cargo.Units, res.Data.Cargo.Capacity),
						res.Data.Nav.Route.Arrival.String(),
					},
				}

				printTable(headers, data)

				return nil
			},
		},
		{
			Name:  "purchase",
			Usage: "purchase a ship",
			Action: func(cCtx *cli.Context) error {
				shipType := cCtx.Args().Get(0)
				if shipType == "" {
					return errors.New("invalid ship type")
				}

				shipyardWaypointSymbol := cCtx.Args().Get(1)
				if shipyardWaypointSymbol == "" {
					return errors.New("invalid shipyard waypoint symbol")
				}

				return apiClient.PurchaseShip(shipType, shipyardWaypointSymbol)
			},
		},
		{
			Name:  "orbit",
			Usage: "move a ship into orbit",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				res, err := apiClient.OrbitShip(shipSymbol)
				if err != nil {
					return err
				}

				headers := []string{
					"System Symbol",
					"Waypoint Symbol",
					"Status",
					"Flight Mode",
					"Arrival Time",
				}

				data := [][]string{
					{
						res.Data.Nav.SystemSymbol,
						res.Data.Nav.WaypointSymbol,
						res.Data.Nav.Status,
						res.Data.Nav.FlightMode,
						res.Data.Nav.Route.Arrival.String(),
					},
				}

				printTable(headers, data)

				return nil
			},
		},
		{
			Name:  "fly",
			Usage: "fly a ship to a waypoint",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				waypointSymbol := cCtx.Args().Get(1)
				if waypointSymbol == "" {
					return errors.New("invalid waypoint symbol")
				}

				return apiClient.FlyShip(shipSymbol, waypointSymbol)
			},
		},
		{
			Name:  "dock",
			Usage: "dock a ship",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				return apiClient.DockShip(shipSymbol)
			},
		},
		{
			Name:  "refuel",
			Usage: "refuel a ship",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				return apiClient.RefuelShip(shipSymbol)
			},
		},
		{
			Name:  "cargo",
			Usage: "view a ship's cargo",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("ship symbol")
				}

				res, err := apiClient.ShipCargo(shipSymbol)
				if err != nil {
					return err
				}

				headers := []string{
					"Capacity",
					"Units",
					"Inventory",
				}

				data := [][]string{
					{
						strconv.Itoa(res.Data.Capacity),
						strconv.Itoa(res.Data.Units),
						strings.Join(res.Inventory(), "\n"),
					},
				}

				printTable(headers, data)

				return nil
			},
		},
		{
			Name:  "extract",
			Usage: "have a ship extract resources from its current location",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				res, err := apiClient.ExtractResources(shipSymbol)
				if err != nil {
					printError(err)
					return nil
				}

				headers := []string{
					"Remaining Cooldown",
					"Extraction",
				}

				data := [][]string{
					{
						strconv.Itoa(res.Data.Cooldown.RemainingSeconds),
						fmt.Sprintf("%d %s", res.Data.Extraction.Yield.Units, res.Data.Extraction.Yield.Symbol),
					},
				}

				printTable(headers, data)

				return nil
			},
		},
		{
			Name:  "survey",
			Usage: "have a ship survey resources from its current location",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				res, err := apiClient.ShipSurvey(shipSymbol)
				if err != nil {
					printError(err)
					return nil
				}

				headers := []string{
					"Symbol",
					"Units",
				}

				var data [][]string
				for _, survey := range res.Data.Surveys {
					data = append(data, []string{survey.Symbol, survey.Size})
				}

				printTable(headers, data)

				log.Info("Survey Expiration", "remaining_seconds", res.Data.Cooldown.RemainingSeconds)

				return nil
			},
		},
		{
			Name:  "sell",
			Usage: "sell a resource from a ship's cargo",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				cargoSymbol := cCtx.Args().Get(1)
				if cargoSymbol == "" {
					return errors.New("invalid cargo symbol")
				}

				cargoUnitsStr := cCtx.Args().Get(2)
				cargoUnits, err := strconv.Atoi(cargoUnitsStr)
				if err != nil || cargoUnits <= 0 {
					return errors.New("invalid cargo units")
				}

				_, err = apiClient.SellShipCargo(shipSymbol, cargoSymbol, cargoUnits)
				if err != nil {
					printError(err)
				}

				return nil
			},
		},
		{
			Name:  "jettison",
			Usage: "jettison a resource from a ship's cargo",
			Action: func(cCtx *cli.Context) error {
				shipSymbol := cCtx.Args().Get(0)
				if shipSymbol == "" {
					return errors.New("invalid ship symbol")
				}

				cargoSymbol := cCtx.Args().Get(1)
				if cargoSymbol == "" {
					return errors.New("invalid cargo symbol")
				}

				cargoUnitsStr := cCtx.Args().Get(2)
				cargoUnits, err := strconv.Atoi(cargoUnitsStr)
				if err != nil || cargoUnits <= 0 {
					return errors.New("invalid cargo units")
				}

				err = apiClient.JettisonShipCargo(shipSymbol, cargoSymbol, cargoUnits)
				if err != nil {
					printError(err)
				}

				return nil
			},
		},
	},
}
