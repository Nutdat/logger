package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// writeLogToFile attempts to append a log line to current monthly log file.
// Returns true on success, false otherwise.

func (l *Logger) writeLogToFile(logLevel, logLine string) bool {
	now := time.Now()
	filename := fmt.Sprintf("%s_%d_%02d.log", logLevel, now.Year(), now.Month())
	filePath := filepath.Join(l.logDir, filename)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		l.addToMemoryBuffer(fmt.Sprintf("[LOGGER ERROR] Could not open file: %v", err))
		return false
	}
	defer file.Close()

	if _, err := file.WriteString(logLine); err != nil {
		l.addToMemoryBuffer(fmt.Sprintf("[LOGGER ERROR] Failed to write to file: %v", err))
		return false
	}

	return true
}

// CleanupLogs deletes log files older than the given number of months.
// If duration is 0 or negative, it defaults to 12 months.

func (l *Logger) CleanupLogs(duration int) {
	if duration <= 0 {
		duration = 12
	}

	files, err := os.ReadDir(l.logDir)
	if err != nil {
		l.LogError(ERROR, "CleanupLogs failed to read log directory: "+err.Error())
		return
	}

	cutoff := time.Now().AddDate(0, -duration, 0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		var year, month int
		n, err := fmt.Sscanf(file.Name(), "error_%d_%02d.log", &year, &month)
		if err != nil || n != 2 {
			continue
		}

		fileTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		if fileTime.Before(cutoff) {
			err = os.Remove(filepath.Join(l.logDir, file.Name()))
			if err != nil {
				l.LogError(ERROR, "CleanupLogs failed to remove file "+file.Name()+": "+err.Error())
			}
		}
	}
}
