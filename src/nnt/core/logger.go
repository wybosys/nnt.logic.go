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

type FnLoggerOutput func(string, ...nnt.Any)
type FnLoggerError func(error)

type prvLogger struct {
	Log       FnLoggerOutput
	Warn      FnLoggerOutput
	Info      FnLoggerOutput
	Fatal     FnLoggerOutput
	Exception FnLoggerError
	Error     FnLoggerError
}

var Logger = &prvLogger{
	Log: func(s string, any ...nnt.Any) {
		log.Printf(s, any)
	},

	Warn: func(s string, any ...nnt.Any) {
		log.Printf(s, any)
	},

	Info: func(s string, any ...nnt.Any) {
		log.Printf(s, any)
	},

	Fatal: func(s string, any ...nnt.Any) {
		log.Fatalf(s, any)
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
		self.Fatal(s)
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
			self.Fatal(s)
		}
	}
}
