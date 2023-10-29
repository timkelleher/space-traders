package api

import (
	"github.com/pkg/errors"
)

func (c client) Register(callSign, faction string) error {
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
