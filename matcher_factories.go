package pegomock

import (
	"reflect"
)

// Deprecated: Use Any[T any]() instead.
func AnyBool() bool {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf(false)))
	return false
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func BoolThat(matcher ArgumentMatcher) bool {
	RegisterMatcher(matcher)
	return false
}

// Deprecated: Use Any[T any]() instead.
func AnyBoolSlice() []bool {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf(false))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func BoolSliceThat(matcher ArgumentMatcher) []bool {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyInt() int {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf(0)))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func IntThat(matcher ArgumentMatcher) int {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyIntSlice() []int {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf(0))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func IntSliceThat(matcher ArgumentMatcher) []int {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyInt8() int8 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((int8)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Int8That(matcher ArgumentMatcher) int8 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyInt8Slice() []int8 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((int8)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Int8SliceThat(matcher ArgumentMatcher) []int8 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyInt16() int16 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((int16)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Int16That(matcher ArgumentMatcher) int16 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyInt16Slice() []int16 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((int16)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Int16SliceThat(matcher ArgumentMatcher) []int16 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyInt32() int32 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((int32)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Int32That(matcher ArgumentMatcher) int32 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyInt32Slice() []int32 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((int32)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Int32SliceThat(matcher ArgumentMatcher) []int32 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyInt64() int64 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((int64)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Int64That(matcher ArgumentMatcher) int64 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyInt64Slice() []int64 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((int64)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Int64SliceThat(matcher ArgumentMatcher) []int64 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyUint() uint {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func UintThat(matcher ArgumentMatcher) uint {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyUintSlice() []uint {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func UintSliceThat(matcher ArgumentMatcher) []uint {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyUint8() uint8 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint8)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Uint8That(matcher ArgumentMatcher) uint8 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyUint8Slice() []uint8 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint8)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Uint8SliceThat(matcher ArgumentMatcher) []uint8 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyUint16() uint16 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint16)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Uint16That(matcher ArgumentMatcher) uint16 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyUint16Slice() []uint16 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint16)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Uint16SliceThat(matcher ArgumentMatcher) []uint16 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyUint32() uint32 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint32)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Uint32That(matcher ArgumentMatcher) uint32 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyUint32Slice() []uint32 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint32)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Uint32SliceThat(matcher ArgumentMatcher) []uint32 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyUint64() uint64 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uint64)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Uint64That(matcher ArgumentMatcher) uint64 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyUint64Slice() []uint64 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uint64)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Uint64SliceThat(matcher ArgumentMatcher) []uint64 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyUintptr() uintptr {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((uintptr)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func UintptrThat(matcher ArgumentMatcher) uintptr {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyUintptrSlice() []uintptr {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((uintptr)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func UintptrSliceThat(matcher ArgumentMatcher) []uintptr {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyFloat32() float32 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((float32)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Float32That(matcher ArgumentMatcher) float32 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyFloat32Slice() []float32 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((float32)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Float32SliceThat(matcher ArgumentMatcher) []float32 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyFloat64() float64 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((float64)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Float64That(matcher ArgumentMatcher) float64 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyFloat64Slice() []float64 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((float64)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Float64SliceThat(matcher ArgumentMatcher) []float64 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyComplex64() complex64 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((complex64)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Complex64That(matcher ArgumentMatcher) complex64 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyComplex64Slice() []complex64 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((complex64)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Complex64SliceThat(matcher ArgumentMatcher) []complex64 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyComplex128() complex128 {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((complex128)(0))))
	return 0
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Complex128That(matcher ArgumentMatcher) complex128 {
	RegisterMatcher(matcher)
	return 0
}

// Deprecated: Use Any[T any]() instead.
func AnyComplex128Slice() []complex128 {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((complex128)(0)))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func Complex128SliceThat(matcher ArgumentMatcher) []complex128 {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyString() string {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf("")))
	return ""
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func StringThat(matcher ArgumentMatcher) string {
	RegisterMatcher(matcher)
	return ""
}

// Deprecated: Use Any[T any]() instead.
func AnyStringSlice() []string {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf(""))))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func StringSliceThat(matcher ArgumentMatcher) []string {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyInterface() interface{} {
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf((*interface{})(nil)).Elem()))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func InterfaceThat(matcher ArgumentMatcher) interface{} {
	RegisterMatcher(matcher)
	return nil
}

// Deprecated: Use Any[T any]() instead.
func AnyInterfaceSlice() []interface{} {
	RegisterMatcher(NewAnyMatcher(reflect.SliceOf(reflect.TypeOf((*interface{})(nil)).Elem())))
	return nil
}

// Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
func InterfaceSliceThat(matcher ArgumentMatcher) []interface{} {
	RegisterMatcher(matcher)
	return nil
}
