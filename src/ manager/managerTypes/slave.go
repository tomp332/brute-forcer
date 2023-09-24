package managerTypes

type Slave struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port int16  `json:"port"`
}

type SlaveDTO struct {
	CustomORMModel
	Slave
}

func (SlaveDTO) TableName() string {
	return "Slaves"
}
