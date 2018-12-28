package kernel

import (
	"github.com/bitly/go-simplejson"
	"log"
)

type JsonObject struct {
	*simplejson.Json
}

func ToJsonObject(buf []byte) (*JsonObject, error) {
	jsobj, err := simplejson.NewJson(buf)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &JsonObject{jsobj}, nil
}

func ToJson(obj *JsonObject) ([]byte, error) {
	if obj == nil {
		return []byte(""), nil
	}
	return obj.Encode()
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

func (self *JsonObject) GetFloat(key string) (float64, bool) {
	t, ok := self.CheckGet(key)
	if !ok {
		return 0, false
	}
	v, err := t.Float64()
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

func (self *JsonObject) GetObject(key string) (interface{}, bool) {
	t, ok := self.CheckGet(key)
	if !ok {
		return nil, false
	}
	return t.Interface(), true
}
