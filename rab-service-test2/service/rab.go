package service

import (
	"fmt"
	rab "rab-service-test2/domain"
	errors "rab-service-test2/utulities"
)

type RabServiceInterface interface {
	CreateRabData(rab.RabDataList) (*rab.RestResponse, *errors.RestError)
}

type rabService struct {
}

var (
	RabService RabServiceInterface = &rabService{}
)

func (r *rabService) CreateRabData(data rab.RabDataList) (*rab.RestResponse, *errors.RestError) {
	// result := &rab.RestResponse{
	// 	Data:    nil,
	// 	Status:  200,
	// 	Message: "in Create Rab Service",
	// }
	// err := errors.SendError("create rab service", 200)
	fmt.Println(data)
	dao := &rab.RabDataList{}
	result, err := dao.CreateRabDataList(data)
	return result,err
}
