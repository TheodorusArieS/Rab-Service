package rab

import (
	"fmt"
	"math"
	"net/http"
	"rab-service-test2/datasource/mysql/rab_db"

	// "database/sql"
	"rab-service-test2/query"
	errors "rab-service-test2/utilities"
)

func (data *RabDataList) CreateRabDataList(r RabDataList) (*RestResponse, *errors.RestError) {
	stmt, err := rab_db.Client.Prepare(query.QueryInsertRabDataList)
	if err != nil {
		fmt.Println("ERROR WHEN PREPARING CREATE RAB DATA LIST QUERY")

		return nil, errors.BadRequestError("ERROR WHEN PREPARING CREATE RAB DATA LIST QUERY")
	}
	defer stmt.Close()
	result, saveErr := stmt.Exec(r.Rab_Id, r.ProductName, r.UnitProduct, r.Quantity, r.UnitPrice)
	if saveErr != nil {
		return nil, errors.BadRequestError("BAD REQUEST IN CREATING RAB DATA LIST IN EXEC COMMAND")
	}

	id, _ := result.LastInsertId()

	stmt, err = rab_db.Client.Prepare(query.QueryCreateRabLog)
	if err != nil {
		return nil, errors.BadRequestError("Failed to create Log in Create RAB List")
	}

	notes := fmt.Sprintf("Pembuatan Create RAB Data List ID %d", id)
	_, err = stmt.Exec(id, notes)

	if err != nil {
		return nil, errors.BadRequestError("Failed to input RAB List ID to log")
	}
	return &RestResponse{
		Status:  http.StatusOK,
		Data:    result,
		Message: "Sukses input barang",
	}, nil
}

func (data *RabDataList) GetRabDataList(offsetInt int64, search string) (*RestResponse, *errors.RestError) {

	stmt, err := rab_db.Client.Prepare(query.QueryGetRabDataList + query.Limit3)
	if err != nil {
		return nil, errors.BadRequestError("ERROR WHEN PREPARING QUERY IN GET RAB DATA LIST")
	}

	searchKey := fmt.Sprint("%" + search + "%")
	result, err := stmt.Query(searchKey, offsetInt)
	fmt.Println(offsetInt)
	if err != nil {
		return nil, errors.InternalServerError("ERROR IN GET RAB DATA LIST QUERY")
	}
	var results1 []RabDataList
	for result.Next() {
		var res RabDataList
		if err := result.Scan(&res.Rab_Id, &res.ProductName, &res.UnitProduct, &res.Quantity, &res.UnitPrice, &res.TotalPrice); err != nil {
			return nil, errors.BadRequestError("ERROR IN LOOPING FROM DATA IN DATABASE")
		}
		results1 = append(results1, res)

	}
	stmt,err = rab_db.Client.Prepare(query.QueryGetRabDataList)
	if err !=nil{
		return nil,errors.BadRequestError("Error in preparing query for total data")
	}
	rows,err := stmt.Query(searchKey)
	if err != nil{
		return nil,errors.BadRequestError("Error in fetching data from database")
	}

	var totalData float64
	var totalPage float64
	for rows.Next(){
		totalData++
	}
	var pageSize int64 = 3
	totalPage = math.Ceil(totalData/float64(pageSize))
	var currentPage  int64 = offsetInt/pageSize
	if currentPage == 0{
		currentPage =1 
	}

	results := &RestResponse{
		Status:  http.StatusOK,
		Meta:&MetaDetail{
			TotalPage:int64(totalPage),
			CurrentPage:currentPage,
			PageSize:pageSize,
			Total:int64(totalData),
		},
		Data:    results1,
		Message: "Sukses ambil data",
	}
	return results, nil
}

func (data *RabList) GetRabDetails(id int64) (*RestResponse, *errors.RestError) {
	stmt, err := rab_db.Client.Prepare(query.QueryGetRabProductDetails)
	if err != nil {
		return nil, errors.BadRequestError("Error When preparing query for get Rab Details")
	}
	productDetails, err := stmt.Query(id)

	var results []RabDataList
	var totalPrice int64
	for productDetails.Next() {
		var res RabDataList
		if err := productDetails.Scan(&res.Rab_Id, &res.ProductName, &res.UnitProduct, &res.Quantity, &res.UnitPrice, &res.TotalPrice); err != nil {
			return nil, errors.BadRequestError("Error when looping product details")
		}
		results = append(results, res)
		totalPrice += int64(res.TotalPrice)
	}

	stmt, err = rab_db.Client.Prepare(query.QueryGetRabListDetail)
	if err != nil {
		return nil, errors.BadRequestError("Error when preparing query to get RAB List")
	}
	rabList, err := stmt.Query(id)
	if err != nil {
		return nil, errors.BadRequestError("Error when trying to get RAB List from database")
	}
	var rabListDetail RabList
	for rabList.Next() {
		var res RabList
		if err := rabList.Scan(&res.Id, &res.RabName, &res.Comodity, &res.Province, &res.City); err != nil {
			return nil, errors.BadRequestError("Error when trying to implemend RAB List Details")
		}
		rabListDetail = res
	}
	rabListDetail.TotalPrice = totalPrice
	rabListDetail.DetailProduct = results

	finishedResult := &RestResponse{
		Message: "Success Getting Rab Details",
		Status:  200,
		Data:    rabListDetail,
	}

	return finishedResult, nil
}

func (data *RabList) CreateRabList(rab RabList) (*RestResponse, *errors.RestError) {
	stmt, err := rab_db.Client.Prepare(query.QueryCreateRabList)
	if err != nil {
		return nil, errors.BadRequestError("Error when preparing Create RAB List query")
	}
	defer stmt.Close()
	insert, err := stmt.Exec(rab.RabName, rab.Comodity, rab.Province, rab.City)
	if err != nil {
		return nil, errors.BadRequestError("Error when trying to input values to query")
	}
	id, _ := insert.LastInsertId()

	stmt, err = rab_db.Client.Prepare(query.QueryCreateRabLog)
	if err != nil {
		return nil, errors.BadRequestError("Error when trying to create log for Create Rab List")
	}

	notes := fmt.Sprintf("Pembuatan Create RAB List ID %d", id)
	_, err = stmt.Exec(id, notes)

	stmt, err = rab_db.Client.Prepare(query.QueryGetRabListCreatedData)
	if err != nil {
		return nil, errors.BadRequestError("Error when trying to get Created RAB Data")
	}

	createdData, err := stmt.Query(id)
	if err != nil {
		return nil, errors.BadRequestError("Error when getting created RAB Data from database")
	}

	var result RabList

	for createdData.Next() {
		var res RabList
		if err := createdData.Scan(&res.RabName, &res.Comodity, &res.Province, &res.City); err != nil {
			return nil, errors.BadRequestError("Error when looping created Data")
		}
		result = res
	}
	sendData := &RestResponse{
		Message: "Success Create RAB List",
		Status:  200,
		Data:    result,
	}

	return sendData, nil
}

func (data *RabList) GetRabList(offsetInt int64, search string) (*RestResponse, *errors.RestError) {
	stmt, err := rab_db.Client.Prepare(query.QueryGetRabList + query.Limit3)
	if err != nil {
		return nil, errors.BadRequestError("Error when trying to prepare query")
	}
	searchKey := fmt.Sprint("%" + search + "%")
	rows, err := stmt.Query(searchKey, searchKey, searchKey, searchKey, offsetInt)
	if err != nil {
		return nil, errors.BadRequestError("ERROR when trying to get data from database")
	}

	var results []RabList
	for rows.Next() {
		var res RabList
		if err := rows.Scan(&res.Id, &res.RabName, &res.Comodity, &res.Province, &res.City); err != nil {
			return nil, errors.BadRequestError("Error when trying to loop data")
		}
		results = append(results, res)
	}

	var pageSize int64 = 3
	stmt, err = rab_db.Client.Prepare(query.QueryGetRabList)
	if err != nil {
		return nil, errors.BadRequestError("Error when preparing query")
	}

	totalData, err := stmt.Query(searchKey, searchKey, searchKey, searchKey)
	if err != nil {
		return nil, errors.BadRequestError("Error when getting data to database")
	}

	var totalPage float64
	var totalData2 float64
	for totalData.Next() {
		totalData2++
	}
	totalPage = math.Ceil(float64(totalData2) / float64(pageSize))
	currentPage := offsetInt / int64(pageSize)
	if currentPage == 0 {
		currentPage = 1
	}

	finishedData := &RestResponse{
		Status: 200,
		Meta: &MetaDetail{
			TotalPage:   int64(totalPage),
			CurrentPage: currentPage,
			PageSize:    pageSize,
			Total:       int64(totalData2),
		},
		Data:    results,
		Message: "Success Getting Data from Database",
	}
	return finishedData, nil
}
