package logger

import (
	"log"
	"nnt/core/entry"
	"reflect"
)

type Console struct {
	BaseLogger
}

func (*Console) Log(str string) {
	log.Printf(str)
}

func (*Console) Warn(str string) {
	log.Printf(str)
}

func (*Console) Info(str string) {
	log.Printf(str)
}

func (*Console) Fatal(e error, str string) {
	log.Printf(e.Error())
	log.Fatalf(str)
}

func (*Console) Exception(e error) {
	log.Printf(e.Error())
}

func (*Console) Error(e error) {
	log.Printf(e.Error())
}

func init() {
	entry.Register(reflect.TypeOf(Console{}))
}
