package core

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSearchLogs_SingleDateMatch(t *testing.T) {

	logDir := "./data/test/"
	logger := NewLogger(logDir)

	filename := filepath.Join(logDir, "error_2025_05.log")
	content := `[2025-05-25 10:00:00] [ERROR] error on 25th
[2025-05-26 11:00:00] [ERROR] error on 26th
`
	os.WriteFile(filename, []byte(content), 0644)

	logs, err := logger.SearchLogs("ERROR", "25.05.2025")
	if err != nil {
		t.Fatalf("SearchLogs failed: %v", err)
	}
	if len(logs) != 1 {
		t.Errorf("Expected 1 log entry for 25.05.2025, got %d", len(logs))
	}
	if !strings.Contains(logs[0], "error on 25th") {
		t.Errorf("Unexpected log content: %s", logs[0])
	}
}

func TestSearchLogs_MultipleMonths(t *testing.T) {

	logDir := "./data/test/"
	logger := NewLogger(logDir)

	// Mai-Datei
	contentMay := `[2025-05-30 10:00:00] [ERROR] error in May
`
	os.WriteFile(filepath.Join(logDir, "error_2025_05.log"), []byte(contentMay), 0644)

	// Juni-Datei
	contentJune := `[2025-06-01 11:00:00] [ERROR] error in June
[2025-06-02 12:00:00] [INFO] info in June
`
	os.WriteFile(filepath.Join(logDir, "error_2025_06.log"), []byte(contentJune), 0644)

	// Juli-Datei
	contentJuly := `[2025-07-01 14:00:00] [ERROR] error in July
`
	os.WriteFile(filepath.Join(logDir, "error_2025_07.log"), []byte(contentJuly), 0644)

	logs, err := logger.SearchLogs("ERROR", "30.05.2025", "01.07.2025")
	if err != nil {
		t.Fatalf("SearchLogs failed: %v", err)
	}
	if len(logs) != 3 {
		t.Errorf("Expected 3 error logs across months, got %d", len(logs))
	}
}

func TestSearchLogs_FiltersByTypeOnly(t *testing.T) {
	logDir := "./data/test/"
	logger := NewLogger(logDir)

	content := `[2025-06-15 10:00:00] [ERROR] something bad
[2025-06-15 11:00:00] [INFO] something fine
`
	os.WriteFile(filepath.Join(logDir, "error_2025_06.log"), []byte(content), 0644)

	logs, err := logger.SearchLogs("INFO", "15.06.2025")
	if err != nil {
		t.Fatalf("SearchLogs failed: %v", err)
	}
	if len(logs) != 1 || !strings.Contains(logs[0], "[INFO]") {
		t.Errorf("Expected 1 INFO log, got: %v", logs)
	}
}
