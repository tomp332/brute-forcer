package internalTypes

import "time"

type Task struct {
	StartTime time.Time
	EndTime   time.Duration
}

type EncryptionTask struct {
	ResultHash     string
	TargetPassword string
	Task
}

type MetaData struct {
	Size        int
	TextMessage string
}

type EncryptionTaskResult struct {
	TargetPassword string
	EncryptedHash  string
	MetaData
	Task
}

type DecryptionTask struct {
	TargetHash     string
	Mode           int16
	WordlistPath   string
	ResultPassword string
	NumAttempts    int
	MetaData
	Task
}
