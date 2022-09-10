package service

import (
	"app.io/data/model"
	"app.io/pkg/validation"
)

func NewTransactionService(transactionGrpcClient interface{}) TransactionService {
	return &transactionServiceImpl{
		transactionGrpcClient: transactionGrpcClient,
	}
}

type transactionServiceImpl struct {
	transactionGrpcClient interface{}
}

func (service *transactionServiceImpl) Create(request model.CreateTransactionRequest) (response model.CreateTransactionResponse) {
	validation.Validate(request)

	response = model.CreateTransactionResponse{
		Id:       "1",
		Name:     "transaction.Name",
		Price:    10,
		Quantity: 10,
	}
	return response
}

func (service *transactionServiceImpl) List() (responses []model.GetTransactionResponse) {
	item := model.GetTransactionResponse{
		Id:       "1",
		Name:     "transaction.Name",
		Price:    10,
		Quantity: 10,
	}
	responses = append(responses, item)
	return
}
