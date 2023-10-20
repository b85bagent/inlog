package main

import (
	"log"

	"github.com/b85bagent/inlog"
)

func main() {

	fileLogger := inlog.NewFileLogger("file.log", 3, 10, 15, false, false)
	fileLogger.SetOutput()

	for i := 0; i < 100000; i++ {
		log.Printf("This is test log #%d. Additional content to make the log entry longer.", i)
	}
}
