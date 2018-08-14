package number

import (
	"nnt"
	"strings"
	"strconv"
)

func Convert(any nnt.Any, def nnt.Number) nnt.Number {
	if any == nil {
		return def;
	}

	switch any.(type) {
	case int:
		return any
	case uint:
		return any
	case int8:
		return any
	case uint8:
		return any
	case int16:
		return any
	case uint16:
		return any
	case int32:
		return any
	case uint32:
		return any
	case int64:
		return any
	case uint64:
		return any
	case uintptr:
		return any
	case bool:
		if v, _ := any.(bool); v {
			return 1
		} else {
			return 0
		}
	case float32:
		return any
	case float64:
		return any
	case string:
		// 判断有没有小数点
		str, _ := any.(string)
		if fnd := strings.Index(str, "."); fnd == -1 {
			r, _ := strconv.Atoi(str)
			return r
		} else {
			r, _ := strconv.ParseFloat(str, 0)
			return r
		}
	}

	return def
}
