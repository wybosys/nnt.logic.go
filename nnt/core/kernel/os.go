package kernel

import "os"

func ArgsContains(str string) bool {
	for idx := range os.Args {
		if os.Args[idx] == str {
			return true
		}
	}
	return false
}
