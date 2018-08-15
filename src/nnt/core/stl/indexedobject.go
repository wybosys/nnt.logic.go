package stl

// 类似于js中的object，可以是array，也可以是map，但不能同时
// 为了降低go中构造非结构数据的难度
type IndexedObject struct {
	arr []interface{}
	mp  map[interface{}]interface{}
}

func (self *IndexedObject) Length() int {
	if self.arr != nil {
		return len(self.arr)
	}
	if self.mp != nil {
		return len(self.mp)
	}
	return 0
}

func (self *IndexedObject) isArray() bool {
	return self.arr != nil
}

func (self *IndexedObject) isMap() bool {
	return self.mp != nil
}

func (self *IndexedObject) Array() []interface{} {
	return self.arr
}

func (self *IndexedObject) Map() map[interface{}]interface{} {
	return self.mp
}

func (self *IndexedObject) AsArray() *IndexedObject {
	if self.isArray() {
		return self
	}
	if self.isMap() {
		self.mp = nil
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
	self.mp = make(map[interface{}]interface{})
	return self
}

func (self *IndexedObject) Clear() *IndexedObject {
	if self.arr != nil {
		self.arr = make([]interface{}, 0)
	} else if (self.mp != nil) {
		self.mp = make(map[interface{}]interface{})
	}
	return self
}

func (self *IndexedObject) Append(v interface{}) *IndexedObject {
	self.arr = append(self.arr, v)
	return self
}

func (self *IndexedObject) AtIndex(idx int) interface{} {
	return self.arr[idx]
}

func (self *IndexedObject) Get(k interface{}) interface{} {
	return self.mp[k]
}

func (self *IndexedObject) Set(k interface{}, v interface{}) *IndexedObject {
	self.mp[k] = v
	return self
}
