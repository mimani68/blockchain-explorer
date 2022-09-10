package service

import "app.io/data/model"

type TransactionService interface {
	List() (responses []model.GetTransactionResponse)
	Single(request model.CreateTransactionRequest) (response model.CreateTransactionResponse)
}
