package types

import "time"

type BruteForceBase struct {
	Hash        string `json:"hash" gorm:"not null;uniqueIndex:idx_password_hash"`
	PlainText   string `json:"plainText"`
	Algorithm   string `json:"algorithm"`
	NumOfSlaves int    `json:"numSlaves"`
}

type BruteForceDTO struct {
	CustomORMModel
	BruteForceBase
}

func (BruteForceDTO) TableName() string {
	return "BruteForces"
}

type IBruteForceCreate struct {
	Hash        string `json:"hash" gorm:"not null;uniqueIndex:idx_hash"`
	Algorithm   string `json:"algorithm"`
	NumOfSlaves int    `json:"numSlaves"`
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
