package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_portal/src"
	"log"
)

func main() {
	bean, err := src.InitBean()
	if err != nil {
		log.Fatalln("can not create bean", err)
	}

	router := gin.Default()
	router.HTMLRender = bean.LoadTemplates()
	portal := router.Group("portal")
	{
		portal.GET("/cashier", bean.CashierService.Get)
		portal.GET("/restaurant", bean.RestaurantService.Get)
		portal.POST("/restaurant", bean.RestaurantService.Post)
	}
	v1 := router.Group("v1")
	{
		v1.GET("/areas/:id/orders", bean.PortalApiService.GetOrder)
	}
	router.Run(":8083")
}
