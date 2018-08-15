package server

import (
	"nnt/core/entry"
	"strings"
	"nnt/core/stl"
	"nnt/core/proto"
	"nnt/core/datetime"
)

type ITransaction interface {
	// 返回事务用来区分客户端的id，通常业务中实现为sid
	SessionId() string

	// 恢复上下文，涉及到数据的恢复，所以是异步模式
	Collect()

	// 是否已经授权
	Auth() bool
}

type TransactionInfo struct {
	// 客户端
	Agent string

	// 访问的主机
	Host   string
	Origin string

	// 客户端的地址
	Addr string

	// 来源
	Referer string
	Path    string
}

type TransactionSubmitOption struct {
	// 仅输出模型
	Model bool

	// 直接输出数据
	Raw bool

	// 输出的类型
	otype string
}

type Transaction struct {
	ITransaction
	entry.IEntry

	// 动作
	action string

	// 映射到router的执行器中
	router string
	call   string

	// 参数
	params stl.StringMap

	// 执行的结果
	Status proto.STATUS

	// 错误信息
	message string

	// 额外数据
	payload stl.IndexObject

	// 输出和输入的model
	model stl.IndexObject

	// 基于哪个服务器运行
	server IServer

	// 是否是压缩数据
	gzip bool

	// 是否暴露接口（通常只有登录会设置为true)
	expose bool

	// 此次的时间
	time int64

	// 环境信息
	info TransactionInfo
}

func (self *Transaction) Init() {
	self.Status = proto.UNKNOWN
	self.time = datetime.Current()
}

func (self *Transaction) Action() string {
	return self.action
}

func (self *Transaction) SetAction(act string) {
	self.action = act
	p := stl.StringArray(strings.Split(act, "."))
	self.router = strings.ToLower(p.At(0, "null"))
	self.call = strings.ToLower(p.At(1, "null"))
}

func (self *Transaction) Modelize(r IRouter) proto.STATUS {
	return proto.OK
}

func (self *Transaction) Submit(opt ...TransactionSubmitOption) {

}

func (self *Transaction) NeedAuth() bool {
	return false
}