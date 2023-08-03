package models

type CredsBase struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type CredsModel struct {
	CustomORMModel
	CredsBase
}

type ICredCreate struct {
	CredsBase
}

type ICredUpdate struct {
	Username    string `json:"username" gorm:"<-"`
	Password    string `json:"password" gorm:"<-"`
	Description string `json:"description" gorm:"<-" `
}
