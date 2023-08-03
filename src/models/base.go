package models

import (
	"gorm.io/gorm"
	"time"
)

type CustomORMModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt" gorm:"<-"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"<-"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"<-"`
}
