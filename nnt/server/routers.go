package server

import (
	"github.com/wybosys/nnt.logic.go/nnt/config"
	"github.com/wybosys/nnt.logic.go/nnt/core/entry"
	"github.com/wybosys/nnt.logic.go/nnt/core/proto"
	"github.com/wybosys/nnt.logic.go/nnt/logger"
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

	// 模型化
	sta := trans.Modelize(r)
	if sta != proto.OK {
		trans.Status = sta
		trans.Submit()
		return
	}

	// 恢复数据上下文
	trans.Collect()

	// 判断权限判断
	if !trans.expose {
		if trans.NeedAuth() {
			if !trans.Auth() {
				trans.Status = proto.NEED_AUTH
				trans.Submit()
				return
			}
		} else {
			pass := self.devopsCheck(trans)
			if !pass {
				trans.Status = proto.PERMISSIO_FAILED
				trans.Submit()
				return
			}
		}
	}

	trans.Status = proto.OK
}

func (self *Routers) devopsCheck(trans *Transaction) bool {
	// devops环境下才进行权限判定
	if config.LOCAL {
		return true
	}

	// 允许客户端访的将无法进行服务端权限判定
	if config.CLIENT_ALLOW {
		return true
	}

	// 如果访问的是api.doc，则不进行判定
	if trans.Action() == "api.doc" {
		return true
	}

	// 判断是否允许跳过权限验证
	if config.DEVOPS_DEVELOP {

	}

	return true
}
