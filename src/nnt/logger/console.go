package logger

import "log"

type Logger struct {
	BaseLogger
}

func (*Logger) Log(str string) {
	log.Printf(str)
}

func (*Logger) Warn(str string) {
	log.Printf(str)
}

func (*Logger) Info(str string) {
	log.Printf(str)
}

func (*Logger) Fatal(e error, str string) {
	log.Printf(e.Error())
	log.Fatalf(str)
}

func (*Logger) Exception(e error) {
	log.Printf(e.Error())
}

func (*Logger) Error(e error) {
	log.Printf(e.Error())
}
