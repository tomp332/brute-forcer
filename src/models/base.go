package models

import (
	"gorm.io/gorm"
	"time"
)

type CustomORMModel struct {
	ID        uint           `gorm:"primaryKey" json:"id" column:"id"`
	CreatedAt time.Time      `json:"createdAt" column:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" column:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" column:"deletedAt"`
}
