package api

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type ContractResponse struct {
	ID            string `json:"id"`
	FactionSymbol string `json:"factionSymbol"`
	Type          string `json:"type"`
	Terms         struct {
		Deadline string `json:"deadline"`
		Payment  struct {
			OnAccepted  int `json:"onAccepted"`
			OnFulfilled int `josn:"onFulfilled"`
		} `json:"payment"`
		Deliver []struct {
			TradeSymbol       string `json:"tradeSymbol"`
			DestinationSymbol string `json:"destinationSymbol"`
			UnitsRequired     int    `json:"unitsRequired"`
			UnitsFulfilled    int    `json:"unitsFulfilled"`
		} `josn:"deliver"`
	} `json:"terms"`
	Accepted         bool   `json:"accepted"`
	Fulfilled        bool   `json:"fulfilled"`
	Expiration       string `json:"expiration"`
	DeadlineToAccept string `json:"deadlineToAccept"`
}

type MyContractsResponse struct {
	Data []ContractResponse `json:"data"`
	Meta struct {
		Total int `json:"total"`
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}

type GetContractResponse struct {
	Data ContractResponse `json:"data"`
}

func (mcr MyContractsResponse) Delieveries(index int) []string {
	var res []string
	for _, deliver := range mcr.Data[index].Terms.Deliver {
		res = append(res, fmt.Sprintf("%d/%d %s to %s",
			deliver.UnitsFulfilled,
			deliver.UnitsRequired,
			deliver.TradeSymbol,
			deliver.DestinationSymbol,
		))
	}
	return res
}

func (c Client) ListContracts() (*MyContractsResponse, error) {
	res := &MyContractsResponse{}
	url := url + "/my/contracts"
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

func (c Client) GetContract(contractId string) (*GetContractResponse, error) {
	res := &GetContractResponse{}
	url := fmt.Sprintf("%s/my/contracts/%s", url, contractId)
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

func (c Client) AcceptContract(id string) error {
	url := fmt.Sprintf("%s/my/contracts/%s/accept", url, id)
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

func (c Client) DeliverContract(contractId, shipSymbol, tradeSymbol string, units int) error {
	req := map[string]any{"shipSymbol": shipSymbol, "tradeSymbol": tradeSymbol, "units": strconv.Itoa(units)}

	url := fmt.Sprintf("%s/my/contracts/%s/deliver", url, contractId)
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
		traceResponse(resp, err)
	}

	return nil
}

func (c Client) FulfillContract(contractId, shipSymbol string) error {
	url := fmt.Sprintf("%s/my/contracts/%s/fulfill", url, contractId)
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
		traceResponse(resp, err)
	}

	return nil
}
