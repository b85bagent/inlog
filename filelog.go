package inlog

import (
	"log"

	"github.com/natefinch/lumberjack"
)

type FileLogger struct {
	logger *lumberjack.Logger
}

// filename: 日誌檔案的路徑。
// maxSize: 日誌檔案的最大大小（以MB為單位）。當日誌檔案的大小超過這個值時，會創建一個新的日誌檔案。
// maxBackups: 保留的舊日誌檔案的最大數量。
// maxAge: 保留舊日誌檔案的最大天數。
// compress: 是否壓縮舊日誌檔案。如果設定為true，舊日誌檔案會被壓縮成gz格式。
// LocalTime: 是否使用本地時間作為日誌檔案的時間戳。如果設定為true，會使用本地時間，否則使用UTC時間。
func NewFileLogger(filename string, maxSize, maxBackups, maxAge int, compress, LocalTime bool) *FileLogger {
	l := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
		LocalTime:  LocalTime,
	}
	return &FileLogger{logger: l}
}

func (f *FileLogger) SetOutput() {
	log.SetOutput(f.logger)
}
