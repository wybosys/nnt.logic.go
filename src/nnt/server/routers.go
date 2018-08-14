package server

import (
	"nnt/core/entry"
	"nnt/logger"
	"nnt/core/proto"
)

type IRouter interface {
	Action() string
}

type Routers struct {
	entry.IEntry
	routers map[string]IRouter
}

type IRouterable interface {
	Routers() *Routers
}

type BaseRouter struct {
	IRouter
	action string
}

func (self *BaseRouter) Action() string {
	return self.action
}

func (self *Routers) Init() {
	self.routers = make(map[string]IRouter)
}

func (self *Routers) Length() int {
	return len(self.routers)
}

func (self *Routers) Register(obj IRouter) {
	actnm := obj.Action()
	if _, fnd := self.routers[actnm]; fnd {
		logger.Fatal(nil, "已经注册了一个同名的路由 "+actnm)
		return
	}
	self.routers[actnm] = obj
}

func (self *Routers) Find(name string) IRouter {
	t, _ := self.routers[name]
	return t
}

func (self *Routers) Foreach(proc func(r IRouter, nm string)) {
	for k, v := range self.routers {
		proc(v, k)
	}
}

func (self *Routers) Process(trans *Transaction) {
	// 查找router
	r := self.Find(trans.router)
	if r == nil {
		trans.Status = proto.ROUTER_NOT_FOUND
		trans.Submit()
		return
	}
}