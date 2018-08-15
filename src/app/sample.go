package app

import (
	"nnt/server"
	"nnt/core/entry"
	"reflect"
)

type Sample struct {
	entry.IEntry
	server.Rest
}

func (self *Sample) Init() {
	self.Rest.Init()
}

func init() {
	entry.Register(reflect.TypeOf(Sample{}))
}
