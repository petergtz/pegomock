package pegomock

import "fmt"

func checkNotNil(arg interface{}, message string, a ...interface{}) {
	if arg == nil {
		panic(fmt.Sprintf(message, a))
	}
}

func checkArgument(arg bool, message string, a ...interface{}) {
	if !arg {
		panic(fmt.Sprintf(message, a))
	}
}
