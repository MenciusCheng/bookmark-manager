package log

import (
	"fmt"
	"log"
	"time"
)

func Debug(msg string, v ...any) {
	log.Println(append([]any{prefix("DEBUG", msg)}, v...)...)
}

func Info(msg string, v ...any) {
	log.Println(append([]any{prefix("INFO", msg)}, v...)...)
}

func Warn(msg string, v ...any) {
	log.Println(append([]any{prefix("WARN", msg)}, v...)...)
}

func Error(msg string, v ...any) {
	log.Println(append([]any{prefix("ERROR", msg)}, v...)...)
}

func prefix(level string, msg string) string {
	return fmt.Sprintf("%s [%s] %s", level, time.Now().Format("2006-01-02 15:04:05 -0700"), msg)
}
