package server

import "nnt/core/kernel"

type IServer interface {
	Config(cfg *kernel.JsonObject) bool
	Start()
	onStart()
}

type BaseServer struct {
	IServer
	id string
}

func (self *BaseServer) Config(cfg *kernel.JsonObject) bool {
	if nodeid, ok := cfg.GetString("id"); ok {
		self.id = nodeid
		return true
	}
	return false
}

func (self *BaseServer) Start() {
	self.onStart()
}

func (self *BaseServer) onStart() {
	// pass
}
