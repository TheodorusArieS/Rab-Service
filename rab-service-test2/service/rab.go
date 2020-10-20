package service

import (
	"fmt"
	rab "rab-service-test2/domain"
	errors "rab-service-test2/utilities"
)

type RabServiceInterface interface {
	CreateRabData(rab.RabDataList) (*rab.RestResponse, *errors.RestError)
	GetRabData() (*rab.RestResponse, *errors.RestError)
	GetRabDetails(int64) (*rab.RestResponse,*errors.RestError)
	CreateRabList(rab.RabList) (*rab.RestResponse,*errors.RestError)
	GetRabList()(*rab.RestResponse,*errors.RestError)
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
	dao := &rab.RabDataList{}

	result1, err1 := dao.GetRabDataList()
	if err1 != nil {
		return nil, err1
	}

	return result1, nil
}

func (r *rabService) GetRabDetails(id int64)(*rab.RestResponse,*errors.RestError){
	dao :=&rab.RabList{}
	fmt.Printf("ada di service get rab details %d",id)
	result,err :=dao.GetRabDetails(id)
	return result,err
}

func (r *rabService) CreateRabList(data rab.RabList)(*rab.RestResponse,*errors.RestError){
	dao := &rab.RabList{}
	result,err :=dao.CreateRabList(data)
	return result,err
}

func (r *rabService) GetRabList() (*rab.RestResponse,*errors.RestError){
	dao :=&rab.RabList{}
	result,err :=dao.GetRabList()
	
	return result,err
}