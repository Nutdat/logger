package core

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCleanupLogs_DeletesOldFiles(t *testing.T) {
	logDir := "./data/test/"
	logger := NewLogger(logDir)

	oldDate := time.Now().AddDate(-1, 0, 0)
	oldFile := filepath.Join(logDir, oldDate.Format("error_2006_01.log"))
	os.WriteFile(oldFile, []byte("old log"), 0644)

	newFile := filepath.Join(logDir, time.Now().Format("error_2006_01.log"))
	os.WriteFile(newFile, []byte("new log"), 0644)

	logger.CleanupLogs(0)

	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Errorf("Old log file was not deleted")
	}
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Errorf("New log file was deleted incorrectly")
	}
}
