package controller

import (
	"github.com/gin-gonic/gin"
	"rab-service-test2/domain"
	"rab-service-test2/service"
	"net/http"
)

func CreateRabData(c *gin.Context) {
	var r rab.RabDataList
	_ = c.ShouldBindJSON(&r)

	result, err := service.RabService.CreateRabData(r)

	if err != nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK,result)
}

func GetRabData(c *gin.Context){
	result,err := service.RabService.GetRabData()
	if err !=nil{
		c.JSON(http.StatusInternalServerError,err)
	}
	c.JSON(http.StatusOK,result)
}