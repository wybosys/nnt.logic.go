package logger

import (
	"nnt/core/number"
	"log"
)

type Level byte

const (
	SPECCIAL  Level = 9
	CUSTOM          = 8
	DEBUG           = 7
	INFO            = 6
	NOTICE          = 5
	WARNING         = 4
	ERROR           = 3
	ALERT           = 2
	CRRITICAL       = 1
	EMERGENCE       = 0
	EMERGENCY       = 0
)

type FnLoggerOutput func(string)
type FnLoggerFatal func(error, string)
type FnLoggerError func(error)
type FnLoggerAssert func(interface{}, string)

type ILogger interface {
	Log(str string)
	Warn(str string)
	Info(str string)
	Fatal(e error, str string)
	Exception(e error)
	Error(e error)
	Assert(any interface{}, str string)
}

type BaseLogger struct {
	ILogger
}

func (*BaseLogger) Log(str string) {

}

func (*BaseLogger) Warn(str string) {

}

func (*BaseLogger) Info(str string) {

}

func (*BaseLogger) Fatal(e error, str string) {

}

func (*BaseLogger) Exception(e error) {

}

func (*BaseLogger) Error(e error) {

}

func (self *BaseLogger) Assert(any interface{}, str string) {
	if any == nil {
		self.Fatal(nil, str)
	} else {
		test := true
		switch any.(type) {
		case bool:
			test, _ = any.(bool)
		default:
			v := number.Convert(any, 0)
			test = v != 0
		}
		if !test {
			self.Fatal(nil, str)
		}
	}
}

var (
	Log       FnLoggerOutput
	Warn      FnLoggerOutput
	Info      FnLoggerOutput
	Fatal     FnLoggerFatal
	Exception FnLoggerError
	Error     FnLoggerError
	Assert    FnLoggerAssert
)

func init() {
	Log = func(s string) {
		log.Printf(s)
	}

	Warn = func(s string) {
		log.Printf(s)
	}

	Info = func(s string) {
		log.Printf(s)
	}

	Fatal = func(e error, s string) {
		log.Printf(e.Error())
		log.Fatalf(s)
	}

	Exception = func(e error) {
		log.Printf(e.Error())
	}

	Error = func(e error) {
		log.Printf(e.Error())
	}
}
