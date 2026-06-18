package main

func main() {
	// int
	const (
		minUint8 uint8 = 0
		maxUint8 uint8 = 255
		minInt8  int8  = -128
		maxInt8  int8  = 127

		minUint16 uint16 = 0
		maxUint16 uint16 = 65535
		minInt16  int16  = -32768
		maxInt16  int16  = 32767

		minUint32 uint32 = 0
		maxUint32 uint32 = 4294967295
		minInt32  int32  = -2147483648
		maxInt32  int32  = 2147483647

		minUint64 uint64 = 0
		maxUint64 uint64 = 18446744073709551615
		minInt64  int64  = -9223372036854775808
		maxInt64  int64  = 9223372036854775807
	)

	// float
	const (
		minFloat32 float32 = -3.4028235e+38
		maxFloat32 float32 = 3.4028235e+38

		minFloat64 float64 = -1.7976931348623157e+308
		maxFloat64 float64 = 1.7976931348623157e+308
	)

	// complex
	const (
		minComplex64 complex64 = -3.4028235e+38 + -3.4028235e+38i
		maxComplex64 complex64 = 3.4028235e+38 + 3.4028235e+38i

		minComplex128 complex128 = -1.7976931348623157e+308 + -1.7976931348623157e+308i
		maxComplex128 complex128 = 1.7976931348623157e+308 + 1.7976931348623157e+308i
	)

	// bool
	const (
		boolFalse bool = false
		boolTrue  bool = true
	)

	// string
	const (
		s1 string = "aaa"
	)

	// rune - aliases para int32
	const (
		rune1 rune = 'a'
		rune2 rune = 'b'
	)

	// byte - alias para uint8
	const (
		byte1 byte = 'a'
		byte2 byte = 'b'
	)

	// valores padrões
	var (
		defaultInt     int       // 0
		defaultFloat   float64   // 0
		defaultComplex complex64 // 0
		defaultBool    bool      // false
		defaultString  string    // ""
		defaultRune    rune      // 0
		defaultByte    byte      // 0
		defaultError   error     // nil
	)

	println(defaultInt, defaultFloat, defaultComplex, defaultBool, defaultString, defaultRune, defaultByte, defaultError)
}
