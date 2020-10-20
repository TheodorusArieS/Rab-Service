package controller

import (
	// "fmt"
	"net/http"
	rab "rab-service-test2/domain"
	"rab-service-test2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRabData(c *gin.Context) {
	var r rab.RabDataList
	_ = c.ShouldBindJSON(&r)

	result, err := service.RabService.CreateRabData(r)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, result)
	}

}

func GetRabData(c *gin.Context) {
	result, err := service.RabService.GetRabData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, result)
	}

}

func GetRabDetails(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := service.RabService.GetRabDetails(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func CreateRabList(c *gin.Context) {
	var rab rab.RabList
	_ = c.ShouldBindJSON(&rab)
	result, err := service.RabService.CreateRabList(rab)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func GetRabList(c *gin.Context){
	result,err :=service.RabService.GetRabList()
	if err !=nil{
		c.JSON(http.StatusBadRequest,err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
