package service

import "app.io/data/entity"

type TransactionService interface {
	LastTransaction() (response entity.Transaction, err error)
	GettransactionByHash(hash string) (response entity.Transaction, err error)
}
