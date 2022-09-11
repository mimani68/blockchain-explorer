package model

type Stats struct {
	TotalAmount       string `json:"totalAmount"`
	TotalTransactions string `json:"totalTransactions"`
}

type StatsRequest struct {
	StartBlock int `json:"startBlock"`
	EndBlock   int `json:"endBlock"`
}

type StatsResponse struct {
	Success string  `json:"success"`
	Message string  `json:"message"`
	Stats   []Stats `json:"stats"`
}
