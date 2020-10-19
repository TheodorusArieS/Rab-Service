package rab

type RestResponse struct {
	Status int64 `json:"status"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

type RabList struct {
	Rab_Id int64 `json:"rab_id"`
	RabName string `json:"rab_name"`
	Comodity string `json:"comodity"`
	Province string `json:"province"`
	City string `json:"city"`
	Quantity int64 `json:"quantity"`

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