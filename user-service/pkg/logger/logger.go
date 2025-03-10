package logger

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLog() {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func makeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return logrus.WithFields(logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func Info(c echo.Context, message string) {
	logEntry := makeLogEntry(c)

	logEntry.Info(message)
}

func Warn(c echo.Context, message string) {
	logEntry := makeLogEntry(c)

	logEntry.Warn(message)
}

func Error(c echo.Context, message string) {
	logEntry := makeLogEntry(c)

	logEntry.Error(message)
}
