package rab

import (
	"fmt"
	"net/http"
	// "rab-service-test2/datasource/mysql/rab_db"
	"rab-service-test2/query"
	"rab-service-test2/utulities"
	"database/sql"

)

var (
	db *sql.DB
)


func (data *RabDataList) CreateRabDataList(r RabDataList) (*RestResponse, *errors.RestError) {
	db,err := sql.Open("mysql","root:admin@tcp(127.0.0.1:3306)/rab_service")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("di dalam dao")
	

	stmt, err := db.Prepare(query.QueryInsertRabDataList)
	if err != nil {
		fmt.Println("error ketika mau persiapan rab list statement")

		return nil, errors.SendError("error ketika mau persiapan rab list statement", http.StatusBadRequest)
	}
	defer stmt.Close()
	result, saveErr := stmt.Exec(r.ProductName, r.UnitProduct, r.Quantity, r.UnitPrice)
	if saveErr != nil {
		return nil, errors.SendError("Error ada di result", http.StatusBadRequest)
	}
	return &RestResponse{
		Status:  http.StatusOK,
		Data:    result,
		// Data : nil,
		Message: "Sukses input barang",
	}, nil
}
