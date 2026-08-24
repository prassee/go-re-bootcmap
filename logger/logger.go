package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel string

const (
	INFO    LogLevel = "info"
	DEBUG   LogLevel = "debug"
	ERROR   LogLevel = "error"
	FATAL   LogLevel = "fatal"
	WARNING LogLevel = "warning"
)

type Logger interface {
	Log(level LogLevel, message string)
	Info(message string)
	Debug(message string)
	Error(message string)
	Fatal(message string)
}

type BaseLogger struct {
	logFunc func(LogLevel, string)
}

func (c BaseLogger) Log(level LogLevel, message string) {
	c.logFunc(level, message)
}

func (c BaseLogger) Info(message string) {
	c.Log(INFO, message)
}

func (c BaseLogger) Error(message string) {
	c.Log(ERROR, message)
}

func (c BaseLogger) Debug(message string) {
	c.Log(DEBUG, message)
}

func (c BaseLogger) Fatal(message string) {
	c.Log(FATAL, message)
}

type ConsoleLogger struct {
	prefix string
	BaseLogger
}

func NewConsoleLogger(prefix string) ConsoleLogger {
	c := ConsoleLogger{
		prefix: prefix,
		logFunc: func(level LogLevel, message string) {
			fmt.Printf("[%s] %v - %v: %s\n", prefix, time.Now().Format(time.RFC3339), level, message)
		}}
	return c
}

type FileLogger struct {
	File *os.File
	BaseLogger
}

func NewFileLogger(file *os.File) FileLogger {
	f := FileLogger{
		File: file,
		logFunc: func(level LogLevel, message string) {
			fmt.Fprintf(file, "[%s] %s : %s\n", time.Now().Format(time.RFC3339), level, message)
		},
	}
	return f
}

var _ Logger = FileLogger{}
var _ Logger = ConsoleLogger{}
