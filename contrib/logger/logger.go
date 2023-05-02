package logger

import (
	"github.com/mlvhub/dummy-core/core"
	log "github.com/sirupsen/logrus"
)

func LogLowercase(str string) {
	log.Info(core.ToLowercase(str))
}
