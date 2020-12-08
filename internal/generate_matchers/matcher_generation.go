package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// Generate matchers:
//
//     go generate github.com/petergtz/pegomock/internal/generate_matchers

//go:generate go run matcher_generation.go
//go:generate go fmt ../../matcher_factories.go
//go:generate go fmt ../../ginkgo_compatible/matchers.go

func main() {
	mustWriteFile("../../matcher_factories.go", GenerateDefaultMatchersFile())
	mustWriteFile("../../ginkgo_compatible/matchers.go", GenerateGinkgoMatchersFile())
}

func mustWriteFile(path string, contents string) {
	err := ioutil.WriteFile(path, []byte(contents), 0644)
	if err != nil {
		panic(err)
	}
}

func GenerateDefaultMatchersFile() string {
	contents := `package pegomock

import (
	"reflect"
)
`

	for _, kind := range primitiveKinds {
		contents += fmt.Sprintf(`
func Eq%[1]s(value %[2]s) %[2]s {
	RegisterMatcher(&EqMatcher{Value: value})
	return %[4]s
}

func NotEq%[1]s(value %[2]s) %[2]s {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return %[4]s
}

func Any%[1]s() %[2]s {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf(%[3]s)))
	return %[4]s
}

func %[1]sThat(matcher ArgumentMatcher) %[2]s {
	RegisterMatcher(matcher)
	return %[4]s
}

func Eq%[1]sSlice(value []%[2]s) []%[2]s {
	RegisterMatcher(&EqMatcher{Value: value})
	return nil
}

func NotEq%[1]sSlice(value []%[2]s) []%[2]s {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return nil
}

func Any%[1]sSlice() []%[2]s {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf(%[3]s))))
	return nil
}

func %[1]sSliceThat(matcher ArgumentMatcher) []%[2]s {
	RegisterMatcher(matcher)
	return nil
}
`, strings.Title(kind.String()), kind.String(), exampleValue(kind), zeroValue(kind))
	}

	// hard-coding this for now as interface{} overall works slightly different than other types.
	return contents + `
func EqInterface(value interface{}) interface{} {
	RegisterMatcher(&EqMatcher{Value: value})
	return nil
}

func NotEqInterface(value interface{}) interface{} {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return nil
}

func AnyInterface() interface{} {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((*interface{})(nil)).Elem()))
	return nil
}

func InterfaceThat(matcher ArgumentMatcher) interface{} {
	RegisterMatcher(matcher)
	return nil
}

func EqInterfaceSlice(value []interface{}) []interface{} {
	RegisterMatcher(&EqMatcher{Value: value})
	return nil
}

func NotEqInterfaceSlice(value []interface{}) []interface{} {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return nil
}

func AnyInterfaceSlice() []interface{} {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((*interface{})(nil)).Elem())))
	return nil
}

func InterfaceSliceThat(matcher ArgumentMatcher) []interface{} {
	RegisterMatcher(matcher)
	return nil
}
`
}

func GenerateGinkgoMatchersFile() string {
	contents := `package mock

import (
	"github.com/petergtz/pegomock"
)

var (`

	for _, kind := range append(primitiveKinds, reflect.Interface) {
		contents += fmt.Sprintf(`
	Eq%[1]s = pegomock.Eq%[1]s
	NotEq%[1]s = pegomock.NotEq%[1]s
	Any%[1]s = pegomock.Any%[1]s
	%[1]sThat = pegomock.%[1]sThat
	Eq%[1]sSlice = pegomock.Eq%[1]sSlice
	NotEq%[1]sSlice = pegomock.NotEq%[1]sSlice
	Any%[1]sSlice = pegomock.Any%[1]sSlice
	%[1]sSliceThat = pegomock.%[1]sSliceThat
`, strings.Title(kind.String()))
	}

	return contents + `
	Times   = pegomock.Times
	AtLeast = pegomock.AtLeast
	AtMost  = pegomock.AtMost
	Never   = pegomock.Never
	Once    = pegomock.Once
	Twice   = pegomock.Twice
)
`
}

var primitiveKinds = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Uintptr,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
	reflect.String,
}

// TODO generate: chan, func matchers

func zeroValue(kind reflect.Kind) string {
	switch {
	case kind == reflect.Bool:
		return `false`
	case reflect.Int <= kind && kind <= reflect.Complex128:
		return `0`
	case kind == reflect.String:
		return `""`
	default:
		return `nil`
	}
}

func exampleValue(kind reflect.Kind) string {
	if kind == reflect.Bool || kind == reflect.Int || kind == reflect.String {
		return zeroValue(kind)
	}
	return fmt.Sprintf("(%s)(%s)", kind.String(), zeroValue(kind))
}
