package proto

import (
	"nnt/core/entry"
	"reflect"
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

func newTagArgument() *TagArgument {
	return &TagArgument{
		IsConst: true,
	}
}

type TagInfo struct {
	// 标记的处理函数名称
	Function string

	// 参数表
	Args []*TagArgument
}

func newTagInfo() *TagInfo {
	return &TagInfo{
		Args: make([]*TagArgument, 0),
	}
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
	r := make(TagInfos)
	var start int
	var info *TagInfo
	var arg, subarg *TagArgument
	var instring, eatchar bool
	for pos, c := range tag {
		if !instring && c == ' ' {
			continue
		}
		if !eatchar {
			start = pos
			eatchar = true
		}
		if instring && c != '"' {
			continue
		}
		switch c {
		case '(':
			info = newTagInfo()
			info.Function = tag[start:pos]
			arg = newTagArgument()
			eatchar = false
		case ')':
			if arg != nil {
				if start != pos {
					arg.Str = tag[start:pos]
					info.Args = append(info.Args, arg)
				}
				arg = nil
			}
			r[info.Function] = info
			info = nil
			eatchar = false
		case '[':
			arg.Array = make([]*TagArgument, 0)
			subarg = newTagArgument()
			eatchar = false
		case ']':
			if subarg != nil {
				subarg.Str = tag[start:pos]
				arg.Array = append(arg.Array, subarg)
				subarg = nil
			}
			eatchar = false
		case ' ':
			// skip
		case '"':
			if instring {
				if subarg != nil {
					subarg.Str = tag[start:pos]
					arg.Array = append(arg.Array, subarg)
					subarg = nil
				} else {
					arg.Str = tag[start:pos]
					info.Args = append(info.Args, arg)
					arg = nil
				}
				instring = false
				eatchar = false
			} else {
				if subarg != nil {
					subarg.IsConst = false
				} else {
					arg.IsConst = false
				}
				instring = true
				eatchar = false
			}
		case ',':
			if subarg != nil {
				subarg.Str = tag[start:pos]
				arg.Array = append(arg.Array, subarg)
				subarg = newTagArgument()
			} else if arg != nil {
				arg.Str = tag[start:pos]
				info.Args = append(info.Args, arg)
				arg = newTagArgument()
			} else {
				arg = newTagArgument()
			}
		}
	}
	return r
}
