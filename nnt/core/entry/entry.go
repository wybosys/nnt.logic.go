package entry

import (
	"log"
	"reflect"
	"strings"
)

// 定义实体的基本接口，Instance时会调用Init达到初始化的目的，如果不需要该特性，使用IDecl
type IEntry interface {
	Init()
}

// 用于描述对象时标记struct描述,否则解析器无法定位信息
type IDecl interface {

}

// 注册实体，用来动态实例化

var entries = make(map[string]reflect.Type)

func Register(clazz reflect.Type, name ...string) {
	var key string
	if len(name) == 0 {
		key = clazz.PkgPath() + "/" + clazz.Name()
	} else {
		key = name[0]
	}
	key = strings.ToLower(key)
	if _, ok := entries[key]; ok {
		log.Fatal("已经注册同名的类型 " + key)
		return
	}
	entries[key] = clazz
}

func Instance(name string) interface{} {
	key := strings.ToLower(name)
	key = strings.Replace(key, ".", "/", -1)
	if typ, ok := entries[key]; ok {
		hdl := reflect.New(typ).Interface()
		if ehdl, ok := hdl.(IEntry); ok {
			ehdl.Init()
			return ehdl
		}
		return hdl
	}
	return nil
}
