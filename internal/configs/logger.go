package configs

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Formatter struct{ logrus.TextFormatter }

func Logger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(new(Formatter))
	log.SetReportCaller(true)
	writer, _ := rotatelogs.New(
		fmt.Sprintf("%s/%%Y-%%m-%%d.log", "logs"),
		rotatelogs.WithMaxAge(time.Hour*24*14),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	log.SetOutput(writer)
	log.SetLevel(logrus.InfoLevel)
	return log
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	format := fmt.Sprintf(
		"%s - %s - %s - %d - %s\n",
		entry.Time.Format("2006-01-02 15:04:05"),
		strings.ToUpper(entry.Level.String()),
		entry.Caller.Function,
		entry.Caller.Line,
		entry.Message,
	)
	return []byte(format), nil
}
