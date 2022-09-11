package model

import "app.io/data/entity"

type ScanRequest struct {
	StartBlock int `json:"startBlock"`
	EndBlock   int `json:"endBlock"`
}

type ScanResponse struct {
	Success string         `json:"success"`
	Message string         `json:"message"`
	Block   []entity.Block `json:"block"`
}
