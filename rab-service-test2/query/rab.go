package query

const (
	QueryInsertRabDataList = `
	INSERT INTO rab_data_list(
		rab_id,
		product_name,
		unit_product,
		quantity,
		unit_price,
		total_price)
	VALUES(?,?,?,?,?,quantity*unit_price)
	`
	QueryGetRabDataList = `
	SELECT rab_id,product_name,unit_product,quantity,unit_price,total_price FROM rab_service.rab_data_list
	`

	QueryCreateRabList = `
	INSERT INTO rab_data(rab_name,comodity,province,city)
	VALUES(?,?,?,?)
	`

	QueryGetRabListCreatedData =`
	SELECT rab_name,comodity,province,city
	FROM rab_data
	WHERE id = ?
	`

	QueryGetRabList =`
	SELECT id,rab_name,comodity,province,city
	FROM rab_data
	`

	QueryCreateRabLog =`
	INSERT INTO rab_log(rab_id,notes)
	VALUES(?,?)
	`

	QueryGetRabProductDetails =`
	SELECT rab_id,product_name,unit_product,quantity,unit_price,total_price
	FROM rab_data_list
	WHERE rab_id = ?
	`

	QueryGetRabListDetail =`
	SELECT id,rab_name,comodity,province,city
	FROM rab_data
	WHERE id = ?
	`

	Limit3 =`
	LIMIT 3 OFFSET ?
	`
)
