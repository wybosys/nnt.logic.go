package manager

import (
	"nnt/logger"
	"nnt/core/kernel"
)

type prvLoggers struct {
}

var Loggers *prvLoggers = &prvLoggers{}

func (*prvLoggers) Start(cfg *kernel.JsonObject) {

}

func init() {
	logger.Log = func(s string) {

	}

	logger.Warn = func(s string) {

	}

	logger.Info = func(s string) {

	}

	logger.Fatal = func(e error, s string) {

	}

	logger.Exception = func(e error) {

	}

	logger.Error = func(e error) {

	}
}
