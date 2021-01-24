package persistence

import (
	"encoding/json"
	"fmt"
	"github.com/hoanganf/pos_domain/entity/exception"
	"github.com/hoanganf/pos_domain/repository"
	"github.com/hoanganf/pos_portal/src/infrastructure/client"
	"net/http"
)

type LocalApiRepositoryImpl struct {
	Client *client.LocalApiClient
}

func NewLocalApiRepository(client *client.LocalApiClient) repository.LocalApiRepository {
	return &LocalApiRepositoryImpl{Client: client}
}

func (r *LocalApiRepositoryImpl) GetOrdersByAreaId(areaId int64) ([]interface{}, *exception.Error) {

	resp, err := r.Client.GetOrdersByAreaId(areaId)
	if err != nil {
		return nil, exception.CreateError(exception.CodeUnknown, err.Error())
	}

	user, err := r.GetResponse(resp)
	if err != nil {
		return nil, exception.CreateError(exception.CodeUnknown, err.Error())
	}

	return user, nil

}

func (s *LocalApiRepositoryImpl) GetResponse(resp *http.Response) ([]interface{}, error) {
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
