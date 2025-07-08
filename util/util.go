package util

import (
	"flag"
	"os"

	"github.com/akshaykumar-tech/school-management/model"
	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {

	Logger = *logrus.New()

	Logger.Out = os.Stdout

}

func SetLogger() {

	port := flag.String(model.LogLevel, model.LogLevelInfo, "log-level(debug, error, info, warning)")

	flag.Parse()

	switch *port {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)

	}

}
