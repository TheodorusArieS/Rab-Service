package rab

type RestResponse struct {
	Status int64 `json:"status"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

type RabList struct {
	Id int64 `json:"id"`
	RabName string `json:"rab_name"`
	Comodity string `json:"comodity"`
	Province string `json:"province"`
	City string `json:"city"`
	DetailProduct []RabDataList `json:"detail_product,omitempty"`
	TotalPrice int64 `json:"total_price"`

}

type RabDataList struct {
	Rab_Id int64 `json:"rab_id"`
	ProductName string `json:"product_name"`
	UnitProduct string `json:"unit_product"`
	Quantity int64 `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	TotalPrice float32 `json:"total_price"`
}

type CreateRab struct {
	RabName string `json:"rab_name"`
	Comodity string `json:"comodity"`
	Province string `json:"province"`
	City string `json:"city"`
	User string `json:"user"`
	DetailProduct []RabDataList `json:"detail_product,omitempty`
}

type RabLogList struct{
	RabId int64 `json:"rab_id"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	Notes string `json:"notes"`
}