package schema

import (
	"time"
	"github.com/jinzhu/gorm"
)

type TableInterface interface{
	Pk() string
	Ref() string
	AddForeignKey()
	InsertDefaults()
}

var (
	Database *gorm.DB
)

type Base struct {
	Id int `gorm:"primary_key"`
	CreateAt time.Time `gorm:"column:created_at;not_null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"colum:updated_at;not_null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP`
	CreatedBy string `gorm:"default:null"`
	UpdatedBy string `gorm:"default:null"`
	DeletedAt *time.Time `sql:"index"`
}

func AutoMigrate(database *gorm.DB){
	Database = database
	database.DropTableIfExists()

	database.AutoMigrate(
		RabDataList{},
	)
}