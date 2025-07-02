package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IrusHunter/MicroserviceCalculator/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) Calculate(ctx context.Context, a, b float32, operation string) (*types.ResultResponce, error) {
	endpoint := fmt.Sprintf("%s?a=%f&b=%f&operation=%s", c.endpoint, a, b, operation)
	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service responsed with non OK status code: %s", httpErr["error"])
	}

	resultResponce := new(types.ResultResponce)
	if err := json.NewDecoder(resp.Body).Decode(resultResponce); err != nil {
		return nil, err
	}

	return resultResponce, nil
}
