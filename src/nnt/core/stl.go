package core

import "nnt"

type _Array struct {
	store nnt.Array
}

type _Map struct {
	store nnt.Map
}

func Array(size int) *_Array {
	return &_Array{
		make(nnt.Array, size),
	}
}

func Map() *_Map {
	return &_Map{
		make(nnt.Map),
	}
}

func (self *_Array) Length() int {
	return len(self.store)
}

func (self *_Array) Append(r *nnt.Any) *_Array {
	self.store = append(self.store, *r)
	return self
}
