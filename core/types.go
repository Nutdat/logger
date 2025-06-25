package core

type LogType string

const (
	INFO    LogType = "INFO"
	WARNING LogType = "WARNING"
	ERROR   LogType = "ERROR"
	FATAL   LogType = "FATAL"
)
