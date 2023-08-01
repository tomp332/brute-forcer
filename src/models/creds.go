package models

import "time"

type Cred struct {
	Id          int64     `gorm:"primary_key:auto_increment" json:"-"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
