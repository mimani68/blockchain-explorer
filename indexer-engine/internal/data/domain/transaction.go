package domain

// swagger:model Transaction
type Transaction struct {
	Base
	BlockNumber int64  `json:"block_number,omitempty" gorm:"primaryKey"`
	Hash        string `json:"hash,omitempty"`
	From        string `json:"From,omitempty"`
	To          string `json:"to,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Nonce       int64  `json:"nonce,omitempty"`
}

func (Transaction) TableName() string {
	return "transaction"
}
