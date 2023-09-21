package models

import "time"

type BruteForceBase struct {
	Hash        string `json:"hash" gorm:"not null;uniqueIndex:idx_hash"`
	NumOfSlaves int    `json:"numSlaves"`
}

type BruteForceDTO struct {
	CustomORMModel
	BruteForceBase
}

func (CredentialsDTO) TableName() string {
	return "BruteForces"
}

type IBruteForceCreate struct {
	BruteForceBase
}

type IBruteForceUpdate struct {
	ID uint `json:"id"`
	BruteForceBase
}

type IBruteForceRead struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	BruteForceBase
}
