package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// removeDataFolder deletes the entire ./data directory including all logs.
func removeDataFolder() error {
	if err := os.RemoveAll("./data"); err != nil {
		return fmt.Errorf("could not remove ./data folder: %w", err)
	}
	return nil
}

// prepareTestLogDir recreates ./data/test/ for use in tests.
func prepareTestLogDir() error {
	logDir := "./data/test/"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("could not create test log directory: %w", err)
	}
	return nil
}

// TestMain handles setup and teardown for all tests.
func TestMain(m *testing.M) {
	// Clean up before running tests
	if err := removeDataFolder(); err != nil {
		fmt.Println("Failed to clean up before tests:", err)
		os.Exit(1)
	}
	if err := prepareTestLogDir(); err != nil {
		fmt.Println("Failed to prepare test log directory:", err)
		os.Exit(1)
	}

	// Run all tests
	code := m.Run()

	// Final cleanup after all tests
	if err := removeDataFolder(); err != nil {
		fmt.Println("Failed to clean up after tests:", err)
	}

	os.Exit(code)
}

func TestNewLogger_CreatesDirectoryAndLogsStart(t *testing.T) {

	logDir := "./data/test/"
	logger := NewLogger(logDir)

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		t.Errorf("Log directory was not created")
	}

	filename := filepath.Join(logDir, time.Now().Format("error_2006_01.log"))
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Logger started") {
		t.Errorf("Log file does not contain 'Logger started'")
	}

	if logger.logDir != logDir {
		t.Errorf("Expected logger logDir to be %q, got %q", logDir, logger.logDir)
	}
}

func TestLogError_WritesToFileAndFallbackToMemory(t *testing.T) {

	logDir := "./data/test/"
	filePath := filepath.Join(logDir, "log.txt")

	// Erstelle eine Datei statt eines Verzeichnisses
	os.WriteFile(filePath, []byte("Not a directory"), 0644)

	logger := NewLogger(filePath)
	logger.memoryBuffer = nil

	logger.LogError(ERROR, "Fallback error message")
	if len(logger.memoryBuffer) == 1 {
		t.Errorf("Expected fallback to memory buffer")
	}
	if !strings.Contains(logger.memoryBuffer[1], "Fallback error message") {
		t.Errorf("Fallback memory buffer does not contain logged message")
	}
}
