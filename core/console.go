package core

import (
	"fmt"
	"time"
)

// LogError writes a log entry to file or memory fallback, and optionally prints it.
func (l *Logger) LogError(logType LogType, errorMessage string) {
	now := time.Now()
	logLine := fmt.Sprintf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), logType, errorMessage)

	success := l.writeLogToFile(logLine)
	if !success {
		l.addToMemoryBuffer(logLine)
	}

	if !l.isProduction() {
		l.printToConsole(now, logType, errorMessage)
	}
}

// printToConsole prints formatted and colored log entry to the console.
func (l *Logger) printToConsole(t time.Time, logType LogType, message string) {
	const (
		colorReset     = "\033[0m"
		colorCyan      = "\033[36m"
		colorBlue      = "\033[34m"
		colorYellow    = "\033[33m"
		colorRed       = "\033[31m"
		colorMagenta   = "\033[35m"
		colorLightGray = "\033[37m"
	)

	logTypeColor := colorCyan
	switch logType {
	case INFO:
		logTypeColor = colorBlue
	case WARNING:
		logTypeColor = colorYellow
	case ERROR:
		logTypeColor = colorRed
	case FATAL:
		logTypeColor = colorMagenta
	}

	formatted := fmt.Sprintf(
		"%s[%s]%s %s[%s]%s %s%s%s\n",
		colorCyan, t.Format("2006-01-02 15:04:05"), colorReset,
		logTypeColor, logType, colorReset,
		colorLightGray, message, colorReset,
	)

	fmt.Print(formatted)
}
