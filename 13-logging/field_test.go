package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("uasername", "Masred").Info("Hello World")

	logger.WithField("username", "Ganteng").WithField("name", "Masred").Info("Hello")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"uasername": "Masred",
		"name":      "Masred Ganteng",
	}).Info("Hello World")
}
