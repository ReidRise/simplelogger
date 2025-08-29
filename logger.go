package simplelog

import "fmt"

type Level int

const (
	ErrorLevel Level = 10
	WarnLevel  Level = 8
	InfoLevel  Level = 6
	DebugLevel Level = 4
	TraceLevel Level = 2
)

var LogLevel Level = InfoLevel

func Error(msg string) {
	if LogLevel <= ErrorLevel {
		fmt.Printf("[ERROR]: %s\n", msg)
	}
}

func Warn(msg string) {
	if LogLevel <= WarnLevel {
		fmt.Printf("[WARN] : %s\n", msg)
	}
}

func Debug(msg string) {
	if LogLevel <= DebugLevel {
		fmt.Printf("[DEBUG]: %s\n", msg)
	}
}

func Info(msg string) {
	if LogLevel <= InfoLevel {
		fmt.Printf("[INFO] : %s\n", msg)
	}
}

func Trace(msg string) {
	if LogLevel <= TraceLevel {
		fmt.Printf("[TRACE]: %s\n", msg)
	}
}
