package logger

import "github.com/Nutdat/logger/core"

var logger *core.Logger

func init() {
	core.NewLogger("./data/log/")
}

func Cleanup(days int) {
	logger.CleanupLogs(days)
}

func Info(msg string) {
	logger.LogError("INFO", msg)
}

func Warn(msg string) {
	logger.LogError("WARN", msg)
}

func Error(msg string) {
	logger.LogError("ERROR", msg)
}
