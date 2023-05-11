package mock

import (
	"github.com/petergtz/pegomock/v4"
)

var (
	EqBool         = pegomock.EqBool         // Deprecated: Use Eq[T any](value T) instead.
	NotEqBool      = pegomock.NotEqBool      // Deprecated: Use NotEq[T any](value T) instead.
	AnyBool        = pegomock.AnyBool        // Deprecated: Use Any[T any]() instead.
	BoolThat       = pegomock.BoolThat       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqBoolSlice    = pegomock.EqBoolSlice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqBoolSlice = pegomock.NotEqBoolSlice // Deprecated: Use NotEq[T any](value T) instead.
	AnyBoolSlice   = pegomock.AnyBoolSlice   // Deprecated: Use Any[T any]() instead.
	BoolSliceThat  = pegomock.BoolSliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqInt         = pegomock.EqInt         // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt      = pegomock.NotEqInt      // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt        = pegomock.AnyInt        // Deprecated: Use Any[T any]() instead.
	IntThat       = pegomock.IntThat       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqIntSlice    = pegomock.EqIntSlice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqIntSlice = pegomock.NotEqIntSlice // Deprecated: Use NotEq[T any](value T) instead.
	AnyIntSlice   = pegomock.AnyIntSlice   // Deprecated: Use Any[T any]() instead.
	IntSliceThat  = pegomock.IntSliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqInt8         = pegomock.EqInt8         // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt8      = pegomock.NotEqInt8      // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt8        = pegomock.AnyInt8        // Deprecated: Use Any[T any]() instead.
	Int8That       = pegomock.Int8That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqInt8Slice    = pegomock.EqInt8Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt8Slice = pegomock.NotEqInt8Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt8Slice   = pegomock.AnyInt8Slice   // Deprecated: Use Any[T any]() instead.
	Int8SliceThat  = pegomock.Int8SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqInt16         = pegomock.EqInt16         // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt16      = pegomock.NotEqInt16      // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt16        = pegomock.AnyInt16        // Deprecated: Use Any[T any]() instead.
	Int16That       = pegomock.Int16That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqInt16Slice    = pegomock.EqInt16Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt16Slice = pegomock.NotEqInt16Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt16Slice   = pegomock.AnyInt16Slice   // Deprecated: Use Any[T any]() instead.
	Int16SliceThat  = pegomock.Int16SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqInt32         = pegomock.EqInt32         // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt32      = pegomock.NotEqInt32      // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt32        = pegomock.AnyInt32        // Deprecated: Use Any[T any]() instead.
	Int32That       = pegomock.Int32That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqInt32Slice    = pegomock.EqInt32Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt32Slice = pegomock.NotEqInt32Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt32Slice   = pegomock.AnyInt32Slice   // Deprecated: Use Any[T any]() instead.
	Int32SliceThat  = pegomock.Int32SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqInt64         = pegomock.EqInt64         // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt64      = pegomock.NotEqInt64      // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt64        = pegomock.AnyInt64        // Deprecated: Use Any[T any]() instead.
	Int64That       = pegomock.Int64That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqInt64Slice    = pegomock.EqInt64Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqInt64Slice = pegomock.NotEqInt64Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyInt64Slice   = pegomock.AnyInt64Slice   // Deprecated: Use Any[T any]() instead.
	Int64SliceThat  = pegomock.Int64SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqUint         = pegomock.EqUint         // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint      = pegomock.NotEqUint      // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint        = pegomock.AnyUint        // Deprecated: Use Any[T any]() instead.
	UintThat       = pegomock.UintThat       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqUintSlice    = pegomock.EqUintSlice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqUintSlice = pegomock.NotEqUintSlice // Deprecated: Use NotEq[T any](value T) instead.
	AnyUintSlice   = pegomock.AnyUintSlice   // Deprecated: Use Any[T any]() instead.
	UintSliceThat  = pegomock.UintSliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqUint8         = pegomock.EqUint8         // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint8      = pegomock.NotEqUint8      // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint8        = pegomock.AnyUint8        // Deprecated: Use Any[T any]() instead.
	Uint8That       = pegomock.Uint8That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqUint8Slice    = pegomock.EqUint8Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint8Slice = pegomock.NotEqUint8Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint8Slice   = pegomock.AnyUint8Slice   // Deprecated: Use Any[T any]() instead.
	Uint8SliceThat  = pegomock.Uint8SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqUint16         = pegomock.EqUint16         // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint16      = pegomock.NotEqUint16      // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint16        = pegomock.AnyUint16        // Deprecated: Use Any[T any]() instead.
	Uint16That       = pegomock.Uint16That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqUint16Slice    = pegomock.EqUint16Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint16Slice = pegomock.NotEqUint16Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint16Slice   = pegomock.AnyUint16Slice   // Deprecated: Use Any[T any]() instead.
	Uint16SliceThat  = pegomock.Uint16SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqUint32         = pegomock.EqUint32         // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint32      = pegomock.NotEqUint32      // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint32        = pegomock.AnyUint32        // Deprecated: Use Any[T any]() instead.
	Uint32That       = pegomock.Uint32That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqUint32Slice    = pegomock.EqUint32Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint32Slice = pegomock.NotEqUint32Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint32Slice   = pegomock.AnyUint32Slice   // Deprecated: Use Any[T any]() instead.
	Uint32SliceThat  = pegomock.Uint32SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqUint64         = pegomock.EqUint64         // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint64      = pegomock.NotEqUint64      // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint64        = pegomock.AnyUint64        // Deprecated: Use Any[T any]() instead.
	Uint64That       = pegomock.Uint64That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqUint64Slice    = pegomock.EqUint64Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqUint64Slice = pegomock.NotEqUint64Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyUint64Slice   = pegomock.AnyUint64Slice   // Deprecated: Use Any[T any]() instead.
	Uint64SliceThat  = pegomock.Uint64SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqUintptr         = pegomock.EqUintptr         // Deprecated: Use Eq[T any](value T) instead.
	NotEqUintptr      = pegomock.NotEqUintptr      // Deprecated: Use NotEq[T any](value T) instead.
	AnyUintptr        = pegomock.AnyUintptr        // Deprecated: Use Any[T any]() instead.
	UintptrThat       = pegomock.UintptrThat       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqUintptrSlice    = pegomock.EqUintptrSlice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqUintptrSlice = pegomock.NotEqUintptrSlice // Deprecated: Use NotEq[T any](value T) instead.
	AnyUintptrSlice   = pegomock.AnyUintptrSlice   // Deprecated: Use Any[T any]() instead.
	UintptrSliceThat  = pegomock.UintptrSliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqFloat32         = pegomock.EqFloat32         // Deprecated: Use Eq[T any](value T) instead.
	NotEqFloat32      = pegomock.NotEqFloat32      // Deprecated: Use NotEq[T any](value T) instead.
	AnyFloat32        = pegomock.AnyFloat32        // Deprecated: Use Any[T any]() instead.
	Float32That       = pegomock.Float32That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqFloat32Slice    = pegomock.EqFloat32Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqFloat32Slice = pegomock.NotEqFloat32Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyFloat32Slice   = pegomock.AnyFloat32Slice   // Deprecated: Use Any[T any]() instead.
	Float32SliceThat  = pegomock.Float32SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqFloat64         = pegomock.EqFloat64         // Deprecated: Use Eq[T any](value T) instead.
	NotEqFloat64      = pegomock.NotEqFloat64      // Deprecated: Use NotEq[T any](value T) instead.
	AnyFloat64        = pegomock.AnyFloat64        // Deprecated: Use Any[T any]() instead.
	Float64That       = pegomock.Float64That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqFloat64Slice    = pegomock.EqFloat64Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqFloat64Slice = pegomock.NotEqFloat64Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyFloat64Slice   = pegomock.AnyFloat64Slice   // Deprecated: Use Any[T any]() instead.
	Float64SliceThat  = pegomock.Float64SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqComplex64         = pegomock.EqComplex64         // Deprecated: Use Eq[T any](value T) instead.
	NotEqComplex64      = pegomock.NotEqComplex64      // Deprecated: Use NotEq[T any](value T) instead.
	AnyComplex64        = pegomock.AnyComplex64        // Deprecated: Use Any[T any]() instead.
	Complex64That       = pegomock.Complex64That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqComplex64Slice    = pegomock.EqComplex64Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqComplex64Slice = pegomock.NotEqComplex64Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyComplex64Slice   = pegomock.AnyComplex64Slice   // Deprecated: Use Any[T any]() instead.
	Complex64SliceThat  = pegomock.Complex64SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqComplex128         = pegomock.EqComplex128         // Deprecated: Use Eq[T any](value T) instead.
	NotEqComplex128      = pegomock.NotEqComplex128      // Deprecated: Use NotEq[T any](value T) instead.
	AnyComplex128        = pegomock.AnyComplex128        // Deprecated: Use Any[T any]() instead.
	Complex128That       = pegomock.Complex128That       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqComplex128Slice    = pegomock.EqComplex128Slice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqComplex128Slice = pegomock.NotEqComplex128Slice // Deprecated: Use NotEq[T any](value T) instead.
	AnyComplex128Slice   = pegomock.AnyComplex128Slice   // Deprecated: Use Any[T any]() instead.
	Complex128SliceThat  = pegomock.Complex128SliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqString         = pegomock.EqString         // Deprecated: Use Eq[T any](value T) instead.
	NotEqString      = pegomock.NotEqString      // Deprecated: Use NotEq[T any](value T) instead.
	AnyString        = pegomock.AnyString        // Deprecated: Use Any[T any]() instead.
	StringThat       = pegomock.StringThat       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqStringSlice    = pegomock.EqStringSlice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqStringSlice = pegomock.NotEqStringSlice // Deprecated: Use NotEq[T any](value T) instead.
	AnyStringSlice   = pegomock.AnyStringSlice   // Deprecated: Use Any[T any]() instead.
	StringSliceThat  = pegomock.StringSliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	EqInterface         = pegomock.EqInterface         // Deprecated: Use Eq[T any](value T) instead.
	NotEqInterface      = pegomock.NotEqInterface      // Deprecated: Use NotEq[T any](value T) instead.
	AnyInterface        = pegomock.AnyInterface        // Deprecated: Use Any[T any]() instead.
	InterfaceThat       = pegomock.InterfaceThat       // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.
	EqInterfaceSlice    = pegomock.EqInterfaceSlice    // Deprecated: Use Eq[T any](value T) instead.
	NotEqInterfaceSlice = pegomock.NotEqInterfaceSlice // Deprecated: Use NotEq[T any](value T) instead.
	AnyInterfaceSlice   = pegomock.AnyInterfaceSlice   // Deprecated: Use Any[T any]() instead.
	InterfaceSliceThat  = pegomock.InterfaceSliceThat  // Deprecated: Use ArgThat[T any](matcher ArgumentMatcher) instead.

	Times   = pegomock.Times
	AtLeast = pegomock.AtLeast
	AtMost  = pegomock.AtMost
	Never   = pegomock.Never
	Once    = pegomock.Once
	Twice   = pegomock.Twice
)
