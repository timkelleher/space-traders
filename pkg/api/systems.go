package api

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

type ListWaypointsResponse struct {
	Data []struct {
		SystemSymbol string `json:"systemSymbol"`
		Symbol       string `json:"symbol"`
		Type         string `json:"type"`
		X            int    `json:"x"`
		Y            int    `json:"y"`
		Orbitals     []any  `json:"orbitals"`
		Traits       []struct {
			Symbol      string `json:"symbol"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"traits"`
		Modifiers []any `json:"modifiers"`
		Chart     struct {
			SubmittedBy string    `json:"submittedBy"`
			SubmittedOn time.Time `json:"submittedOn"`
		} `json:"chart"`
		Faction struct {
			Symbol string `json:"symbol"`
		} `json:"faction"`
		Orbits              string `json:"orbits"`
		IsUnderConstruction bool   `json:"isUnderConstruction"`
	} `json:"data"`
	Meta struct {
		Total int `json:"total"`
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}

func (lwr ListWaypointsResponse) Traits(index int) []string {
	var res []string
	for _, trait := range lwr.Data[index].Traits {
		res = append(res, trait.Name)
	}
	return res
}

func (c Client) ListWaypointsByType(systemSymbol, kind string) (*ListWaypointsResponse, error) {
	res := &ListWaypointsResponse{}
	url := fmt.Sprintf("%s/systems/%s/waypoints?type=%s", url, systemSymbol, kind)
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
		traceResponse(resp, err)
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
		traceResponse(resp, err)
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
		traceResponse(resp, err)
	}

	return res, nil
}
