package automation

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/timkelleher/space-traders/pkg/api"
)

var discardResources = []string{
	"COPPER_ORE",
	"ICE_WATER",
	"QUARTZ_SAND",
	"SILICON_CRYSTALS",
}

func FillShipCargo(apiClient *api.Client, shipSymbol string) error {
	cargo, _ := apiClient.ShipCargo(shipSymbol)
	log.Info("Ship Cargo Status", "units", cargo.Data.Units, "capacity", cargo.Data.Capacity)
	time.Sleep(1 * time.Second)

	for cargo.Data.Units < cargo.Data.Capacity {
		extractRes, err := apiClient.ExtractResources(shipSymbol)
		if err != nil {
			return err
		}
		log.Info("Extracted Resources", "units", extractRes.Data.Extraction.Yield.Units, "symbol", extractRes.Data.Extraction.Yield.Symbol)
		timer := extractRes.Data.Cooldown.RemainingSeconds

		log.Info("Waiting for Extraction Cooldown", "seconds", timer)
		time.Sleep(time.Duration(timer+1) * time.Second)

		AutoDiscardResources(apiClient, shipSymbol)
		time.Sleep(1 * time.Second)

		cargo, err = apiClient.ShipCargo(shipSymbol)
		if err != nil {
			return err
		}
	}

	return nil
}

func FlyShip(apiClient *api.Client, shipSymbol, destinationWaypoint string) error {
	ship, err := apiClient.GetShip(shipSymbol)
	if err != nil {
		return err
	}
	log.Info("Ship Status",
		"flight_status", ship.Data.Nav.Status,
		"current_fuel", ship.Data.Fuel.Current, "fuel_capacity", ship.Data.Fuel.Capacity,
	)
	time.Sleep(1 * time.Second)

	if ship.Data.Nav.WaypointSymbol == destinationWaypoint {
		log.Warn("Ship Already at Destination")
		return nil
	}

	log.Info("Ship Itinerary", "current", ship.Data.Nav.WaypointSymbol, "dest", destinationWaypoint)

	if ship.Data.Nav.Status != "IN_TRANSIT" {
		if err := apiClient.FlyShip(shipSymbol, destinationWaypoint); err != nil {
			log.Warn("yup", "error", err)
			return err
		}
	}
	time.Sleep(1 * time.Second)

	ship, err = apiClient.GetShip(shipSymbol)
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	for ship.Data.Nav.Status != "IN_ORBIT" {
		log.Info("Waiting for Ship to Arrive to Destination")
		time.Sleep(10 * time.Second)

		ship, err = apiClient.GetShip(shipSymbol)
		if err != nil {
			return err
		}
	}

	return nil
}

func AutoDiscardResources(apiClient *api.Client, shipSymbol string) error {
	time.Sleep(1 * time.Second)

	cargo, err := apiClient.ShipCargo(shipSymbol)
	if err != nil {
		return err
	}

	for _, item := range cargo.Data.Inventory {
		for _, discarded := range discardResources {
			if discarded == item.Symbol {
				if err := apiClient.JettisonShipCargo(shipSymbol, item.Symbol, item.Units); err != nil {
					return err
				}

				time.Sleep(1 * time.Second)
			}
		}
	}
	return nil
}

func AutoRefillFuel(apiClient *api.Client, contractId, shipSymbol string) error {
	ship, err := apiClient.GetShip(shipSymbol)
	if err != nil {
		return err
	}

	log.Info("Performing Optional Dock and Refuel", "status", ship.Data.Nav.Status)

	if ship.Data.Nav.Status != "DOCKED" {
		if err := apiClient.DockShip(shipSymbol); err != nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	if err := apiClient.RefuelShip(shipSymbol); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	if _, err := apiClient.OrbitShip(shipSymbol); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	return nil
}

func FulfillContract(apiClient *api.Client, contractId, shipSymbol string) error {
	contract, err := apiClient.GetContract(contractId)
	if err != nil {
		return err
	}

	log.Info("Contract Accepted Status", "accepted", contract.Data.Accepted)
	if !contract.Data.Accepted {
		if err := apiClient.AcceptContract(contractId); err != nil {
			return err
		}
	}

	cargo, err := apiClient.ShipCargo(shipSymbol)
	if err != nil {
		return err
	}

	for i := 0; i < len(contract.Data.Terms.Deliver); i++ {
		log.Info(fmt.Sprintf("Starting Delivery %d/%d", i+1, len(contract.Data.Terms.Deliver)))

		for contract.Data.Terms.Deliver[i].UnitsFulfilled < contract.Data.Terms.Deliver[i].UnitsRequired {
			log.Info("Delivery Progress",
				"delivery", i+1,
				"fulfilled", contract.Data.Terms.Deliver[i].UnitsFulfilled,
				"required", contract.Data.Terms.Deliver[i].UnitsRequired,
			)

			if err := AutoRefillFuel(apiClient, contractId, shipSymbol); err != nil {
				return nil
			}

			destinationWaypoint := contract.Data.Terms.Deliver[i].DestinationSymbol
			log.Info("Setting Destination", "symbol", destinationWaypoint)

			resourceSymbol := contract.Data.Terms.Deliver[i].TradeSymbol
			log.Info("Setting Resource", "symbol", resourceSymbol)

			sourceWaypoints, err := apiClient.ListWaypointsByType("X1-FH96", "ENGINEERED_ASTEROID")
			if err != nil {
				return err
			}
			if len(sourceWaypoints.Data) == 0 {
				log.Warn("Cannot find source for resource")
				return errors.New("cannot find source for resource %s")
			}

			sourceWaypoint := sourceWaypoints.Data[0].Symbol
			log.Info("Setting Source", "symbol", sourceWaypoint)

			if cargo.Data.Units == 0 {
				log.Info("Flying to Source")
				if err := FlyShip(apiClient, shipSymbol, sourceWaypoint); err != nil {
					return err
				}

				log.Info("Filling Ship Cargo", "symbol", shipSymbol)
				if err := FillShipCargo(apiClient, shipSymbol); err != nil {
					return err
				}
			}

			log.Info("Flying to Destination")
			if err := FlyShip(apiClient, shipSymbol, destinationWaypoint); err != nil {
				return err
			}

			if err := apiClient.DockShip(shipSymbol); err != nil {
				return err
			}

			cargo, err = apiClient.ShipCargo(shipSymbol)
			if err != nil {
				return err
			}

			for _, item := range cargo.Data.Inventory {
				if item.Symbol == resourceSymbol {
					if err := apiClient.DeliverContract(contractId, shipSymbol, item.Symbol, item.Units); err != nil {
						return nil
					}
					log.Info("Delivered Cargo for Contract", "symbol", item.Symbol, "units", item.Units)
				} else {
					sell, err := apiClient.SellShipCargo(shipSymbol, item.Symbol, item.Units)
					if err != nil {
						time.Sleep(1 * time.Second)

						log.Warn("Cargo Sale Failed", "symbol", item.Symbol)
						apiClient.JettisonShipCargo(shipSymbol, item.Symbol, item.Units)
					} else {
						log.Info("Sold Cargo", "symbol", item.Symbol, "price", sell.Data.Transaction.TotalPrice)
					}
				}

				time.Sleep(1 * time.Second)
			}

			contract, err = apiClient.GetContract(contractId)
			if err != nil {
				log.Warn("why here tho")
				return err
			}
		}
	}

	log.Info("Contract Fulfilled")
	apiClient.FulfillContract(contractId, shipSymbol)

	return nil
}

func JettisonCargoExclude(apiClient *api.Client, shipSymbol, cargoSymbol string, cargo *api.MyShipsCargoResponse) error {
	for _, item := range cargo.Data.Inventory {
		if cargoSymbol != item.Symbol {
			log.Info("Jettisoning Ship Cargo", "units", item.Units, "symbol", item.Symbol)
			if err := apiClient.JettisonShipCargo(shipSymbol, item.Symbol, item.Units); err != nil {
				return err
			}
		}
	}
	return nil
}
