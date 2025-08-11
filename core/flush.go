package core

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

var mu sync.Mutex

func (l *Logger) Flush() {
	mu.Lock()
	defer mu.Unlock()

	filename := "./.Nutdat/log/crash_report.log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Fehler beim Öffnen der Crash-Report-Datei: %v\n", err)
		return
	}
	defer f.Close()

	// Header mit Timestamp
	timestamp := time.Now().Format(time.RFC3339)
	_, _ = f.WriteString(fmt.Sprintf("\n--- Crash Report - %s ---\n", timestamp))

	// Stacktrace anhängen
	stack := debug.Stack()
	_, _ = f.WriteString("Stacktrace:\n")
	_, _ = f.Write(stack)

	// Speicher-Buffer anhängen
	if len(l.memoryBuffer) > 0 {
		_, _ = f.WriteString("\nBuffered Logs:\n")
		for _, line := range l.memoryBuffer {
			_, _ = f.WriteString(line + "\n")
		}
		// Buffer leeren, damit es beim nächsten Mal nicht wieder drinsteht
		l.memoryBuffer = []string{}
	}

	_, _ = f.WriteString("--- Ende Crash Report ---\n\n")
}
