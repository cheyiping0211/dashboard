package global

import (
	"time"
)

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey json:"id""`
	CreatedAt time.Time
	UpdatedAt time.Time
}