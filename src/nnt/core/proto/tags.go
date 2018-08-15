package proto

import (
	"reflect"
	"nnt/core/entry"
)

// 解析符合nnt规范的标注

type TaggedClass struct {
	// 类定义
	Class TagInfos

	// 成员变量定义
	Fields map[string]TagInfos
}

func newTaggedClass() *TaggedClass {
	return &TaggedClass{
		Fields: make(map[string]TagInfos),
	}
}

type TagArgument struct {
	// 是否是常量定义
	IsConst bool

	// 字串数据
	Str string

	// 数组数据
	Array []*TagArgument
}

type TagInfo struct {
	// 标记的处理函数名称
	Function string

	// 参数表
	Args []*TagArgument
}

type TagInfos map[string]*TagInfo

var (
	typEntry = reflect.TypeOf((*entry.IEntry)(nil)).Elem()
	typDecl  = reflect.TypeOf((*entry.IDecl)(nil)).Elem()
)

func ParseClass(typ reflect.Type) *TaggedClass {
	// 只有实现 IEntry 或者 IDecl 的类型才支持解析
	var implTyp string
	if typ.Implements(typEntry) {
		implTyp = "IEntry"
	} else if typ.Implements(typDecl) {
		implTyp = "IDecl"
	} else {
		return nil
	}

	r := newTaggedClass()

	// 获得类定义
	typFd, _ := typ.FieldByName(implTyp)

	// 解析类tag
	r.Class = ParseTag(string(typFd.Tag))

	// 解析成员变量
	for i, l := 0, typ.NumField(); i < l; i++ {
		typFd := typ.Field(i)
		if typFd.Name == implTyp {
			continue
		}
		if typFd.Tag == "" {
			continue
		}
		infos := ParseTag(string(typFd.Tag))
		if infos != nil {
			r.Fields[typFd.Name] = infos
		}
	}

	return r
}

func ParseTag(tag string) TagInfos {
	if tag == "" {
		return nil
	}
	return nil
}
