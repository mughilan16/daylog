package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
	debugLogger *log.Logger

	debugEnabled bool
)

func Init(debug bool) {
	debugEnabled = debug

	infoLogger = log.New(os.Stdout, "INFO  ", log.LstdFlags)
	errorLogger = log.New(os.Stderr, "ERROR ", log.LstdFlags)
	fatalLogger = log.New(os.Stderr, "FATAL ", log.LstdFlags)

	if debugEnabled {
		debugLogger = log.New(os.Stdout, "DEBUG ", log.LstdFlags)
	}
}

func Info(msg string) {
	infoLogger.Println(msg)
}

func Error(msg string) {
	errorLogger.Println(msg)
}

func Fatal(msg string) {
	fatalLogger.Println(msg)
	os.Exit(1)
}

func Debug(msg string) {
	if !debugEnabled {
		return
	}
	debugLogger.Println(msg)
}
