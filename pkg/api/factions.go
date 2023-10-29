package api

import (
	"os"

	"github.com/pkg/errors"
)

type ListFactionsResponse struct {
	Data []struct {
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
	} `json:"data"`
	Meta struct {
		Total int `json:"total"`
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"meta"`
}

func (lfr ListFactionsResponse) Traits(index int) []string {
	var res []string
	for _, trait := range lfr.Data[index].Traits {
		res = append(res, trait.Name)
	}
	return res
}

func (c Client) ListFactions() (*ListFactionsResponse, error) {
	res := &ListFactionsResponse{}
	url := url + "/factions"
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
