package schema

type RabData struct {
	Base
	RabName string `gorm:"not_null;column:rab_name" json:"rab_name"`
	Comodity string `gorm:"not_null;column:comodity" json:"comodity"`
	Province string `gorm:"not_null;column:province" json:"province"`
	City string `gorm:"not_null;column:city" json:"city"`
}

func (RabData)TableName() string{
	return "rab_data"
}

func (RabData) Pk() string {
	return "id"
}

func (f RabData) Ref() string{
	return f.TableName() +"(" + f.Pk() + ")"
}

func (f RabData) AddForeignKeys(){

}
func (f RabData) InsertDefaults(){
	
}
