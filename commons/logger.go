package commons

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

type MyFormatter struct{}

var levelList = []string{
	"PANIC",
	"FATAL",
	"ERROR",
	"WARN",
	"INFO",
	"DEBUG",
	"TRACE",
}

func (mf *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// define supported log levels
	level := levelList[int(entry.Level)]
	strList := strings.Split(entry.Caller.File, "/")
	// get the go file name
	fileName := strList[len(strList)-1]
	b.WriteString(fmt.Sprintf("%28s [%5s] %s:%d - %s\n",
		// Custom Time Format
		entry.Time.Format("2006-01-02T15:04:05"),
		level, fileName, entry.Caller.Line, entry.Message))
	return b.Bytes(), nil
}

func GetLogger() *logrus.Logger {
	logFile := "log.txt"
	var f *os.File
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	var standardLogger = logrus.New()

	standardLogger.SetLevel(logrus.DebugLevel)
	standardLogger.SetFormatter(&MyFormatter{})
	standardLogger.SetOutput(io.MultiWriter(f, os.Stdout))
	standardLogger.SetReportCaller(true)

	return standardLogger
}
