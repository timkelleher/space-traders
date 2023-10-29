package api

import (
	"os"

	"github.com/pkg/errors"
)

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

func (c client) MyAgent() (*MyAgentResponse, error) {
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
		traceResponse(resp, err)
	}

	return res, nil
}
