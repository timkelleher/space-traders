package api

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

type RegisterAgentResponse struct {
	Data struct {
		Agent struct {
			AccountID       string `json:"accountId"`
			Symbol          string `json:"symbol"`
			Headquarters    string `json:"headquarters"`
			Credits         int    `json:"credits"`
			StartingFaction string `json:"startingFaction"`
			ShipCount       int    `json:"shipCount"`
		} `json:"agent"`
		Contract struct {
			ID            string `json:"id"`
			FactionSymbol string `json:"factionSymbol"`
			Type          string `json:"type"`
			Terms         struct {
				Deadline time.Time `json:"deadline"`
				Payment  struct {
					OnAccepted  int `json:"onAccepted"`
					OnFulfilled int `json:"onFulfilled"`
				} `json:"payment"`
				Deliver []struct {
					TradeSymbol       string `json:"tradeSymbol"`
					DestinationSymbol string `json:"destinationSymbol"`
					UnitsRequired     int    `json:"unitsRequired"`
					UnitsFulfilled    int    `json:"unitsFulfilled"`
				} `json:"deliver"`
			} `json:"terms"`
			Accepted         bool      `json:"accepted"`
			Fulfilled        bool      `json:"fulfilled"`
			Expiration       time.Time `json:"expiration"`
			DeadlineToAccept time.Time `json:"deadlineToAccept"`
		} `json:"contract"`
		Faction struct {
			Symbol       string `json:"symbol"`
			Name         string `json:"name"`
			Description  string `json:"description"`
			Headquarters string `json:"headquarters"`
			Traits       []struct {
				Symbol      string `json:"symbol"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"traits"`
			IsRecruiting bool `json:"isRecruiting"`
		} `json:"faction"`
		Ship struct {
			Symbol       string `json:"symbol"`
			Registration struct {
				Name          string `json:"name"`
				FactionSymbol string `json:"factionSymbol"`
				Role          string `json:"role"`
			} `json:"registration"`
			Nav struct {
				SystemSymbol   string `json:"systemSymbol"`
				WaypointSymbol string `json:"waypointSymbol"`
				Route          struct {
					Destination struct {
						Symbol       string `json:"symbol"`
						Type         string `json:"type"`
						SystemSymbol string `json:"systemSymbol"`
						X            int    `json:"x"`
						Y            int    `json:"y"`
					} `json:"destination"`
					Departure struct {
						Symbol       string `json:"symbol"`
						Type         string `json:"type"`
						SystemSymbol string `json:"systemSymbol"`
						X            int    `json:"x"`
						Y            int    `json:"y"`
					} `json:"departure"`
					Origin struct {
						Symbol       string `json:"symbol"`
						Type         string `json:"type"`
						SystemSymbol string `json:"systemSymbol"`
						X            int    `json:"x"`
						Y            int    `json:"y"`
					} `json:"origin"`
					DepartureTime time.Time `json:"departureTime"`
					Arrival       time.Time `json:"arrival"`
				} `json:"route"`
				Status     string `json:"status"`
				FlightMode string `json:"flightMode"`
			} `json:"nav"`
			Crew struct {
				Current  int    `json:"current"`
				Required int    `json:"required"`
				Capacity int    `json:"capacity"`
				Rotation string `json:"rotation"`
				Morale   int    `json:"morale"`
				Wages    int    `json:"wages"`
			} `json:"crew"`
			Frame struct {
				Symbol         string `json:"symbol"`
				Name           string `json:"name"`
				Description    string `json:"description"`
				Condition      int    `json:"condition"`
				ModuleSlots    int    `json:"moduleSlots"`
				MountingPoints int    `json:"mountingPoints"`
				FuelCapacity   int    `json:"fuelCapacity"`
				Requirements   struct {
					Power int `json:"power"`
					Crew  int `json:"crew"`
					Slots int `json:"slots"`
				} `json:"requirements"`
			} `json:"frame"`
			Reactor struct {
				Symbol       string `json:"symbol"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				Condition    int    `json:"condition"`
				PowerOutput  int    `json:"powerOutput"`
				Requirements struct {
					Power int `json:"power"`
					Crew  int `json:"crew"`
					Slots int `json:"slots"`
				} `json:"requirements"`
			} `json:"reactor"`
			Engine struct {
				Symbol       string `json:"symbol"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				Condition    int    `json:"condition"`
				Speed        int    `json:"speed"`
				Requirements struct {
					Power int `json:"power"`
					Crew  int `json:"crew"`
					Slots int `json:"slots"`
				} `json:"requirements"`
			} `json:"engine"`
			Cooldown struct {
				ShipSymbol       string    `json:"shipSymbol"`
				TotalSeconds     int       `json:"totalSeconds"`
				RemainingSeconds int       `json:"remainingSeconds"`
				Expiration       time.Time `json:"expiration"`
			} `json:"cooldown"`
			Modules []struct {
				Symbol       string `json:"symbol"`
				Capacity     int    `json:"capacity"`
				Range        int    `json:"range"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				Requirements struct {
					Power int `json:"power"`
					Crew  int `json:"crew"`
					Slots int `json:"slots"`
				} `json:"requirements"`
			} `json:"modules"`
			Mounts []struct {
				Symbol       string   `json:"symbol"`
				Name         string   `json:"name"`
				Description  string   `json:"description"`
				Strength     int      `json:"strength"`
				Deposits     []string `json:"deposits"`
				Requirements struct {
					Power int `json:"power"`
					Crew  int `json:"crew"`
					Slots int `json:"slots"`
				} `json:"requirements"`
			} `json:"mounts"`
			Cargo struct {
				Capacity  int `json:"capacity"`
				Units     int `json:"units"`
				Inventory []struct {
					Symbol      string `json:"symbol"`
					Name        string `json:"name"`
					Description string `json:"description"`
					Units       int    `json:"units"`
				} `json:"inventory"`
			} `json:"cargo"`
			Fuel struct {
				Current  int `json:"current"`
				Capacity int `json:"capacity"`
				Consumed struct {
					Amount    int       `json:"amount"`
					Timestamp time.Time `json:"timestamp"`
				} `json:"consumed"`
			} `json:"fuel"`
		} `json:"ship"`
		Token string `json:"token"`
	} `json:"data"`
}

func (c Client) RegisterAgent(callSign, faction string) (*RegisterAgentResponse, error) {
	res := &RegisterAgentResponse{}
	req := map[string]any{"symbol": callSign, "faction": faction}

	url := url + "/register"
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetBody(req).
		SetResult(res).
		SetError(&errorResponse{}).
		Post(url)
	if err != nil {
		return nil, errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return nil, errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return res, nil
}

type MyAgentResponse struct {
	Data struct {
		AccountId       string `json:"accountId"`
		Symbol          string `json:"symbol"`
		Headquarters    string `json:"headquarters"`
		Credits         int    `json:"credits"`
		StartingFaction string `json:"startingFaction"`
		ShipCount       int    `json:"shipCount"`
	} `json:"data"`
}

func (c Client) AgentData() (*MyAgentResponse, error) {
	res := &MyAgentResponse{}

	url := url + "/my/agent"
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetResult(res).
		SetError(&errorResponse{}).
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return nil, errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return res, nil
}

type ShipNavResponse struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          struct {
		Destination struct {
			Symbol       string `json:"symbol"`
			Type         string `json:"type"`
			SystemSymbol string `json:"systemSymbol"`
			X            int    `json:"x"`
			Y            int    `json:"y"`
		} `json:"destination"`
		Departure struct {
			Symbol       string `json:"symbol"`
			Type         string `json:"type"`
			SystemSymbol string `json:"systemSymbol"`
			X            int    `json:"x"`
			Y            int    `json:"y"`
		} `json:"departure"`
		Origin struct {
			Symbol       string `json:"symbol"`
			Type         string `json:"type"`
			SystemSymbol string `json:"systemSymbol"`
			X            int    `json:"x"`
			Y            int    `json:"y"`
		} `json:"origin"`
		DepartureTime time.Time `json:"departureTime"`
		Arrival       time.Time `json:"arrival"`
	} `json:"route"`
	Status     string `json:"status"`
	FlightMode string `json:"flightMode"`
}

type ListMyShipsResponse struct {
	Data []ShipResponse `json:"data"`
	Meta struct {
		Total int `json:"total"`
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}

type GetShipResponse struct {
	Data ShipResponse `json:"data"`
}

type ShipResponse struct {
	Symbol       string `json:"symbol"`
	Registration struct {
		Name          string `json:"name"`
		FactionSymbol string `json:"factionSymbol"`
		Role          string `json:"role"`
	} `json:"registration"`
	Nav  ShipNavResponse `json:"nav"`
	Crew struct {
		Current  int    `json:"current"`
		Required int    `json:"required"`
		Capacity int    `json:"capacity"`
		Rotation string `json:"rotation"`
		Morale   int    `json:"morale"`
		Wages    int    `json:"wages"`
	} `json:"crew"`
	Frame struct {
		Symbol         string `json:"symbol"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		Condition      int    `json:"condition"`
		ModuleSlots    int    `json:"moduleSlots"`
		MountingPoints int    `json:"mountingPoints"`
		FuelCapacity   int    `json:"fuelCapacity"`
		Requirements   struct {
			Power int `json:"power"`
			Crew  int `json:"crew"`
			Slots int `json:"slots"`
		} `json:"requirements"`
	} `json:"frame"`
	Reactor struct {
		Symbol       string `json:"symbol"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Condition    int    `json:"condition"`
		PowerOutput  int    `json:"powerOutput"`
		Requirements struct {
			Power int `json:"power"`
			Crew  int `json:"crew"`
			Slots int `json:"slots"`
		} `json:"requirements"`
	} `json:"reactor"`
	Engine struct {
		Symbol       string `json:"symbol"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Condition    int    `json:"condition"`
		Speed        int    `json:"speed"`
		Requirements struct {
			Power int `json:"power"`
			Crew  int `json:"crew"`
			Slots int `json:"slots"`
		} `json:"requirements"`
	} `json:"engine"`
	Cooldown struct {
		ShipSymbol       string    `json:"shipSymbol"`
		TotalSeconds     int       `json:"totalSeconds"`
		RemainingSeconds int       `json:"remainingSeconds"`
		Expiration       time.Time `json:"expiration"`
	} `json:"cooldown"`
	Modules []struct {
		Symbol       string `json:"symbol"`
		Capacity     int    `json:"capacity"`
		Range        int    `json:"range"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Requirements struct {
			Power int `json:"power"`
			Crew  int `json:"crew"`
			Slots int `json:"slots"`
		} `json:"requirements"`
	} `json:"modules"`
	Mounts []struct {
		Symbol       string   `json:"symbol"`
		Name         string   `json:"name"`
		Description  string   `json:"description"`
		Strength     int      `json:"strength"`
		Deposits     []string `json:"deposits"`
		Requirements struct {
			Power int `json:"power"`
			Crew  int `json:"crew"`
			Slots int `json:"slots"`
		} `json:"requirements"`
	} `json:"mounts"`
	Cargo struct {
		Capacity  int `json:"capacity"`
		Units     int `json:"units"`
		Inventory []struct {
			Symbol      string `json:"symbol"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Units       int    `json:"units"`
		} `json:"inventory"`
	} `json:"cargo"`
	Fuel struct {
		Current  int `json:"current"`
		Capacity int `json:"capacity"`
		Consumed struct {
			Amount    int       `json:"amount"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"consumed"`
	} `json:"fuel"`
}

func (c Client) ListMyShips() (*ListMyShipsResponse, error) {
	res := &ListMyShipsResponse{}

	url := url + "/my/ships"
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetResult(res).
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return nil, errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return res, nil
}

func (c Client) GetShip(symbol string) (*GetShipResponse, error) {
	res := &GetShipResponse{}

	url := fmt.Sprintf("%s/my/ships/%s", url, symbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetResult(res).
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return nil, errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return res, nil
}

func (c Client) PurchaseShip(shipType, shipyardWaypointSymbol string) error {
	req := map[string]any{"shipType": shipType, "waypointSymbol": shipyardWaypointSymbol}

	url := url + "/my/ships"
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetBody(req).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return nil
}

type OrbitResponse struct {
	Data struct {
		Nav ShipNavResponse `json:"nav"`
	} `json:"data"`
}

func (c Client) OrbitShip(miningShipSymbol string) (*OrbitResponse, error) {
	res := &OrbitResponse{}

	url := fmt.Sprintf("%s/my/ships/%s/orbit", url, miningShipSymbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetResult(res).
		Post(url)
	if err != nil {
		return nil, errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return nil, errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return res, nil
}

func (c Client) FlyShip(shipSymbol, waypointSymbol string) error {
	req := map[string]any{"waypointSymbol": waypointSymbol}

	url := fmt.Sprintf("%s/my/ships/%s/navigate", url, shipSymbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetBody(req).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return nil
}

func (c Client) DockShip(shipSymbol string) error {
	url := fmt.Sprintf("%s/my/ships/%s/dock", url, shipSymbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return nil
}

func (c Client) RefuelShip(shipSymbol string) error {
	url := fmt.Sprintf("%s/my/ships/%s/refuel", url, shipSymbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return nil
}

type ExtractResourcesResponse struct {
	Data struct {
		Cooldown struct {
			ShipSymbol       string    `json:"shipSymbol"`
			TotalSeconds     int       `json:"totalSeconds"`
			RemainingSeconds int       `json:"remainingSeconds"`
			Expiration       time.Time `json:"expiration"`
		} `json:"cooldown"`
		Extraction struct {
			ShipSymbol string `json:"shipSymbol"`
			Yield      struct {
				Symbol string `json:"symbol"`
				Units  int    `json:"units"`
			} `json:"yield"`
		} `json:"extraction"`
		Cargo struct {
			Capacity  int `json:"capacity"`
			Units     int `json:"units"`
			Inventory []struct {
				Symbol      string `json:"symbol"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Units       int    `json:"units"`
			} `json:"inventory"`
		} `json:"cargo"`
	} `json:"data"`
}

func (c Client) ExtractResources(shipSymbol string) (*ExtractResourcesResponse, error) {
	res := &ExtractResourcesResponse{}

	url := fmt.Sprintf("%s/my/ships/%s/extract", url, shipSymbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetResult(res).
		Post(url)
	if err != nil {
		return nil, errors.Wrap(err, "making request")
	}
	if resp.Error() != nil {
		return nil, errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return res, nil
}

type MyShipsCargoResponse struct {
	Data struct {
		Capacity  int `json:"capacity"`
		Units     int `json:"units"`
		Inventory []struct {
			Symbol      string `json:"symbol"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Units       int    `json:"units"`
		} `json:"inventory"`
	} `json:"data"`
}

func (mscr MyShipsCargoResponse) Inventory() []string {
	var res []string
	for _, inv := range mscr.Data.Inventory {
		res = append(res, fmt.Sprintf("%d %s", inv.Units, inv.Symbol))
	}
	return res
}

func (c Client) ShipCargo(shipSymbol string) (*MyShipsCargoResponse, error) {
	res := &MyShipsCargoResponse{}

	url := fmt.Sprintf("%s/my/ships/%s/cargo", url, shipSymbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetResult(res).
		Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		traceResponse(resp, err)
	}

	return res, nil
}

type SellShipCargoResponse struct {
	Data struct {
		Agent struct {
			AccountID       string `json:"accountId"`
			Symbol          string `json:"symbol"`
			Headquarters    string `json:"headquarters"`
			Credits         int    `json:"credits"`
			StartingFaction string `json:"startingFaction"`
			ShipCount       int    `json:"shipCount"`
		} `json:"agent"`
		Cargo struct {
			Capacity  int `json:"capacity"`
			Units     int `json:"units"`
			Inventory []struct {
				Symbol      string `json:"symbol"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Units       int    `json:"units"`
			} `json:"inventory"`
		} `json:"cargo"`
		Transaction struct {
			WaypointSymbol string    `json:"waypointSymbol"`
			ShipSymbol     string    `json:"shipSymbol"`
			TradeSymbol    string    `json:"tradeSymbol"`
			Type           string    `json:"type"`
			Units          int       `json:"units"`
			PricePerUnit   int       `json:"pricePerUnit"`
			TotalPrice     int       `json:"totalPrice"`
			Timestamp      time.Time `json:"timestamp"`
		} `json:"transaction"`
	} `json:"data"`
}

func (c Client) SellShipCargo(shipSymbol, cargoSymbol string, cargoUnits int) (*SellShipCargoResponse, error) {
	res := &SellShipCargoResponse{}

	req := map[string]any{"symbol": cargoSymbol, "units": cargoUnits}

	url := fmt.Sprintf("%s/my/ships/%s/sell", url, shipSymbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetBody(req).
		SetResult(res).
		Post(url)
	if err != nil {
		return nil, errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return nil, errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return res, nil
}

func (c Client) JettisonShipCargo(shipSymbol, cargoSymbol string, cargoUnits int) error {
	req := map[string]any{"symbol": cargoSymbol, "units": cargoUnits}

	url := fmt.Sprintf("%s/my/ships/%s/jettison", url, shipSymbol)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetError(&errorResponse{}).
		SetBody(req).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "making request")
	}

	if resp.Error() != nil {
		return errors.New(resp.Error().(*errorResponse).Error.Message)
	}

	return nil
}
