package model

import "app.io/data/entity"

type BlockRequest struct {
	BlockNumber int `json:"blockNumber"`
}

type BlockResponse struct {
	Success string       `json:"success"`
	Message string       `json:"message"`
	Block   entity.Block `json:"block"`
}
