package query

const (
	QueryInsertRabDataList =`
	INSERT INTO rab_data_list(
		product_name,
		unit_product,
		quantity,
		unit_price,
		total_price)
	VALUES(?,?,?,?,quantity*unit_price)
	`
	QueryGetRabDataList =`
	SELECT product_name,unit_product,quantity,unit_price,total_price FROM rab_service.rab_data_list;
	`
)