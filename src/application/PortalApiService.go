package application

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_domain/entity/exception"
	"github.com/hoanganf/pos_domain/repository"
	"log"
	"net/http"
	"strconv"
)

//https://github.com/gin-gonic/gin/issues/339#issuecomment-111694462

type PortalApiService struct {
	CommonService      *CommonService
	LocalApiRepository repository.LocalApiRepository
}

func NewPortalApiService(commonService *CommonService,
	localApiRepository repository.LocalApiRepository) *PortalApiService {
	return &PortalApiService{CommonService: commonService,
		LocalApiRepository: localApiRepository,
	}
}

func (s *PortalApiService) GetOrder(c *gin.Context) {
	/*if !s.CommonService.IsLogin(c) {
		c.JSON(http.StatusUnauthorized, exception.CreateError(exception.CodeSignatureInvalid, "Access denied."))
		return
	}*/
	areaId, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		log.Print(paramErr)
		c.JSON(http.StatusBadRequest, exception.CreateError(exception.CodeValueInvalid, "areaId invalid."))
		return
	}

	var orders, err = s.LocalApiRepository.GetOrdersByAreaId(areaId)
	if err != nil || len(orders) == 0 {
		log.Print(err)
		c.JSON(http.StatusNotFound, exception.CreateError(exception.CodeNotFound, "order not found."))
		return
	}
	c.JSON(http.StatusOK, orders)
}
