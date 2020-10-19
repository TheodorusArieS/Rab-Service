package app

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"rab-service-test2/datasource/mysql/rab_db"
)

var (
	router = gin.Default()
)

func Start(){
	UrlMap()
	rab_db.PrintMe()
	router.Run()
}