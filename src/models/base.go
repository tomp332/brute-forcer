package models

import (
	"gorm.io/gorm"
	"time"
)

type CustomORMModel struct {
	ID        uint           `gorm:"primaryKey" gorm:"-" json:"id"`
	CreatedAt time.Time      `gorm:"-" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"-" json:"updatedAt" `
	DeletedAt gorm.DeletedAt `gorm:"-" json:"deletedAt" `
}
