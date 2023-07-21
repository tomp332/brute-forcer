package models

import "time"

type Slave struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Ip        string    `json:"ip"`
	Port      int16     `json:"port"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
