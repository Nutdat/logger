package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GetLastLogs returns the last 'count' logs from memory buffer.
// If count > stored logs, returns all.
func (l *Logger) GetLastLogs(count int) []string {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	filename := fmt.Sprintf("error_%d_%02d.log", now.Year(), now.Month())
	filePath := filepath.Join(l.logDir, filename)

	data, err := os.ReadFile(filePath)
	if err != nil {
		// Falls Datei nicht gelesen werden kann, gib leere Liste zurück
		l.addToMemoryBuffer(fmt.Sprintf("[LOGGER ERROR] Could not read log file: %v", err))
		return []string{}
	}

	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")

	n := len(lines)
	if count > n {
		count = n
	}

	return lines[n-count:]
}
