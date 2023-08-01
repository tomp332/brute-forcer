package models

type ServiceStatus uint8

const (
	ONLINE ServiceStatus = iota
	PENDING
	ERROR
	ShuttingDown
)

func (s ServiceStatus) String() string {
	switch s {
	case ONLINE:
		return "ONLINE"
	case PENDING:
		return "PENDING"
	case ERROR:
		return "ERROR"
	case ShuttingDown:
		return "SHUTTING_DOWN"
	}
	return "unknown"
}

type Health struct {
	ID     string        `json:"id"`
	Status ServiceStatus `json:"status"`
	Port   int16         `json:"port"`
}
