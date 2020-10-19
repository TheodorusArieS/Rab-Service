package service

import (
	"fmt"
	rab "rab-service-test2/domain"
	errors "rab-service-test2/utulities"
)

type RabServiceInterface interface {
	CreateRabData(rab.RabDataList) (*rab.RestResponse, *errors.RestError)
	GetRabData() (*rab.RestResponse, *errors.RestError)
}

type rabService struct {
}

var (
	RabService RabServiceInterface = &rabService{}
)

func (r *rabService) CreateRabData(data rab.RabDataList) (*rab.RestResponse, *errors.RestError) {
	fmt.Println(data)
	dao := &rab.RabDataList{}
	result, err := dao.CreateRabDataList(data)
	return result, err
}

func (r *rabService) GetRabData() (*rab.RestResponse, *errors.RestError) {
	// result := &rab.RestResponse{
	// 	Status:  200,
	// 	Data:    nil,
	// 	Message: "Error",
	// }
	// err := errors.SendError("test error", 400)
	dao := &rab.RabDataList{}

	result1, err1 := dao.GetRabDataList()
	if err1 != nil {
		return nil, err1
	}

	return result1, nil
}
