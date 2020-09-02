package provider

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger = nil

func newLogrus() *logrus.Logger {
	log := logrus.New()
	return log
}

func GetLogger() *logrus.Logger {
	if log == nil {
		log = newLogrus()
	}
	return log
}
