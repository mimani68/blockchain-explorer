package dto

type TatumTransaction struct {
	BlockNumber int64  `json:"blockNumber,omitempty"`
	Hash        string `json:"hash,omitempty"`
	From        string `json:"From,omitempty"`
	To          string `json:"to,omitempty"`
	Amount      string `json:"value,omitempty"`
	Nonce       int64  `json:"nonce,omitempty"`
}
