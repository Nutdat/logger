package core

type LogType string

const (
	INFO  LogType = "INFO"
	WARN  LogType = "WARN"
	ERROR LogType = "ERROR"
	FATAL LogType = "FATAL"
)
