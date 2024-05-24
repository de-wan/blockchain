package common

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func ConfigureLogging(filename string) {
	// configure logging
	logFilePath := filename + ".log"

	// Initialize the logger with log rotation settings
	logFile := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // compress the rotated logs
	}
	defer logFile.Close()

	// Create a multi-writer that writes to both file and console
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Set log output to the multi-writer
	log.SetOutput(multiWriter)
}
