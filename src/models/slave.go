package models

type Slave struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Ip        string `json:"ip"`
	Port      int16  `json:"port"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
