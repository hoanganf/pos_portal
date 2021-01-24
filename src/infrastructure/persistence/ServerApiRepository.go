package persistence

import (
	"encoding/json"
	"fmt"
	"github.com/hoanganf/pos_domain/entity/exception"
	"github.com/hoanganf/pos_domain/repository"
	"github.com/hoanganf/pos_portal/src/infrastructure/client"
	"net/http"
)

type ServerApiRepositoryImpl struct {
	Client *client.ServerApiClient
}

func NewServerApiRepository(client *client.ServerApiClient) repository.ServerApiRepository {
	return &ServerApiRepositoryImpl{Client: client}
}

func (r *ServerApiRepositoryImpl) GetAreasByRestaurantId(restId int64) ([]interface{}, *exception.Error) {

	resp, err := r.Client.GetAreasRequest(restId)
	if err != nil {
		return nil, exception.CreateError(exception.CodeUnknown, err.Error())
	}

	user, err := r.GetResponse(resp)
	if err != nil {
		return nil, exception.CreateError(exception.CodeUnknown, err.Error())
	}

	return user, nil

}

func (r *ServerApiRepositoryImpl) GetRestaurants() ([]interface{}, *exception.Error) {

	resp, err := r.Client.GetRestaurantsRequest()
	if err != nil {
		return nil, exception.CreateError(exception.CodeUnknown, err.Error())
	}

	result, err := r.GetResponse(resp)
	if err != nil {
		return nil, exception.CreateError(exception.CodeUnknown, err.Error())
	}

	return result, nil

}

func (s *ServerApiRepositoryImpl) GetResponse(resp *http.Response) ([]interface{}, error) {
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResponse exception.Error
		err := json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return nil, fmt.Errorf("can not decode login response. [%w]", err)
		}
		if errResponse.ErrorCode == exception.CodeNotFound {
			return nil, fmt.Errorf("Not found.")
		}
		if errResponse.ErrorCode == exception.CodeValueInvalid {
			return nil, fmt.Errorf("request value invalid.")
		}
		return nil, fmt.Errorf("api has error. [%s]", errResponse.ErrorMessage)
	}

	var items []interface{}
	err := json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return nil, fmt.Errorf("can not decode response. [%w]", err)
	}
	return items, nil
}
