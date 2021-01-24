package application

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_domain/repository"
	"net/http"
)

//https://github.com/gin-gonic/gin/issues/339#issuecomment-111694462

type CashierService struct {
	CommonService       *CommonService
	ServerApiRepository repository.ServerApiRepository
}

func NewCashierService(commonService *CommonService,
	serverApiRepository repository.ServerApiRepository) *CashierService {
	return &CashierService{CommonService: commonService,
		ServerApiRepository: serverApiRepository,
	}
}

func (s *CashierService) Get(c *gin.Context) {
	/*if !s.IsLogin(c) {
		log.Print("host %q", c.Request.URL.Host)
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s?origin=%s", s.LoginUrl, c.Request.URL))
		return
	}*/

	restaurantId, restIdErr := s.CommonService.GetRestaurantId(c)
	if restIdErr != nil {
		c.Redirect(http.StatusMovedPermanently, "/portal/restaurant")
	}

	resource := s.CommonService.GetCommonResource()
	resource.IsCashier = true
	resource.PageTitle = "Cashier"

	areas, err := s.ServerApiRepository.GetAreasByRestaurantId(restaurantId)
	if err != nil {
		resource.ErrorMessage = err.ErrorMessage
	} else {
		resource.Areas = areas
	}

	c.HTML(http.StatusOK, "cashier", gin.H{
		"resource": resource,
	})
	return
}

func (s *CashierService) Post(c *gin.Context) {
	/*if !s.IsLogin(c) {
		c.JSON(http.StatusUnauthorized, exception.CreateError(exception.CodeSignatureInvalid, "Access denied."))
		return
	}*/

	//	c.JSON(http.StatusOK, gin.H{"filePath": c.Request.URL.String() + "/view/" + directory + "/" + baseFileName, "message": fmt.Sprintf("'%s' uploaded!", file.Filename)})
}
