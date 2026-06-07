package logger

import (
	"io"
	"os"

	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	AppLogger    *slog.Logger
	ErrorLogger  *slog.Logger
	AccessLogger *slog.Logger
)

func Init() {

	// Dossier logs
	_ = os.MkdirAll("logs", os.ModePerm)

	//--------------------------------------------
	// APP Logger
	//--------------------------------------------
	appLogFile := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     30,   // days
		Compress:   true, // disabled by
	}
	appMultiWriter := io.MultiWriter(os.Stdout, appLogFile)
	AppLogger = slog.New(slog.NewJSONHandler(appMultiWriter, nil))

	//--------------------------------------------
	// ERROR Logger
	//--------------------------------------------
	errorLogFile := &lumberjack.Logger{
		Filename:   "logs/error.log",
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     30,   // days
		Compress:   true, // disabled by
	}
	errorMultiWriter := io.MultiWriter(os.Stdout, errorLogFile)
	ErrorLogger = slog.New(slog.NewJSONHandler(errorMultiWriter, nil))

	//--------------------------------------------
	// ACCESS Logger
	//--------------------------------------------
	accessLogFile := &lumberjack.Logger{
		Filename:   "logs/access.log",
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     30,   // days
		Compress:   true, // disabled by
	}
	accessMultiWriter := io.MultiWriter(os.Stdout, accessLogFile)
	AccessLogger = slog.New(slog.NewJSONHandler(accessMultiWriter, nil))
}
