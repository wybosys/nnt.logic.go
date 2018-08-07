package core

import "nnt"

type prvArrayT struct {
}

var ArrayT = &prvArrayT{}

func (*prvArrayT) QueryObject(arr nnt.Any, filter func(any *nnt.Any, idx int) bool) *nnt.Any {
	ref := arr.(nnt.Array)
	for idx := range ref {
		if filter(&ref[idx], idx) {
			return &ref[idx]
		}
	}
	return nil
}

func (self *prvArrayT) Contains(arr nnt.Any, tgt nnt.Any, compar func(l *nnt.Any, r nnt.Any) bool) bool {
	return self.IndexOf(arr, tgt, compar) != -1
}

func (*prvArrayT) IndexOf(arr nnt.Any, tgt nnt.Any, compar func(l *nnt.Any, r nnt.Any) bool) int {
	ref := arr.(nnt.Array)
	if compar == nil {
		for idx := range ref {
			if ref[idx] == tgt {
				return idx
			}
		}
	} else {
		for idx := range ref {
			if compar(&ref[idx], tgt) {
				return idx
			}
		}
	}
	return -1
}

func (*prvArrayT) Convert(arr nnt.Any, to func(any *nnt.Any, idx int) nnt.Any, skipnull bool) nnt.Array {
	ref := arr.(nnt.Array)
	ret := make([]nnt.Any, len(ref))
	for idx := range ref {
		t := to(&ref[idx], idx)
		if t == nil && skipnull {
			continue
		}
		ret = append(ret, t)
	}
	return ret
}
