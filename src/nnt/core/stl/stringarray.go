package stl

type StringSlice []string

type implStringArray struct {
	StringSlice
}

func StringArray(arr []string) *implStringArray {
	return &implStringArray{arr}
}

func (self *StringSlice) Length() int {
	return len(*self)
}

type FnStringArrayEqual func(idx int, e string) bool

func (self *StringSlice) Query(compr FnStringArrayEqual) (string, bool) {
	for idx, e := range *self {
		if compr(idx, e) {
			return e, true
		}
	}
	return "", false
}
