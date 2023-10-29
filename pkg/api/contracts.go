package api

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type MyContractsResponse struct {
	Data []struct {
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
		Meta             struct {
			Total int `json:"total"`
			Page  int `json:"page"`
			Limit int `json:"limit"`
		} `json:"meta"`
	} `json:"data"`
}

func (mcr MyContractsResponse) TermsDelieveries(index int) []string {
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
		traceResponse(resp, err)
	}

	return nil
}
