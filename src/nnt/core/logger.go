package core

import "nnt"

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

type FnLoggerOutput func(string, *nnt.Any)

var (
	Log FnLoggerOutput = func(_ string, _ *nnt.Any) {

	}
)
