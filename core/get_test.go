package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGetLastLogs_ReturnsCorrectNumber(t *testing.T) {

	logDir := "./data/test/"
	logger := NewLogger(logDir)

	filename := filepath.Join(logDir, time.Now().Format("error_2006_01.log"))
	lines := []string{
		"[2025-06-01 10:00:00] [ERROR] error 1",
		"[2025-06-01 11:00:00] [ERROR] error 2",
		"[2025-06-01 12:00:00] [ERROR] error 3",
		"[2025-06-01 13:00:00] [ERROR] error 4",
		"[2025-06-01 14:00:00] [ERROR] error 5",
	}
	err := os.WriteFile(filename, []byte(strings.Join(lines, "\n")+"\n"), 0644)
	if err != nil {
		fmt.Println(err)
	}

	logs := logger.GetLastLogs(3)

	if len(logs) != 3 {
		t.Errorf("Expected 3 logs, got %d", len(logs))
	}
	expected := "error 3"
	if !strings.Contains(logs[0], expected) {
		t.Errorf("Expected first returned log to contain %q", expected)
	}
}
