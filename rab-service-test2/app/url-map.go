package app

import (
	"fmt"
	"rab-service-test2/controller"
)

func UrlMap(){

	rab :=router.Group("/rab")
	{
		rab.GET("/",controller.GetRabData)
		rab.POST("/",controller.CreateRabData)
	}

	fmt.Println("ada di ulrmap")

}