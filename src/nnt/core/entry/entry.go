package entry

import "reflect"

// 注册实体，用来动态实例化

var entries = make(map[string]reflect.Type)

func Register(name string, clazz reflect.Type) {
	entries[name] = clazz
}
