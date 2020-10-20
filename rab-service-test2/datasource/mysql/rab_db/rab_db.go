package rab_db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"rab-service-test2/schema"
	"github.com/joho/godotenv"
	"os"
)

var (
	Client *sql.DB
	Database *gorm.DB
)

func PrintMe(){
	fmt.Println("Print Me")
}

func init(){
	_ = godotenv.Load()


	username :=os.Getenv("DB_USER")
	password :=os.Getenv("DB_PASSWORD")
	host :=os.Getenv("DB_HOST")
	dbSchema :=os.Getenv("DB_SCHEMA")

	dataSource:=fmt.Sprintf("%s:%s@tcp(%s)/%s",username,password,host,dbSchema)
	// Client, err := sql.Open("mysql",dataSource)
	// if err !=nil{
	// 	panic(err)
	// }
	// if err =Client.Ping();err !=nil{
	// 	panic(err)
	// }
	// fmt.Println("Database configured")
	db, err := gorm.Open("mysql",dataSource)
	if err != nil{
		fmt.Println("mySQL:",err)
	}
	fmt.Println("Dropping All Tables")
	Database = db
	Client = db.DB()
	schema.AutoMigrate(db)
	fmt.Println("Recreating All Table")
}