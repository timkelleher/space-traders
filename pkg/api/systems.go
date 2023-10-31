package api

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

type SystemResponse struct {
	Symbol       string `json:"symbol"`
	SectorSymbol string `json:"sectorSymbol"`
	Type         string `json:"type"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
	Waypoints    []struct {
		Symbol   string `json:"symbol"`
		Type     string `json:"type"`
		X        int    `json:"x"`
		Y        int    `json:"y"`
		Orbitals []struct {
			Symbol string `json:"symbol"`
		} `json:"orbitals"`
		Orbits string `json:"orbits"`
	} `json:"waypoints"`
	Factions []struct {
		Symbol string `json:"symbol"`
	} `json:"factions"`
}

func (sr SystemResponse) GetFactions() []string {
	var res []string
	for _, faction := range sr.Factions {
		res = append(res, faction.Symbol)
	}
	return res
}

func (sr SystemResponse) GetWaypoints() []string {
	var res []string
	for _, waypoint := range sr.Waypoints {
		res = append(res, fmt.Sprintf("%s - %s - (%d,%d)",
			waypoint.Symbol,
			waypoint.Type,
			waypoint.X,
			waypoint.Y,
		))
	}
	return res
}

type ListSystemsResponse struct {
	Data []SystemResponse `json:"data"`
	Meta struct {
		Total int `json:"total"`
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}

func (c Client) ListSystems(page string) (*ListSystemsResponse, error) {
	res := &ListSystemsResponse{}

	params := map[string]string{"page": page}
	url := fmt.Sprintf("%s/systems", url)
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetAuthToken(os.Getenv("SPACE_TRADERS_TOKEN")).
		SetResult(res).
		SetQueryParams(params).
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

type GetSystemResponse struct {
	Data SystemResponse `json:"data"`
}

func (c Client) GetSystem(systemSymbol string) (*GetSystemResponse, error) {
	res := &GetSystemResponse{}

	url := fmt.Sprintf("%s/systems/%s", url, systemSymbol)
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

type WaypointResponse struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
	Orbitals     []struct {
		Symbol string `json:"symbol"`
	} `json:"orbitals"`
	Orbits  string `json:"orbits"`
	Faction struct {
		Symbol string `json:"symbol"`
	} `json:"faction"`
	Traits []struct {
		Symbol      string `json:"symbol"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"traits"`
	Modifiers []struct {
		Symbol      string `json:"symbol"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"modifiers"`
	Chart struct {
		WaypointSymbol string    `json:"waypointSymbol"`
		SubmittedBy    string    `json:"submittedBy"`
		SubmittedOn    time.Time `json:"submittedOn"`
	} `json:"chart"`
	IsUnderConstruction bool `json:"isUnderConstruction"`
}

func (wr WaypointResponse) GetTraits() []string {
	var res []string
	for _, trait := range wr.Traits {
		res = append(res, trait.Name)
	}
	return res
}

type ListWaypointsResponse struct {
	Data []WaypointResponse `json:"data"`
	Meta struct {
		Total int `json:"total"`
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}

type GetWaypointResponse struct {
	Data WaypointResponse `json:"data"`
}

func (c Client) GetWaypoint(systemSymbol, waypointSymbol string) (*GetWaypointResponse, error) {
	res := &GetWaypointResponse{}
	url := fmt.Sprintf("%s/systems/%s/waypoints/%s", url, systemSymbol, waypointSymbol)
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

func (c Client) ListWaypointsByType(systemSymbol, waypointType string) (*ListWaypointsResponse, error) {
	res := &ListWaypointsResponse{}
	url := fmt.Sprintf("%s/systems/%s/waypoints?type=%s", url, systemSymbol, waypointType)
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

func (c Client) ListWaypointsByTrait(systemSymbol, trait string) (*ListWaypointsResponse, error) {
	res := &ListWaypointsResponse{}
	url := fmt.Sprintf("%s/systems/%s/waypoints?traits=%s", url, systemSymbol, trait)
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

type ListShipyardShipsResponse struct {
	Data struct {
		Symbol    string `json:"symbol"`
		ShipTypes []struct {
			Type string `json:"type"`
		} `json:"shipTypes"`
		Transactions []struct {
			ShipSymbol     string    `json:"shipSymbol"`
			WaypointSymbol string    `json:"waypointSymbol"`
			AgentSymbol    string    `json:"agentSymbol"`
			Price          int       `json:"price"`
			Timestamp      time.Time `json:"timestamp"`
		} `json:"transactions"`
		Ships []struct {
			Type          string `json:"type"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			Supply        string `json:"supply"`
			Activity      string `json:"activity"`
			PurchasePrice int    `json:"purchasePrice"`
			Frame         struct {
				Symbol         string `json:"symbol"`
				Name           string `json:"name"`
				Description    string `json:"description"`
				ModuleSlots    int    `json:"moduleSlots"`
				MountingPoints int    `json:"mountingPoints"`
				FuelCapacity   int    `json:"fuelCapacity"`
				Requirements   struct {
					Power int `json:"power"`
					Crew  int `json:"crew"`
				} `json:"requirements"`
			} `json:"frame"`
			Reactor struct {
				Symbol       string `json:"symbol"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				PowerOutput  int    `json:"powerOutput"`
				Requirements struct {
					Crew int `json:"crew"`
				} `json:"requirements"`
			} `json:"reactor"`
			Engine struct {
				Symbol       string `json:"symbol"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				Speed        int    `json:"speed"`
				Requirements struct {
					Power int `json:"power"`
					Crew  int `json:"crew"`
				} `json:"requirements"`
			} `json:"engine"`
			Modules []struct {
				Symbol       string `json:"symbol"`
				Name         string `json:"name"`
				Description  string `json:"description"`
				Capacity     int    `json:"capacity"`
				Requirements struct {
					Crew  int `json:"crew"`
					Power int `json:"power"`
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
					Crew  int `json:"crew"`
					Power int `json:"power"`
				} `json:"requirements"`
			} `json:"mounts"`
			Crew struct {
				Required int `json:"required"`
				Capacity int `json:"capacity"`
			} `json:"crew"`
		} `json:"ships"`
		ModificationsFee int `json:"modificationsFee"`
	} `json:"data"`
}

func (lssr ListShipyardShipsResponse) ShipModules(index int) []string {
	var res []string
	for _, module := range lssr.Data.Ships[index].Modules {
		res = append(res, module.Name)
	}
	return res
}

func (c Client) ListShipyardShips(systemSymbol, shipyardWaypointSymbol string) (*ListShipyardShipsResponse, error) {
	res := &ListShipyardShipsResponse{}
	url := fmt.Sprintf("%s/systems/%s/waypoints/%s/shipyard", url, systemSymbol, shipyardWaypointSymbol)
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
