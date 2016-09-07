package pegomock

import (
	"reflect"

	"github.com/petergtz/pegomock/matcher"
)

func EqBool(value bool) bool {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return false
}

func AnyBool() bool {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Bool})
	return false
}

func AnyBoolSlice() []bool {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqInt(value int) int {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyInt() int {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Int})
	return 0
}

func AnyIntSlice() []int {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqInt8(value int8) int8 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyInt8() int8 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Int8})
	return 0
}

func AnyInt8Slice() []int8 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqInt16(value int16) int16 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyInt16() int16 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Int16})
	return 0
}

func AnyInt16Slice() []int16 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqInt32(value int32) int32 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyInt32() int32 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Int32})
	return 0
}

func AnyInt32Slice() []int32 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqInt64(value int64) int64 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyInt64() int64 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Int64})
	return 0
}

func AnyInt64Slice() []int64 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqUint(value uint) uint {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyUint() uint {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Uint})
	return 0
}

func AnyUintSlice() []uint {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqUint8(value uint8) uint8 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyUint8() uint8 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Uint8})
	return 0
}

func AnyUint8Slice() []uint8 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqUint16(value uint16) uint16 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyUint16() uint16 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Uint16})
	return 0
}

func AnyUint16Slice() []uint16 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqUint32(value uint32) uint32 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyUint32() uint32 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Uint32})
	return 0
}

func AnyUint32Slice() []uint32 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqUint64(value uint64) uint64 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyUint64() uint64 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Uint64})
	return 0
}

func AnyUint64Slice() []uint64 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqUintptr(value uintptr) uintptr {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyUintptr() uintptr {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Uintptr})
	return 0
}

func AnyUintptrSlice() []uintptr {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqFloat32(value float32) float32 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyFloat32() float32 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Float32})
	return 0
}

func AnyFloat32Slice() []float32 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqFloat64(value float64) float64 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyFloat64() float64 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Float64})
	return 0
}

func AnyFloat64Slice() []float64 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqComplex64(value complex64) complex64 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyComplex64() complex64 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Complex64})
	return 0
}

func AnyComplex64Slice() []complex64 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqComplex128(value complex128) complex128 {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

func AnyComplex128() complex128 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Complex128})
	return 0
}

func AnyComplex128Slice() []complex128 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

func EqString(value string) string {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return ""
}

func AnyString() string {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.String})
	return ""
}

func AnyStringSlice() []string {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}
