package core

import "nnt"

func QueryObject(arr *nnt.Array, filter func(any *nnt.Any, idx int) bool) *nnt.Any {
	for idx := range *arr {
		if filter(&(*arr)[idx], idx) {
			return &(*arr)[idx]
		}
	}
	return nil
}

func Convert(arr *nnt.Array, to func(any *nnt.Any, idx int) *nnt.Any, skipnull bool) []*nnt.Any {
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
