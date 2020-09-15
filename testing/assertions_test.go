package testing

import (
	"math/cmplx"
	"strconv"
	"testing"
)

var (
	StringValue string     = "To jest piękny dzień"
	IntValue    int        = 12345
	UIntValue   uint       = 246810
	Byte        byte       = 1<<8 - 1
	Rune        rune       = 1<<16 - 1
	MaxInt8     int8       = 1<<4 - 1
	MaxInt16    int16      = 1<<8 - 1
	MaxInt32    int32      = 1<<16 - 1
	MaxInt64    int64      = 1<<32 - 1
	MaxUInt8    uint8      = 1<<8 - 1
	MaxUInt16   uint16     = 1<<16 - 1
	MaxUInt32   uint32     = 1<<32 - 1
	MaxUInt64   uint64     = 1<<64 - 1
	Complex64   complex64  = 1.0i
	Complex128  complex128 = cmplx.Sqrt(-5 + 12i)
	Float32     float32
	Float64     float64
)

func TestAssertEqual(t *testing.T) {
	// convert to a float32
	f32, _ := strconv.ParseFloat("3.14159265", 32)
	Float32 = float32(f32)

	f64, _ := strconv.ParseFloat("3.14159265", 64)
	Float64 = f64

	unitTests := []struct {
		name     string
		expected interface{}
		data     Expected
	}{
		{"string success test", "To jest piękny dzień", Expected{T: t, Value: StringValue}},
		{"boolean success test", true, Expected{T: t, Value: true}},
		{"int success test", int(12345), Expected{T: t, Value: IntValue}},
		{"uint success test", uint(246810), Expected{T: t, Value: UIntValue}},
		{"int8 success test", int8(15), Expected{T: t, Value: MaxInt8}},
		{"int16 success test", int16(255), Expected{T: t, Value: MaxInt16}},
		{"int32 success test", int32(65535), Expected{T: t, Value: MaxInt32}},
		{"int64 success test", int64(4294967295), Expected{T: t, Value: MaxInt64}},
		{"uint8 success test", uint8(255), Expected{T: t, Value: MaxUInt8}},
		{"uint16 success test", uint16(65535), Expected{T: t, Value: MaxUInt16}},
		{"uint32 success test", uint32(4294967295), Expected{T: t, Value: MaxUInt32}},
		{"uint64 success test", uint64(18446744073709551615), Expected{T: t, Value: MaxUInt64}},
		{"byte success test", byte(255), Expected{T: t, Value: Byte}},
		{"rune success test", rune(65535), Expected{T: t, Value: Rune}},
		{"complex64 success test", complex64(1.0i), Expected{T: t, Value: Complex64}},
		{"complex128 success test", complex128(cmplx.Sqrt(-5 + 12i)), Expected{T: t, Value: Complex128}},
		{"float32 success test", float32(f32), Expected{T: t, Value: Float32}},
		{"float64 success test", float64(f64), Expected{T: t, Value: Float64}},
	}

	for _, test := range unitTests {
		t.Run(test.name, func(t *testing.T) {
			if !test.data.AssertEqual(test.expected) {
				t.Fatalf("%s expected %v to equal %v", test.name, test.expected, test.data.Value)
			}
		})
	}
}
