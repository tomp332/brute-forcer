package internalTypes

import "time"

type Task struct {
	StartTime time.Time
	EndTime   time.Duration
}

type EncryptionTask struct {
	Hash     string
	Password string
	Task
}

type TaskResult struct {
	Password string
	Hash     string
	Task
}

type DecryptionTask struct {
	Hash         string
	Mode         int16
	WordlistPath string
	Password     string
	NumAttempts  int
	Task
}
