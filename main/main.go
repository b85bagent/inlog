package main

import (
	"log"

	"github.com/b85bagent/inlog"
)

func main() {

	l := inlog.NewLogger(false)

	l.Debug("This is a Debug")
	l.Info("This is a Info")
	l.Warn("This is a Warn")
	l.Error("This is an Error")

	l.WithFields(map[string]interface{}{
		"trace_id":      "1234",
		"game_id":       "5678",
		"身分證":           "9012",
		"名稱":            "chess",
		"user_id":       "3456",
		"exchange_rate": 7.8,
	}).Debug("testLex")

	fileLogger := inlog.NewFileLogger("file.log", 3, 10, 15, false, false)
	fileLogger.SetOutput()

	for i := 0; i < 100000; i++ {
		log.Printf("This is test log #%d. Additional content to make the log entry longer.", i)
	}
}
