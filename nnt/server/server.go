package server

import (
	"github.com/wybosys/nnt.logic.go/nnt/core/kernel"
	"github.com/wybosys/nnt.logic.go/nnt/core/stl"
)

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

type IHttpServer interface {
}

type IConsoleServer interface {
	// 通过控制台执行
	// @params 调用参数
	// @req 请求对象
	// @rsp 响应对象
	Invoke(params *stl.IndexedObject, req interface{}, rsp interface{});
}
