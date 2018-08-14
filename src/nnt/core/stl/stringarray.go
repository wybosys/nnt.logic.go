package stl

type StringArray []string

func (self *StringArray) Length() int {
	return len(*self)
}

type FnStringArrayEqual func(idx int, e string) bool

func (self *StringArray) Query(compr FnStringArrayEqual) (string, bool) {
	for idx, e := range *self {
		if compr(idx, e) {
			return e, true
		}
	}
	return "", false
}
