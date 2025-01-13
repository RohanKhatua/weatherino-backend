package customlogger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type FileLineHook struct{}

func (hook *FileLineHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook FileLineHook) Fire(entry *logrus.Entry) error {
	_, file, line, ok := runtime.Caller(8)

	if ok {
		entry.Data["file"] = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	return nil
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := map[logrus.Level]string{
		logrus.DebugLevel: "\033[36m",
		logrus.InfoLevel:  "\033[32m",
		logrus.WarnLevel:  "\033[33m",
		logrus.ErrorLevel: "\033[31m",
		logrus.FatalLevel: "\033[35m",
		logrus.PanicLevel: "\033[35m",
	}

	color := levelColor[entry.Level]
	reset := "\033[0m"

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fileLine := entry.Data["file"]

	logMsg := fmt.Sprintf(
		"%s[%s] %s%-7s%s %s - %v\n",
		color, timestamp, color, entry.Level, reset, fileLine, entry.Message,
	)

	return []byte(logMsg), nil
}

var Logger *logrus.Logger

func Init() {
	Logger = logrus.New()
	Logger.SetOutput(Logger.Out)
	Logger.SetLevel(Logger.Level)
	Logger.SetFormatter(&CustomFormatter{})
	Logger.AddHook(&FileLineHook{})
}
