package models

type SlaveBase struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port int16  `json:"port"`
}

type SlaveModel struct {
	CustomORMModel
	SlaveBase
}
