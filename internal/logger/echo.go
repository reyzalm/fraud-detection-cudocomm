package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/CudoCommunication/cudocomm/internal/domain"
	"github.com/labstack/echo/v4"
)

type EchoLogger struct {
	logger  echo.Logger
	logFile *os.File
	mu      sync.Mutex
}

func NewEchoLogger(logger echo.Logger) domain.Logger {

	now := time.Now()
	logDir := filepath.Join(".log", now.Format("2006"), now.Format("01"))
	logFileName := filepath.Join(logDir, now.Format("02")+".log")

	if err := os.MkdirAll(logDir, 0755); err != nil {
		logger.Fatalf("Failed to create log directory: %v", err)
	}

	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}

	return &EchoLogger{
		logger:  logger,
		logFile: file,
	}
}

func (el *EchoLogger) writeToFile(level string, payload *domain.LoggerPayload) {
	el.mu.Lock()
	defer el.mu.Unlock()

	payload.Time = time.Now().Format("2006-01-02 15:04:05")

	logEntry, err := json.Marshal(payload)
	if err != nil {

		fmt.Fprintf(os.Stderr, "Failed to marshal log payload: %v\n", err)
		return
	}

	logLine := fmt.Sprintf("[%s] %s\n", level, string(logEntry))
	if _, err := el.logFile.WriteString(logLine); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log file: %v\n", err)
	}
}

func (el *EchoLogger) Error(payload *domain.LoggerPayload) {

	el.writeToFile("ERROR", payload)

	el.logger.Errorf("Location: %s, Message: %s, Request: %+v", payload.Loc, payload.Msg, payload.Req)
}

func (el *EchoLogger) Info(payload *domain.LoggerPayload) {

	el.writeToFile("INFO", payload)
	
	el.logger.Infof("Location: %s, Message: %s", payload.Loc, payload.Msg)
}
