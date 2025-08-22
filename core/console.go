package core

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// LogError writes a log entry to file or memory fallback, and optionally prints it.
func (l *Logger) LogError(logType LogType, errorMessage string) {
	now := time.Now()
	logLine := fmt.Sprintf("[%s] [%s] %s\n", now.Format("2006-01-02 15:04:05"), logType, errorMessage)

	success := l.writeLogToFile(string(logType), logLine)
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
		colorGreen     = "\033[32m"
	)

	logTypeColor := colorCyan
	switch logType {
	case INFO:
		logTypeColor = colorGreen
	case WARN:
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

// simple log just to console
func LogtoConsole(module, message string) {
	const (
		colorReset     = "\033[0m"
		colorDarkGray  = "\033[90m"
		colorTimestamp = "\033[36m" // cyan
		colorModule    = "\033[34m" // blue
	)

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	formatted := fmt.Sprintf(
		"%s[%s]%s %s[%s]%s %s%s%s\n",
		colorTimestamp, timestamp, colorReset,
		colorModule, module, colorReset,
		colorDarkGray, message, colorReset,
	)

	fmt.Print(formatted)
}

// LogInitMessage logs a message with a dark yellow module tag and dark gray message.
func LogInitMessage(module string) {
	const (
		colorReset     = "\033[0m"
		colorDarkGray  = "\033[90m"
		colorTimestamp = "\033[36m" // cyan
		colorModule    = "\033[97m" // white
	)
	message := module + " successfully initialized"
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	formatted := fmt.Sprintf(
		"%s[%s]%s %s[%s]%s %s%s%s\n",
		colorTimestamp, timestamp, colorReset,
		colorModule, "INIT", colorReset,
		colorDarkGray, message, colorReset,
	)

	fmt.Print(formatted)
}

func PrettyPrintJSON(v interface{}) {
	// Erst JSON erzeugen
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Fehler beim Marshallen: %v\n", err)
		return
	}

	// Dann wieder nach interface{} parsen
	var data interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Printf("Fehler beim Unmarshallen: %v\n", err)
		return
	}

	printColored(data, 0)
	fmt.Println()
}

func printColored(v interface{}, indent int) {
	const (
		reset   = "\033[0m"
		green   = "\033[32m" // Numbers
		cyan    = "\033[36m" // Keys
		yellow  = "\033[33m" // Bool/Null
		magenta = "\033[35m" // Strings
	)
	ind := strings.Repeat("  ", indent)

	switch val := v.(type) {
	case map[string]interface{}:
		fmt.Println("{")
		i := 0
		for k, vv := range val {
			fmt.Print(ind + "  " + cyan + `"` + k + `"` + reset + ": ")
			printColored(vv, indent+1)
			i++
			if i < len(val) {
				fmt.Print(",")
			}
			fmt.Println()
		}
		fmt.Print(ind + "}")
	case []interface{}:
		fmt.Println("[")
		for i, vv := range val {
			fmt.Print(ind + "  ")
			printColored(vv, indent+1)
			if i < len(val)-1 {
				fmt.Print(",")
			}
			fmt.Println()
		}
		fmt.Print(ind + "]")
	case string:
		fmt.Print(magenta + strconv.Quote(val) + reset)
	case float64:
		fmt.Print(green + fmt.Sprintf("%v", val) + reset)
	case bool:
		fmt.Print(yellow + fmt.Sprintf("%v", val) + reset)
	case nil:
		fmt.Print(yellow + "null" + reset)
	default:
		fmt.Print(fmt.Sprintf("%v", val))
	}
}
