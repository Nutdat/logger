package core

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	logDir       string
	memoryBuffer []string
	mu           sync.Mutex
}

// NewLogger creates a Logger and ensures log directory exists.

func NewLogger(logDir string) *Logger {
	if logDir == "" {
		logDir = "./data/logs/"
	}

	logger := &Logger{logDir: logDir}

	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		logLine := fmt.Sprintf("[%s] [%s] Failed to create log directory: %v\n",
			time.Now().Format("2006-01-02 15:04:05"), ERROR, err)
		logger.addToMemoryBuffer(logLine)
	}

	// Log that logger started
	LogInitMessage("Logger")

	return logger
}

// addToMemoryBuffer adds a log line safely to in-memory buffer.
func (l *Logger) addToMemoryBuffer(line string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.memoryBuffer) > 1000 {
		l.memoryBuffer = l.memoryBuffer[1:]
	}
	l.memoryBuffer = append(l.memoryBuffer, line)
}

// DumpMemoryLogs prints all logs currently stored in memory.
func (l *Logger) DumpMemoryLogs() {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Println("=== In-Memory Logs ===")
	for _, entry := range l.memoryBuffer {
		fmt.Println(entry)
	}
	fmt.Println("======================")
}

// isProduction returns true if environment is production.
func (l *Logger) isProduction() bool {
	return strings.ToLower(os.Getenv("APP_ENV")) == "production"
}
