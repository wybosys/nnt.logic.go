package stl

type AnyArray []interface{}

func (self *AnyArray) Length() int {
	return len(*self)
}

type FnAnyArrayEqual func(idx int, e interface{}) bool

func (self *AnyArray) Query(compr FnAnyArrayEqual) (interface{}, bool) {
	for idx, e := range *self {
		if compr(idx, e) {
			return e, true
		}
	}
	return "", false
}
