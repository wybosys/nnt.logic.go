package kernel

import (
	"github.com/wybosys/nnt.logic.go/nnt/core/kernel/variant_type"
	"strconv"
)

type VariantType int

type Variant struct {
	// 原始数据
	raw interface{}

	// 类型
	typ VariantType

	// 携带的数据
	buf   []byte
	str   string
	bol   bool
	num   interface{}
	jsobj *JsonObject
}

func NewVairant(o interface{}) *Variant {
	r := new(Variant)
	r.setRaw(o)
	return r
}

func (r *Variant) setRaw(o interface{}) {
	r.raw = o
	if o == nil {
		return
	}

	switch o.(type) {
	case string:
		r.str = o.(string)
		r.typ = variant_type.STRING;
		break
	case []byte:
		r.buf = o.([]byte)
		r.typ = variant_type.BUFFER
		break
	case bool:
		r.bol = o.(bool)
		r.typ = variant_type.BOOLEAN
		break
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
		r.num = o.(int)
		r.typ = variant_type.NUMBER
		break
	case float32:
	case float64:
		r.num = o.(float64)
		r.typ = variant_type.NUMBER
		break
	case JsonObject:
		r.jsobj = o.(*JsonObject)
		r.typ = variant_type.OBJECT
		break
	}
}

func (self *Variant) Raw() interface{} {
	return self.raw
}

func (self *Variant) Object() interface{} {
	return self.jsobj
}

func (self *Variant) Value() interface{} {
	switch self.typ {
	case variant_type.STRING:
		return self.str
	case variant_type.BUFFER:
		return self.buf
	case variant_type.OBJECT:
		return self.jsobj
	case variant_type.BOOLEAN:
		return self.bol
	case variant_type.NUMBER:
		return self.num
	}
	return nil
}

func (self *Variant) SetValue(v interface{}) {
	switch self.typ {
	case variant_type.STRING:
		self.str = v.(string)
		break
	case variant_type.BUFFER:
		self.buf = v.([]byte)
		break
	case variant_type.OBJECT:
		self.jsobj = v.(*JsonObject)
		break
	case variant_type.BOOLEAN:
		self.bol = v.(bool)
		break
	case variant_type.NUMBER:
		self.num = v
		break
	}
}

func (self *Variant) toBuffer() []byte {
	if self.buf != nil {
		return self.buf
	}
	self.buf = []byte(self.ToString())
	return self.buf
}

func (self *Variant) ToString() string {
	if self.str != "" {
		return self.str
	}
	switch self.typ {
	case variant_type.BUFFER:
		self.str = string(self.buf)
		break
	case variant_type.OBJECT:
		b, err := ToJson(self.jsobj)
		if err == nil {
			self.str = string(b)
		}
		break
	case variant_type.BOOLEAN:
		if self.bol {
			self.str = "true"
		} else {
			self.str = "false"
		}
		break
	case variant_type.NUMBER:
		switch self.num.(type) {
		case int:
			self.str = strconv.Itoa(self.num.(int))
			break
		default:
			self.str = strconv.FormatFloat(self.num.(float64), 'f', 6, 64)
			break
		}
	}
	return self.str
}

func (self *Variant) ToJsObj() *JsonObject {
	if self.jsobj != nil {
		return self.jsobj
	}
	self.jsobj, _ = ToJsonObject([]byte(self.ToString()))
	return self.jsobj
}

func (self *Variant) Serialize() []byte {
	js := new(JsonObject)
	js.Set("_t", self.typ)
	js.Set("_i", "vo")
	js.Set("_d", self.Value())
	r, _ := ToJson(js)
	return r
}

func (self *Variant) Unserialize(str []byte) bool {
	obj, _ := ToJsonObject(str)
	if obj == nil {
		return false
	}
	if v, _ := obj.GetString("_i"); v != "vo" {
		self.setRaw(obj)
	} else {
		vt, _ := obj.GetInt("_t")
		self.typ = VariantType(vt)
		vv, _ := obj.GetObject("_d")
		self.SetValue(vv)
	}
	return true
}

func VariantFromString(str []byte) *Variant {
	if str == nil {
		return nil
	}
	r := new(Variant)
	if !r.Unserialize(str) {
		return nil
	}
	return r
}
