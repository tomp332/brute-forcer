package models

type CredentialsDTO struct {
	ICredentialsCreate
	CustomORMModel
}

func (CredentialsDTO) TableName() string {
	return "Credentials"
}

type ICredentialsCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type IUpdateCredentials struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"<-"`
	Password string `json:"password" gorm:"<-"`
}

type IReadCredentials struct {
	ID uint `json:"id"`
	ICredentialsCreate
}
