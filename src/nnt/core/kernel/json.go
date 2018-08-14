package kernel

import (
	"github.com/bitly/go-simplejson"
	"nnt/core"
)

type JsonObject struct {
	*simplejson.Json
}

func ToJsonObject(buf []byte) *JsonObject {
	jsobj, err := simplejson.NewJson(buf)
	if err != nil {
		core.Logger.Error(err)
		return nil
	}
	return &JsonObject{jsobj}
}
