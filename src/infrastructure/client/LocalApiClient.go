package client

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type LocalApiClient struct {
	Host                 string
	AreaEntryPoint string
	Client               *http.Client
}

func NewLocalApiClient(host string, timeout int) *LocalApiClient {
	return &LocalApiClient{
		Host:                 host,
		AreaEntryPoint: "/v1/areas",
		Client: &http.Client{
			Timeout: time.Millisecond * time.Duration(timeout)},
	}
}

func (c *LocalApiClient) GetOrdersByAreaId(areaId int64) (*http.Response, error) {
	url := c.Host + c.AreaEntryPoint + "/" + strconv.FormatInt(areaId, 10) + "/orders?fields=numberId,tableName,countSum,priceSum"
	return c.DoGetRequest(url)
}

func (c *LocalApiClient) DoGetRequest(url string) (*http.Response, error) {
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
