package app

import (
	"github.com/wybosys/nnt.logic.go/nnt/core/entry"
	"github.com/wybosys/nnt.logic.go/nnt/server"
	"reflect"
)

type Sample struct {
	entry.IEntry
	server.Rest
}

type Echoo struct {
	entry.IDecl `model(auth) table("db" , "echoo")`
	Input       string `string(1, [input], "输入") colstring(  )`
	Output      string `string(2, [output], "输出") colstring()`
	Time        int    `integer(3, [output], "服务器时间") colinteger()`
	Tmp         string // 测试用，不输出
}

func init() {
	entry.Register(reflect.TypeOf(Sample{}))
	entry.Register(reflect.TypeOf(Echoo{}))
}
