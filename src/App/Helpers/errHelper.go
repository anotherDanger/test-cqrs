package helpers

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func NewErr(fileName string, level logrus.Level, err error) (*logrus.Logger, error) {
	if err != nil {
		logger := logrus.New()
		logger.SetLevel(level)
		logger.SetFormatter(&logrus.JSONFormatter{})
		fileOut, fileErr := os.OpenFile(fmt.Sprintf("%s.log", fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if fileErr != nil {
			panic(err)
		}

		logger.SetOutput(fileOut)
		switch level {
		case logrus.DebugLevel:
			logger.Debug(err.Error())
		case logrus.InfoLevel:
			logger.Info(err.Error())
		case logrus.WarnLevel:
			logger.Warn(err.Error())
		case logrus.ErrorLevel:
			logger.Error(err.Error())
		case logrus.FatalLevel:
			logger.Fatal(err.Error())
		case logrus.PanicLevel:
			logger.Panic(err.Error())
		}
		return logger, fileErr
	} else {
		return nil, nil
	}
}
