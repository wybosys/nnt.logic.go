package entry

import (
	"reflect"
	"strings"
	"log"
)

type IEntry interface {
	Init()
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
		log.Fatal(name + " 没有实现 IEntry 接口")
		return nil
	}
	return nil
}
