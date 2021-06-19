package server

import "github.com/wybosys/nnt.logic.go/nnt/core/entry"

type Rest struct {
	entry.IEntry
	BaseServer
	IRouterable
	IHttpServer
	IConsoleServer
}

func (self *Rest) Init() {

}
