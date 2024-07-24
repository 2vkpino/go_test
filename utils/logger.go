package utils

import (
	"log"
	"os"
	"path/filepath"
	"s3_file_uploader/config"
)

const (
	INFO = iota
	WARNING
	ERROR
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	logLevel      int
)

func InitLogger(logFile string) {
	// Создаем все родительские директории, если они не существуют
	err := os.MkdirAll(filepath.Dir(logFile), 0755)
	if err != nil {
		log.Fatalf("Failed to create directories for log file: %v", err)
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	logLevel = config.LogLevel

	LogInfo("Logging initialized")
}

func LogInfo(v ...interface{}) {
	if logLevel <= INFO {
		infoLogger.Println(v...)
	}
}

func LogWarning(v ...interface{}) {
	if logLevel <= WARNING {
		warningLogger.Println(v...)
	}
}

func LogError(v ...interface{}) {
	if logLevel <= ERROR {
		errorLogger.Println(v...)
	}
}
