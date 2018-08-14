package core

import (
	"nnt"
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

type prvLogger struct {
	Log       FnLoggerOutput
	Warn      FnLoggerOutput
	Info      FnLoggerOutput
	Fatal     FnLoggerFatal
	Exception FnLoggerError
	Error     FnLoggerError
}

var Logger = &prvLogger{
	Log: func(s string) {
		log.Printf(s)
	},

	Warn: func(s string) {
		log.Printf(s)
	},

	Info: func(s string) {
		log.Printf(s)
	},

	Fatal: func(e error, s string) {
		log.Printf(e.Error())
		log.Fatalf(s)
	},

	Exception: func(e error) {
		log.Printf(e.Error())
	},

	Error: func(e error) {
		log.Printf(e.Error())
	},
}

func (self *prvLogger) Assert(exp nnt.Any, s string) {
	if exp == nil {
		self.Fatal(nil, s)
	} else {
		test := true
		switch exp.(type) {
		case bool:
			test, _ = exp.(bool)
		default:
			v := 0 //ToNumber(exp, 0)
			test = v != 0
		}
		if !test {
			self.Fatal(nil, s)
		}
	}
}
