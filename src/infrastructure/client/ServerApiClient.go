package client

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type ServerApiClient struct {
	Host                 string
	RestaurantEntryPoint string
	Client               *http.Client
}

func NewServerApiClient(host string, timeout int) *ServerApiClient {
	return &ServerApiClient{
		Host:                 host,
		RestaurantEntryPoint: "/v1/restaurants",
		Client: &http.Client{
			Timeout: time.Millisecond * time.Duration(timeout)},
	}
}

func (c *ServerApiClient) GetAreasRequest(restId int64) (*http.Response, error) {
	url := c.Host + c.RestaurantEntryPoint + "/" + strconv.FormatInt(restId, 10) + "/areas?fields=name,id"
	return c.DoGetRequest(url)
}

func (c *ServerApiClient) GetRestaurantsRequest() (*http.Response, error) {
	url := c.Host + c.RestaurantEntryPoint + "?fields=name,id"
	return c.DoGetRequest(url)
}

func (c *ServerApiClient) DoGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("can not create http request. [%w]", err)
	}
	// Send request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client has error. [%w]", err)
	}

	return resp, nil
}
