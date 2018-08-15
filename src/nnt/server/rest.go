package server

import "nnt/core/entry"

type Rest struct {
	entry.IEntry
	BaseServer
	IRouterable
	IHttpServer
	IConsoleServer
}

func (self *Rest) Init() {

}
