package model

type CreateTransactionRequest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}

type CreateTransactionResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}

type GetTransactionResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}
