package pegomock

import (
	"reflect"
)

func EqBool(value bool) bool {
	RegisterMatcher(&EqMatcher{Value: value})
	return false
}

func NotEqBool(value bool) bool {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return false
}

func AnyBool() bool {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf(false)))
	return false
}

func AnyBoolSlice() []bool {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf(false))))
	return nil
}

func EqInt(value int) int {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqInt(value int) int {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyInt() int {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf(0)))
	return 0
}

func AnyIntSlice() []int {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf(0))))
	return nil
}

func EqInt8(value int8) int8 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqInt8(value int8) int8 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyInt8() int8 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((int8)(0))))
	return 0
}

func AnyInt8Slice() []int8 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((int8)(0)))))
	return nil
}

func EqInt16(value int16) int16 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqInt16(value int16) int16 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyInt16() int16 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((int16)(0))))
	return 0
}

func AnyInt16Slice() []int16 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((int16)(0)))))
	return nil
}

func EqInt32(value int32) int32 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqInt32(value int32) int32 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyInt32() int32 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((int32)(0))))
	return 0
}

func AnyInt32Slice() []int32 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((int32)(0)))))
	return nil
}

func EqInt64(value int64) int64 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqInt64(value int64) int64 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyInt64() int64 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((int64)(0))))
	return 0
}

func AnyInt64Slice() []int64 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((int64)(0)))))
	return nil
}

func EqUint(value uint) uint {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqUint(value uint) uint {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyUint() uint {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint)(0))))
	return 0
}

func AnyUintSlice() []uint {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint)(0)))))
	return nil
}

func EqUint8(value uint8) uint8 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqUint8(value uint8) uint8 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyUint8() uint8 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint8)(0))))
	return 0
}

func AnyUint8Slice() []uint8 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint8)(0)))))
	return nil
}

func EqUint16(value uint16) uint16 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqUint16(value uint16) uint16 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyUint16() uint16 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint16)(0))))
	return 0
}

func AnyUint16Slice() []uint16 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint16)(0)))))
	return nil
}

func EqUint32(value uint32) uint32 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqUint32(value uint32) uint32 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyUint32() uint32 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint32)(0))))
	return 0
}

func AnyUint32Slice() []uint32 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint32)(0)))))
	return nil
}

func EqUint64(value uint64) uint64 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqUint64(value uint64) uint64 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyUint64() uint64 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint64)(0))))
	return 0
}

func AnyUint64Slice() []uint64 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint64)(0)))))
	return nil
}

func EqUintptr(value uintptr) uintptr {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqUintptr(value uintptr) uintptr {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyUintptr() uintptr {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uintptr)(0))))
	return 0
}

func AnyUintptrSlice() []uintptr {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uintptr)(0)))))
	return nil
}

func EqFloat32(value float32) float32 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqFloat32(value float32) float32 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyFloat32() float32 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((float32)(0))))
	return 0
}

func AnyFloat32Slice() []float32 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((float32)(0)))))
	return nil
}

func EqFloat64(value float64) float64 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqFloat64(value float64) float64 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyFloat64() float64 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((float64)(0))))
	return 0
}

func AnyFloat64Slice() []float64 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((float64)(0)))))
	return nil
}

func EqComplex64(value complex64) complex64 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqComplex64(value complex64) complex64 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyComplex64() complex64 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((complex64)(0))))
	return 0
}

func AnyComplex64Slice() []complex64 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((complex64)(0)))))
	return nil
}

func EqComplex128(value complex128) complex128 {
	RegisterMatcher(&EqMatcher{Value: value})
	return 0
}

func NotEqComplex128(value complex128) complex128 {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return 0
}

func AnyComplex128() complex128 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((complex128)(0))))
	return 0
}

func AnyComplex128Slice() []complex128 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((complex128)(0)))))
	return nil
}

func EqString(value string) string {
	RegisterMatcher(&EqMatcher{Value: value})
	return ""
}

func NotEqString(value string) string {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return ""
}

func AnyString() string {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf("")))
	return ""
}

func EqStringSlice(value []string) []string {
	RegisterMatcher(&EqMatcher{Value: value})
	return nil
}

func NotEqStringSlice(value []string) []string {
	RegisterMatcher(&NotEqMatcher{Value: value})
	return nil
}

func AnyStringSlice() []string {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf(""))))
	return nil
}

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
