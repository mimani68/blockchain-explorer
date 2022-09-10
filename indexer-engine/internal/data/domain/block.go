package domain

// swagger:model Block
type Block struct {
	Base
	Number  int64  `json:"number,omitempty" gorm:"primaryKey"`
	Hash    string `json:"hash,omitempty"`
	TxCount int64  `json:"tx_count,omitempty"`
}

func (Block) TableName() string {
	return "block"
}
