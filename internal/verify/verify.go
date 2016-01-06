package verify

import "fmt"

func NotNil(arg interface{}, message string, a ...interface{}) {
	if arg == nil {
		panic(fmt.Sprintf(message, a))
	}
}

func Argument(arg bool, message string, a ...interface{}) {
	if !arg {
		panic(fmt.Sprintf(message, a))
	}
}
