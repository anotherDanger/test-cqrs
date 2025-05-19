package helpers

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func NewErr(fileName string, level logrus.Level, err error) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{})

	fileOut, fileErr := os.OpenFile(fmt.Sprintf("%s.log", fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if fileErr != nil {
		fmt.Println("Gagal buka file log:", fileErr)
		return nil, fileErr
	}

	logger.SetOutput(fileOut)

	if err != nil {
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
		default:
			logger.Error(err.Error())
		}
		return logger, nil
	}

	switch level {
	case logrus.InfoLevel:
		logger.Info("Info: operasi berhasil")
	case logrus.DebugLevel:
		logger.Debug("Debug: operasi berhasil")
	case logrus.TraceLevel:
		logger.Trace("Trace: operasi berhasil")
	}

	return logger, nil
}
