package app

import (

	"rab-service-test2/controller"
)

func UrlMap(){

	rab :=router.Group("/rab")
	{
		rab.POST("/",controller.CreateRabList)
		rab.GET("/:id",controller.GetRabDetails)
		rab.GET("/",controller.GetRabList)
	}
	data :=router.Group("/data")
	{
		data.POST("/",controller.CreateRabData)
		data.GET("/",controller.GetRabData)
	}
}