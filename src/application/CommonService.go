package application

import (
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_domain/repository"
	"github.com/hoanganf/pos_portal/src/application/resource"
	"log"
	"strconv"
)

const (
	ORIGIN = "origin"
)

type CommonService struct {
	LoginUrl        string
	ImageUrl        string
	DesignUrl       string
	TokenName       string
	LoginRepository repository.LoginRepository
}

func NewCommonService(loginUrl string, designUrl string, imageUrl string, tokenName string, loginRepository repository.LoginRepository) *CommonService {

	return &CommonService{
		LoginUrl:        loginUrl,
		DesignUrl:       designUrl,
		ImageUrl:        imageUrl,
		TokenName:       tokenName,
		LoginRepository: loginRepository}
}

func (s *CommonService) GetCommonResource() *resource.PageResource {
	pageResource := &resource.PageResource{}
	pageResource.ImageUrl = s.ImageUrl
	pageResource.DesignUrl = s.DesignUrl
	pageResource.LoginUrl = s.LoginUrl
	return pageResource
}

func (s *CommonService) GetRestaurantId(c *gin.Context) (int64, error) {
	restaurantId, err := c.Cookie("restaurantId")
	if err != nil {
		log.Print(err)
		return -1, err
	}

	restId, err := strconv.ParseInt(restaurantId, 0, 64)
	if err != nil {
		log.Print(err)
		return -1, err
	}
	return restId, nil
}

func (s *CommonService) IsLogin(c *gin.Context) bool {
	cookie, err := c.Cookie(s.TokenName)
	if err != nil {
		log.Print(err)
		return false
	}
	user, cErr := s.LoginRepository.GetUserByJwt(cookie)
	if cErr != nil {
		log.Print(cErr.ErrorMessage)
		return false
	}

	if user.JWT != "" {
		return true
	}
	return false
}
