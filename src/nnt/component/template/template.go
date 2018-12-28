package template

import (
	"fmt"
	"nnt/core/stl"
	"regexp"
	"strings"
)

var RE_PARAMETER, _ = regexp.Compile(`\\{\\{([a-zA-Z0-9_.]+)\\}\\}'`)

type Template struct {
	buf string
}

func (self *Template) CompileCustom(str string, pat *regexp.Regexp) *Template {
	// 标记str中的变量，避免循环填数据
	self.buf = string(pat.ReplaceAll([]byte("@@__$1"), []byte(str)))
	return self
}

func (self *Template) Compile(str string) *Template {
	self.CompileCustom(str, RE_PARAMETER)
	return self
}

func (self *Template) Render(parameters *stl.IndexedObject) string {
	var fn_each_param = func(v interface{}, k interface{}) {
		self.buf = strings.Replace(self.buf, fmt.Sprintf("@@__%s", k.(string)), v.(string), -1)
	}
	parameters.Foreach(fn_each_param)
	return self.buf
}
