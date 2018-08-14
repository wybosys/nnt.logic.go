package stl

import "nnt"

type prvArray struct {
	store nnt.Array
}

func Array(arr nnt.Any) *prvArray {
	return &prvArray{
		arr.(nnt.Array),
	}
}

func (self *prvArray) Length() int {
	return len(self.store)
}

func (self *prvArray) Append(r nnt.Any) *prvArray {
	self.store = append(self.store, r)
	return self
}
