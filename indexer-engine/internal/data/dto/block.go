package dto

import "app.io/internal/data/domain"

type CompleteBlock struct {
	Number      int64                `json:"number,omitempty" gorm:"primaryKey"`
	Hash        string               `json:"hash,omitempty"`
	TxCount     int64                `json:"tx_count,omitempty"`
	Transaction []domain.Transaction `json:"transactions,omitempty"`
}
