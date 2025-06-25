package logger

import "github.com/Nutdat/logger/core"

// logger is the internal instance of the core.Logger used by this package.
var logger *core.Logger

// init initializes the default logger instance with a log directory.
// It is automatically called when the package is imported.
func init() {
	logger = core.NewLogger("./data/log/")
}

// Cleanup removes log files older than the specified number of days.
//
// Parameters:
//   - days: number of days to keep log files. Older logs will be deleted.
func Cleanup(days int) {
	logger.CleanupLogs(days)
}

// Info logs an informational message with the "INFO" level.
//
// Parameters:
//   - msg: the message to log.
func Info(msg string) {
	logger.LogError("INFO", msg)
}

// Warn logs a warning message with the "WARN" level.
//
// Parameters:
//   - msg: the message to log.
func Warn(msg string) {
	logger.LogError("WARN", msg)
}

// Error logs an error message with the "ERROR" level.
//
// Parameters:
//   - msg: the message to log.
func Error(msg string) {
	logger.LogError("ERROR", msg)
}

// Fatal logs a critical error message with the "FATAL" level.
//
// Parameters:
//   - msg: the message to log.
func Fatal(msg string) {
	logger.LogError("FATAL", msg)
}

// Console prints a custom message directly to the console with a module tag.
//
// Parameters:
//   - module: a short identifier (e.g., "SQL", "CACHE") to prefix the message.
//   - msg: the message to print.
func Console(module, msg string) {
	core.LogtoConsole(module, msg)
}

// LogInit prints an initialization-related message to the console,
// using a dark yellow module tag and dark gray message text.
//
// Parameters:
//   - module: name of the module or package being initialized.
//   - msg: descriptive message about the init process.
func LogInit(module, msg string) {
	core.LogInitMessage(module, msg)
}
