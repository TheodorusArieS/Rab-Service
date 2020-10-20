package schema

type RabLog struct {
	Base
	RabId int64 `gorm:"not_null,column:rab_id" json:"rab_id"`
	Notes string `gorm:"not_null,column:notes" json:"notes"` 
}

func (RabLog) TableName() string {
	return "rab_log"
}
func (RabLog) Pk() string{
	return "id"
}

func (f RabLog) Ref() string{
	return f.TableName() +"(" + f.Pk() + ")"
}

func (f RabLog) AddForeignKeys(){

}
func (f RabLog) InsertDefaults(){
	
}