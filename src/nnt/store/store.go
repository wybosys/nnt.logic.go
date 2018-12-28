package store

import "nnt/core/kernel"

// 数据库基类接口
type IDbms interface {
	// 打开连接
	Open();

	// 关闭连接
	Close();
}

// 数据库默认实现
type AbstractDbms struct {
	// 唯一标记
	Id string
}

// 开始事务
func (self *AbstractDbms) Begin() {
	// pass
}

// 完成事务
func (self *AbstractDbms) Complete() {
	// pass
}

// 取消事务
func (self *AbstractDbms) Cancel() {
	// pass
}

// 读取配置
func (self *AbstractDbms) Config(cfg *kernel.JsonObject) bool {
	self.Id, _ = cfg.GetString("id")
	return true
}

// 数据库执行情况
type DbExecuteStat struct {
	// 增加的行数
	Insert int

	// 修改行数
	Update int

	// 删除行数
	Remove int
}
