package mock

import (
	"github.com/petergtz/pegomock"
)

var (
	EqBool         = pegomock.EqBool
	NotEqBool      = pegomock.NotEqBool
	AnyBool        = pegomock.AnyBool
	BoolThat       = pegomock.BoolThat
	EqBoolSlice    = pegomock.EqBoolSlice
	NotEqBoolSlice = pegomock.NotEqBoolSlice
	AnyBoolSlice   = pegomock.AnyBoolSlice
	BoolSliceThat  = pegomock.BoolSliceThat

	EqInt         = pegomock.EqInt
	NotEqInt      = pegomock.NotEqInt
	AnyInt        = pegomock.AnyInt
	IntThat       = pegomock.IntThat
	EqIntSlice    = pegomock.EqIntSlice
	NotEqIntSlice = pegomock.NotEqIntSlice
	AnyIntSlice   = pegomock.AnyIntSlice
	IntSliceThat  = pegomock.IntSliceThat

	EqInt8         = pegomock.EqInt8
	NotEqInt8      = pegomock.NotEqInt8
	AnyInt8        = pegomock.AnyInt8
	Int8That       = pegomock.Int8That
	EqInt8Slice    = pegomock.EqInt8Slice
	NotEqInt8Slice = pegomock.NotEqInt8Slice
	AnyInt8Slice   = pegomock.AnyInt8Slice
	Int8SliceThat  = pegomock.Int8SliceThat

	EqInt16         = pegomock.EqInt16
	NotEqInt16      = pegomock.NotEqInt16
	AnyInt16        = pegomock.AnyInt16
	Int16That       = pegomock.Int16That
	EqInt16Slice    = pegomock.EqInt16Slice
	NotEqInt16Slice = pegomock.NotEqInt16Slice
	AnyInt16Slice   = pegomock.AnyInt16Slice
	Int16SliceThat  = pegomock.Int16SliceThat

	EqInt32         = pegomock.EqInt32
	NotEqInt32      = pegomock.NotEqInt32
	AnyInt32        = pegomock.AnyInt32
	Int32That       = pegomock.Int32That
	EqInt32Slice    = pegomock.EqInt32Slice
	NotEqInt32Slice = pegomock.NotEqInt32Slice
	AnyInt32Slice   = pegomock.AnyInt32Slice
	Int32SliceThat  = pegomock.Int32SliceThat

	EqInt64         = pegomock.EqInt64
	NotEqInt64      = pegomock.NotEqInt64
	AnyInt64        = pegomock.AnyInt64
	Int64That       = pegomock.Int64That
	EqInt64Slice    = pegomock.EqInt64Slice
	NotEqInt64Slice = pegomock.NotEqInt64Slice
	AnyInt64Slice   = pegomock.AnyInt64Slice
	Int64SliceThat  = pegomock.Int64SliceThat

	EqUint         = pegomock.EqUint
	NotEqUint      = pegomock.NotEqUint
	AnyUint        = pegomock.AnyUint
	UintThat       = pegomock.UintThat
	EqUintSlice    = pegomock.EqUintSlice
	NotEqUintSlice = pegomock.NotEqUintSlice
	AnyUintSlice   = pegomock.AnyUintSlice
	UintSliceThat  = pegomock.UintSliceThat

	EqUint8         = pegomock.EqUint8
	NotEqUint8      = pegomock.NotEqUint8
	AnyUint8        = pegomock.AnyUint8
	Uint8That       = pegomock.Uint8That
	EqUint8Slice    = pegomock.EqUint8Slice
	NotEqUint8Slice = pegomock.NotEqUint8Slice
	AnyUint8Slice   = pegomock.AnyUint8Slice
	Uint8SliceThat  = pegomock.Uint8SliceThat

	EqUint16         = pegomock.EqUint16
	NotEqUint16      = pegomock.NotEqUint16
	AnyUint16        = pegomock.AnyUint16
	Uint16That       = pegomock.Uint16That
	EqUint16Slice    = pegomock.EqUint16Slice
	NotEqUint16Slice = pegomock.NotEqUint16Slice
	AnyUint16Slice   = pegomock.AnyUint16Slice
	Uint16SliceThat  = pegomock.Uint16SliceThat

	EqUint32         = pegomock.EqUint32
	NotEqUint32      = pegomock.NotEqUint32
	AnyUint32        = pegomock.AnyUint32
	Uint32That       = pegomock.Uint32That
	EqUint32Slice    = pegomock.EqUint32Slice
	NotEqUint32Slice = pegomock.NotEqUint32Slice
	AnyUint32Slice   = pegomock.AnyUint32Slice
	Uint32SliceThat  = pegomock.Uint32SliceThat

	EqUint64         = pegomock.EqUint64
	NotEqUint64      = pegomock.NotEqUint64
	AnyUint64        = pegomock.AnyUint64
	Uint64That       = pegomock.Uint64That
	EqUint64Slice    = pegomock.EqUint64Slice
	NotEqUint64Slice = pegomock.NotEqUint64Slice
	AnyUint64Slice   = pegomock.AnyUint64Slice
	Uint64SliceThat  = pegomock.Uint64SliceThat

	EqUintptr         = pegomock.EqUintptr
	NotEqUintptr      = pegomock.NotEqUintptr
	AnyUintptr        = pegomock.AnyUintptr
	UintptrThat       = pegomock.UintptrThat
	EqUintptrSlice    = pegomock.EqUintptrSlice
	NotEqUintptrSlice = pegomock.NotEqUintptrSlice
	AnyUintptrSlice   = pegomock.AnyUintptrSlice
	UintptrSliceThat  = pegomock.UintptrSliceThat

	EqFloat32         = pegomock.EqFloat32
	NotEqFloat32      = pegomock.NotEqFloat32
	AnyFloat32        = pegomock.AnyFloat32
	Float32That       = pegomock.Float32That
	EqFloat32Slice    = pegomock.EqFloat32Slice
	NotEqFloat32Slice = pegomock.NotEqFloat32Slice
	AnyFloat32Slice   = pegomock.AnyFloat32Slice
	Float32SliceThat  = pegomock.Float32SliceThat

	EqFloat64         = pegomock.EqFloat64
	NotEqFloat64      = pegomock.NotEqFloat64
	AnyFloat64        = pegomock.AnyFloat64
	Float64That       = pegomock.Float64That
	EqFloat64Slice    = pegomock.EqFloat64Slice
	NotEqFloat64Slice = pegomock.NotEqFloat64Slice
	AnyFloat64Slice   = pegomock.AnyFloat64Slice
	Float64SliceThat  = pegomock.Float64SliceThat

	EqComplex64         = pegomock.EqComplex64
	NotEqComplex64      = pegomock.NotEqComplex64
	AnyComplex64        = pegomock.AnyComplex64
	Complex64That       = pegomock.Complex64That
	EqComplex64Slice    = pegomock.EqComplex64Slice
	NotEqComplex64Slice = pegomock.NotEqComplex64Slice
	AnyComplex64Slice   = pegomock.AnyComplex64Slice
	Complex64SliceThat  = pegomock.Complex64SliceThat

	EqComplex128         = pegomock.EqComplex128
	NotEqComplex128      = pegomock.NotEqComplex128
	AnyComplex128        = pegomock.AnyComplex128
	Complex128That       = pegomock.Complex128That
	EqComplex128Slice    = pegomock.EqComplex128Slice
	NotEqComplex128Slice = pegomock.NotEqComplex128Slice
	AnyComplex128Slice   = pegomock.AnyComplex128Slice
	Complex128SliceThat  = pegomock.Complex128SliceThat

	EqString         = pegomock.EqString
	NotEqString      = pegomock.NotEqString
	AnyString        = pegomock.AnyString
	StringThat       = pegomock.StringThat
	EqStringSlice    = pegomock.EqStringSlice
	NotEqStringSlice = pegomock.NotEqStringSlice
	AnyStringSlice   = pegomock.AnyStringSlice
	StringSliceThat  = pegomock.StringSliceThat

	EqInterface         = pegomock.EqInterface
	NotEqInterface      = pegomock.NotEqInterface
	AnyInterface        = pegomock.AnyInterface
	InterfaceThat       = pegomock.InterfaceThat
	EqInterfaceSlice    = pegomock.EqInterfaceSlice
	NotEqInterfaceSlice = pegomock.NotEqInterfaceSlice
	AnyInterfaceSlice   = pegomock.AnyInterfaceSlice
	InterfaceSliceThat  = pegomock.InterfaceSliceThat

	Times   = pegomock.Times
	AtLeast = pegomock.AtLeast
	AtMost  = pegomock.AtMost
	Never   = pegomock.Never
	Once    = pegomock.Once
	Twice   = pegomock.Twice
)
