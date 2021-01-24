package application

import (
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_domain/repository"
	"log"
	"net/http"
	"strconv"
)

type RestaurantService struct {
	CommonService       *CommonService
	ServerApiRepository repository.ServerApiRepository
}

func NewRestaurantService(commonService *CommonService,
	serverApiRepository repository.ServerApiRepository) *RestaurantService {
	return &RestaurantService{CommonService: commonService,
		ServerApiRepository: serverApiRepository}
}

func (s *RestaurantService) Get(c *gin.Context) {
	/*if !s.IsLogin(c) {
		log.Print("host %q", c.Request.URL.Host)
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s?origin=%s", s.LoginUrl, c.Request.URL))
		return
	}*/

	resource := s.CommonService.GetCommonResource()
	resource.Origin = c.DefaultQuery(ORIGIN, "cashier")

	results, err := s.ServerApiRepository.GetRestaurants()
	if err != nil {
		resource.ErrorMessage = err.ErrorMessage
	} else {
		resource.Restaurants = results
	}

	c.HTML(http.StatusOK, "restaurant", gin.H{
		"resource": resource,
	})
	return
}

func (s *RestaurantService) Post(c *gin.Context) {
	/*if !s.IsLogin(c) {
		c.JSON(http.StatusUnauthorized, exception.CreateError(exception.CodeSignatureInvalid, "Access denied."))
		return
	}*/
	restaurantIdString := c.PostForm("restaurant_id")
	origin := c.DefaultPostForm(ORIGIN, "cashier")
	_, err := strconv.ParseInt(restaurantIdString, 0, 64)
	if err != nil {
		log.Print(err)
		resource := s.CommonService.GetCommonResource()
		resource.ErrorMessage = "restaurantId not found."
		resource.Origin = origin
		results, err := s.ServerApiRepository.GetRestaurants()
		if err != nil {
			resource.ErrorMessage = resource.ErrorMessage + " AND " + err.ErrorMessage
		} else {
			resource.Restaurants = results
		}
		c.HTML(http.StatusOK, "restaurant", gin.H{
			"resource": resource,
		})
		return
	}

	c.SetCookie("restaurantId", restaurantIdString, 0, "/", "", http.SameSiteNoneMode, false, true)
	c.Redirect(http.StatusMovedPermanently, "/portal/"+origin)
	//	c.JSON(http.StatusOK, gin.H{"filePath": c.Request.URL.String() + "/view/" + directory + "/" + baseFileName, "message": fmt.Sprintf("'%s' uploaded!", file.Filename)})
}
