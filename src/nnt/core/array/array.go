package array

import "nnt"

func QueryObject(arr nnt.Any, filter func(any *nnt.Any, idx int) bool) *nnt.Any {
	ref := arr.(nnt.Array)
	for idx := range ref {
		if filter(&ref[idx], idx) {
			return &ref[idx]
		}
	}
	return nil
}

func Contains(arr nnt.Any, tgt nnt.Any, compar func(l *nnt.Any, r nnt.Any) bool) bool {
	return IndexOf(arr, tgt, compar) != -1
}

func IndexOf(arr nnt.Any, tgt nnt.Any, compar func(l *nnt.Any, r nnt.Any) bool) int {
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

func Convert(arr nnt.Any, to func(any *nnt.Any, idx int) nnt.Any, skipnull bool) nnt.Array {
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
