package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// InitLogger initializes the logger
func InitLogger() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using system environment variables")
	}

	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO"
	}

	level, err := logrus.ParseLevel(strings.ToUpper(logLevel))
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = filepath.Join(xdg.DataHome, "MultiCrypt", "logs")
	}
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
		return
	}

	logFile := filepath.Join(logDir, "multicrypt.log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return
	}

	logger.SetOutput(file)
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}
