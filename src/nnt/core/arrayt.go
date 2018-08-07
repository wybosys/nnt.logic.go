package core

import "nnt"

type _ArrayT struct {
}

var ArrayT = &_ArrayT{}

func (*_ArrayT) QueryObject(arr *nnt.Array, filter func(any *nnt.Any, idx int) bool) *nnt.Any {
	for idx := range *arr {
		if filter(&(*arr)[idx], idx) {
			return &(*arr)[idx]
		}
	}
	return nil
}

func (*_ArrayT) Contains(arr *nnt.Array, tgt *nnt.Any, compar func(l *nnt.Any, r *nnt.Any) bool) bool {
	return ArrayT.IndexOf(arr, tgt, compar) != -1
}

func (*_ArrayT) IndexOf(arr *nnt.Array, tgt *nnt.Any, compar func(l *nnt.Any, r *nnt.Any) bool) int {
	for idx := range (*arr) {
		if compar(&(*arr)[idx], tgt) {
			return idx
		}
	}
	return -1
}

func (*_ArrayT) Convert(arr *nnt.Array, to func(any *nnt.Any, idx int) *nnt.Any, skipnull bool) []*nnt.Any {
	ret := make([]*nnt.Any, len(*arr))
	for idx := range (*arr) {
		t := to(&(*arr)[idx], idx)
		if t == nil && skipnull {
			continue
		}
		ret = append(ret, t)
	}
	return ret
}
