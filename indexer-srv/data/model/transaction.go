package model

import "app.io/data/entity"

type TransactionResponse struct {
	Success     string             `json:"success"`
	Message     string             `json:"message"`
	Transaction entity.Transaction `json:"transaction"`
}

type TransactionRequest struct {
	TransactionHash string `json:"transactionHash"`
}
