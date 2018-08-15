package app

import (
	"nnt/server"
	"nnt/core/entry"
	"reflect"
	"nnt/core/proto"
)

type Sample struct {
	entry.IEntry
	server.Rest
}

type Echoo struct {
	entry.IDecl `model(auth) table("db", "echoo")`
	Input       string `string(1, [input], "输入") colstring()`
	Output      string `string(2, [output], "输出") colstring()`
	Time        int    `integer(3, [output], "服务器时间") colinteger()`
}

func init() {
	proto.ParseClass(reflect.TypeOf(Echoo{}))
	entry.Register(reflect.TypeOf(Sample{}))
	entry.Register(reflect.TypeOf(Echoo{}))
}
