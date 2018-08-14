package url

import (
	"strings"
	"os"
	"path/filepath"
	"fmt"
	"nnt/logger"
)

type FnSchemeProcessor func(string) string

var schemes = make(map[string]FnSchemeProcessor)
var HOME, _ = os.Getwd()
var ROOT, _ = filepath.Abs("/")

// 展开url
// 如果包含 :// 则拆分成 scheme 和 body，再根绝 scheme 注册的转换器转换
// 否则按照 / 来打断各个部分，再处理 ~、/ 的设置
func Expand(url string) string {
	if fnd := strings.Index(url, "://"); fnd != -1 {
		ps := strings.Split(url, "://")
		proc := schemes[ps[0]]
		if proc == nil {
			logger.Log(fmt.Sprintf("没有注册该类型 %s 的处理器", ps[0]))
			return ""
		}
		return proc(ps[1])
	}

	ps := strings.Split(url, "/")
	if ps[0] == "~" {
		ps[0] = HOME
	} else if ps[0] == "" {
		ps[0] = ROOT
	} else {
		return url
	}

	return strings.Join(ps, "/")
}

func RegisterScheme(scheme string, proc FnSchemeProcessor) {
	schemes[scheme] = proc
}

func init() {
	RegisterScheme("http", func(s string) string {
		return s
	})
	RegisterScheme("https", func(s string) string {
		return s
	})
}
