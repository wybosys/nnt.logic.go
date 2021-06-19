package stl

// 类似于js中的object，可以是array，也可以是map，但不能同时
// 为了降低go中构造非结构数据的难度
type IndexedObject struct {
	arr  []interface{}
	mapp map[interface{}]interface{}
}

func (self *IndexedObject) Length() int {
	if self.arr != nil {
		return len(self.arr)
	}
	if self.mapp != nil {
		return len(self.mapp)
	}
	return 0
}

func (self *IndexedObject) isArray() bool {
	return self.arr != nil
}

func (self *IndexedObject) isMap() bool {
	return self.mapp != nil
}

func (self *IndexedObject) Array() []interface{} {
	return self.arr
}

func (self *IndexedObject) Map() map[interface{}]interface{} {
	return self.mapp
}

func (self *IndexedObject) AsArray() *IndexedObject {
	if self.isArray() {
		return self
	}
	if self.isMap() {
		self.mapp = nil
	}
	self.arr = make([]interface{}, 0)
	return self
}

func (self *IndexedObject) AsMap() *IndexedObject {
	if self.isMap() {
		return self
	}
	if self.isArray() {
		self.arr = nil
	}
	self.mapp = make(map[interface{}]interface{})
	return self
}

func (self *IndexedObject) Clear() *IndexedObject {
	if self.arr != nil {
		self.arr = make([]interface{}, 0)
	} else if (self.mapp != nil) {
		self.mapp = make(map[interface{}]interface{})
	}
	return self
}

func (self *IndexedObject) Append(v interface{}) *IndexedObject {
	self.arr = append(self.arr, v)
	return self
}

func (self *IndexedObject) At(idx int) interface{} {
	return self.arr[idx]
}

func (self *IndexedObject) Get(k interface{}) interface{} {
	return self.mapp[k]
}

func (self *IndexedObject) Set(k interface{}, v interface{}) *IndexedObject {
	self.mapp[k] = v
	return self
}

func (self *IndexedObject) Foreach(proc func(v interface{}, k interface{})) {
	if self.arr != nil {
		for idx, e := range self.arr {
			proc(e, idx)
		}
	} else if self.mapp != nil {
		for k, v := range self.mapp {
			proc(v, k)
		}
	}
}
