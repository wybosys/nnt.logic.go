package kernel

import (
	"github.com/bitly/go-simplejson"
	"log"
)

type JsonObject struct {
	*simplejson.Json
}

func ToJsonObject(buf []byte) *JsonObject {
	jsobj, err := simplejson.NewJson(buf)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	return &JsonObject{jsobj}
}

func (self *JsonObject) CheckGet(key string) (*JsonObject, bool) {
	v, ok := self.Json.CheckGet(key)
	return &JsonObject{v}, ok
}

func (self *JsonObject) JsArray() ([]*JsonObject, error) {
	t, err := self.Array()
	if err != nil {
		return make([]*JsonObject, 0), err
	}
	r := make([]*JsonObject, 0)
	for idx := range t {
		r = append(r, &JsonObject{self.GetIndex(idx)})
	}
	return r, nil
}

func (self *JsonObject) GetString(key string) (string, bool) {
	t, ok := self.CheckGet(key)
	if !ok {
		return "", false
	}
	v, err := t.String()
	if err != nil {
		return "", false
	}
	return v, true
}

func (self *JsonObject) GetInt(key string) (int, bool) {
	t, ok := self.CheckGet(key)
	if !ok {
		return 0, false
	}
	v, err := t.Int()
	if err != nil {
		return 0, false
	}
	return v, true
}

func (self *JsonObject) GetBool(key string) (bool, bool) {
	t, ok := self.CheckGet(key)
	if !ok {
		return false, false
	}
	v, err := t.Bool()
	if err != nil {
		return false, false
	}
	return v, true
}
