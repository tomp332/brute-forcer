package managerTypes

import "time"

type CredentialsBase struct {
	Username string `json:"username" gorm:"uniqueIndex:idx_username_password"`
	Password string `json:"password" gorm:"uniqueIndex:idx_username_password"`
	Hash     string `json:"hash" gorm:"not null;uniqueIndex:idx_hash"`
}

type CredentialsDTO struct {
	CustomORMModel
	CredentialsBase
}

func (CredentialsDTO) TableName() string {
	return "Credentials"
}

type ICredentialsCreate struct {
	CredentialsBase
}

type IUpdateCredentials struct {
	ID uint `json:"id"`
	CredentialsBase
}

type IReadCredentials struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	CredentialsBase
}
