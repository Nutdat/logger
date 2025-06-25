package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SearchLogs searches logs filtered by log type, and a date or date range.
//
// Parameters:
// - logTypeStr: e.g. "ERROR", empty string for no filter
// - startDateStr: date string "02.01.2006" (required)
// - endDateStr: optional, same format, if empty or missing, searches only startDate
//
// Returns matching logs or an error if date parsing fails.
func (l *Logger) SearchLogs(logTypeStr, startDateStr string, endDateStr ...string) ([]string, error) {
	const inputLayout = "02.01.2006"
	const logLayout = "2006-01-02 15:04:05"

	startDate, err := time.Parse(inputLayout, strings.TrimSpace(startDateStr))
	if err != nil {
		return nil, fmt.Errorf("invalid start date format (expected DD.MM.YYYY): %w", err)
	}

	// Default: only that one day
	start := startDate.Truncate(24 * time.Hour)
	end := start.Add(24*time.Hour - time.Nanosecond)

	// If endDateStr is given, override end
	if len(endDateStr) > 0 && strings.TrimSpace(endDateStr[0]) != "" {
		endDate, err := time.Parse(inputLayout, strings.TrimSpace(endDateStr[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid end date format (expected DD.MM.YYYY): %w", err)
		}
		end = endDate.Truncate(24 * time.Hour).Add(24*time.Hour - time.Nanosecond)
	}

	// Optional log type filter
	var logType *LogType
	if logTypeStr != "" {
		lt := LogType(strings.ToUpper(strings.TrimSpace(logTypeStr)))
		logType = &lt
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	var results []string
	// Iterate by month from start to end
	for t := time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.UTC); t.Year() < end.Year() || (t.Year() == end.Year() && t.Month() <= end.Month()); t = t.AddDate(0, 1, 0) {
		filename := fmt.Sprintf("error_%d_%02d.log", t.Year(), t.Month())
		filePath := filepath.Join(l.logDir, filename)

		data, err := os.ReadFile(filePath)
		if err != nil {
			continue // Datei evtl. nicht vorhanden
		}

		lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
		for _, line := range lines {
			if len(line) < 21 {
				continue
			}

			timeStr := line[1:20]
			logTime, err := time.Parse(logLayout, timeStr)
			if err != nil {
				continue
			}

			if logTime.Before(start) || logTime.After(end) {
				continue
			}

			if logType != nil {
				if !strings.Contains(line, "["+string(*logType)+"]") {
					continue
				}
			}

			results = append(results, line)
		}
	}

	return results, nil
}
