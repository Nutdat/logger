package logger

import "github.com/Nutdat/logger/core"

var logger *core.Logger

func init() {
	logger = core.NewLogger("./data/log/")
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

func Fatal(msg string) {
	logger.LogError("FATAL", msg)
}

func Console(module, msg string) {
	core.LogtoConsole(module, msg)
}

func LogInit(module, msg string) {
	core.LogInitMessage(module, msg)
}
