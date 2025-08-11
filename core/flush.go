package core

func (l *Logger) Flush() {
	l.DumpMemoryLogs()
}
