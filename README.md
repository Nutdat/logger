# 📝 Nut Logger

A simple, colorful logging library for Go — logs to files and console.

---

## ⚙️ Features

- 📁 File logging with rotation
- 🖥️ Console logging with colors
- 📦 Log levels: `INFO`, `WARN`, `ERROR`, `FATAL`
- 🧹 Auto cleanup of old logs
- 🧩 Module-based console output

---

## 🚀 Usage

### Import

```go
import "github.com/Nutdat/logger"
```

### Log to File

```go
logger.Info("App started")
logger.Warn("Low disk space")
logger.Error("DB connection failed")
logger.Fatal("Out of memory")
```

### Log to Console

```go 
logger.Console("SQL", "SELECT * FROM users")
```

## or Init
(different color)

```go
logger.LogInit("INIT", "Cache initialized")
```
## JSON Print

```go
logger.PrettyPrintJSON(interface)
```

### clean up all old logs

```go
logger.Cleanup(7) // Keep logs from last 7 months
```

