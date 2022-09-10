package entity

// swagger:model Transaction
type Transaction struct {
	BlockNumber int64  `json:"blockNumber,omitempty"`
	Hash        string `json:"hash,omitempty"`
	From        string `json:"From,omitempty"`
	To          string `json:"to,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Nonce       int64  `json:"nonce,omitempty"`
}
