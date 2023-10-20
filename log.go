package inlog

import (
	"fmt"
	"log"
	"sync"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

var levelNames = map[LogLevel]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
}

type Logger struct {
	debugMode bool
	fields    map[string]interface{}
	mu        sync.Mutex
}

func NewLogger(debugMode bool) *Logger {
	return &Logger{debugMode: debugMode}
}

func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	return &Logger{
		debugMode: l.debugMode,
		fields:    fields,
	}
}

func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if level >= Info || l.debugMode {
		msg := fmt.Sprintf(format, v...)
		if l.fields != nil {
			for k, v := range l.fields {
				msg = fmt.Sprintf("%s %s=%v", msg, k, v)
			}
		}
		log.Printf("[" + levelNames[level] + "] " + msg)
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(Debug, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(Info, format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(Warn, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(Error, format, v...)
}
