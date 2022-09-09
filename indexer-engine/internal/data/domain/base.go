package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID                   string     `json:"id,omitempty" bson:"_id" gorm:"primaryKey"`
	CreatedAt            time.Time  `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt            *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
	DeletedAt            *time.Time `json:"deletedAt,omitempty" bson:"deletedAt"`
	OperationDescription string     `json:"opsDescription,omitempty" bson:"opsDescription"`
	Status               string     `json:"status,omitempty" bson:"status"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return
}
