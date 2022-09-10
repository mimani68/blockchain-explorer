package service

import "app.io/data/model"

type TransactionService interface {
	Create(request model.CreateTransactionRequest) (response model.CreateTransactionResponse)
	List() (responses []model.GetTransactionResponse)
}
