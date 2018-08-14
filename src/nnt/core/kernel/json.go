package kernel

import (
	"github.com/bitly/go-simplejson"
	"nnt/logger"
)

type JsonObject struct {
	*simplejson.Json
}

func ToJsonObject(buf []byte) *JsonObject {
	jsobj, err := simplejson.NewJson(buf)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return &JsonObject{jsobj}
}
