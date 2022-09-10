package entity

// swagger:model Block
type Block struct {
	Number  int64  `json:"number,omitempty"`
	Hash    string `json:"hash,omitempty"`
	TxCount int64  `json:"tx_count,omitempty"`
}
