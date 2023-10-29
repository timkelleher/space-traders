package api

import (
	"os"

	"github.com/pkg/errors"
)

func (c Client) RegisterAgent(callSign, faction string) error {
	req := map[string]interface{}{"symbol": callSign, "faction": faction}

	url := url + "/register"
	resp, err := c.resty.R().
		EnableTrace().
		ForceContentType("application/json").
		SetBody(req).
		SetError(&errorResponse{}).
		Post(url)
	if err != nil {
		return errors.Wrap(err, "making request")
	}

	traceResponse(resp, err)

	if resp.Error() != nil {
		return errors.New(resp.Error().(*errorResponse).Error.Message)
	}
	return nil
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
		traceResponse(resp, err)
	}

	return res, nil
}
