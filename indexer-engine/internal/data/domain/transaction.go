package domain

// swagger:model Transaction
type Transaction struct {
	Base
	BlockNumber string `json:"block_number,omitempty"`
	Hash        string `json:"hash,omitempty"`
	From        string `json:"From,omitempty"`
	To          string `json:"to,omitempty"`
	Amount      string `json:"amount,omitempty"`
	Nonce       string `json:"nonce,omitempty"`
}

func (Transaction) TableName() string {
	return "transaction"
}
