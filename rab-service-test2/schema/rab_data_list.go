package schema

type RabDataList struct{
	Base
	Rab_Id int64 `gorm:"not_null;column:rab_id" json:"rab_id"`
	ProductName string `gorm:"not_null;column:product_name" json:"product_name"`
	UnitProduct string `gorm:"not_null;column:unit_product" json:unit_product`
	Quantity int64 `gorm:"not_null;column:quantity" json:"quantity"`
	UnitPrice float64 `gorm:"unit_price;column:unit_price" json:"unit_price"`
	TotalPrice float32 `gorm:"total_price;column:total_price" json:"total_price"` 
}

func (RabDataList) TableName() string{
	return "rab_data_list"
}

func (RabDataList) Pk() string{
	return "id"
}

func (f RabDataList) Ref() string{
	return f.TableName() +"(" + f.Pk() + ")"
}

func (f RabDataList) AddForeignKeys(){

}
func (f RabDataList) InsertDefaults(){
	
}

