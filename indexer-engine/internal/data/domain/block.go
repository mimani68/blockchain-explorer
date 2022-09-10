package domain

// swagger:model Block
type Block struct {
	Base
	Number  string `json:"number,omitempty"`
	Hash    string `json:"hash,omitempty"`
	TxCount string `json:"tx_count,omitempty"`
}

func (Block) TableName() string {
	return "app.block"
}
