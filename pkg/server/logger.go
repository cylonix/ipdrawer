package server

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.JSONFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}
