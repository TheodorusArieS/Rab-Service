package rab_db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"rab-service-test2/schema"
)

var (
	Client *sql.DB
	Database *gorm.DB
)

func PrintMe(){
	fmt.Println("Print Me")
}

func init(){
	username :="root"
	password :="admin"
	host :="127.0.0.1:3306"
	dbSchema :="rab_service"

	dataSource:=fmt.Sprintf("%s:%s@tcp(%s)/%s",username,password,host,dbSchema)
	Client, err := sql.Open("mysql",dataSource)
	if err !=nil{
		panic(err)
	}
	if err =Client.Ping();err !=nil{
		panic(err)
	}
	fmt.Println("Database configured")


	db, err := gorm.Open("mysql",dataSource)
	if err != nil{
		fmt.Println("mySQL:",err)
	}
	fmt.Println("Dropping All Tables")
	Database = db
	schema.AutoMigrate(db)
	fmt.Println("Recreating All Table")
}