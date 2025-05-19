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
		return logger, fileErr
	} else {
		return nil, nil
	}
}
