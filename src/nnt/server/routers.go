package server

type IRouter interface {
	Action() string
}

type Routers struct {
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

func (self *Routers) Length() int {
	return 0
}