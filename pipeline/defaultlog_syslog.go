// +build !windows,!nacl,!plan9

package pipeline

import (
	"log"
	"log/syslog"
)

// ForceLog should rarely be used. It forceable logs an entry to the
// Windows Event Log (on Windows) or to the SysLog (on Linux)
func ForceLog(severity LogSeverity, msg string) {
	if defaultLogger == nil {
		return // Return fast if we failed to create the logger.
	}
	// We are logging it, ensure trailing newline
	if len(msg) == 0 || msg[len(msg)-1] != '\n' {
		msg += "\n" // Ensure trailing newline
	}
	switch severity {
	case LogFatal:
		defaultLogger.Fatal(msg)
	case LogPanic:
		defaultLogger.Panic(msg)
	case LogError, LogWarning, LogInfo:
		defaultLogger.Print(msg)
	}
}

var defaultLogger = func() *log.Logger {
	l, _ := syslog.NewLogger(syslog.LOG_USER|syslog.LOG_WARNING, log.LstdFlags)
	return l
}()
