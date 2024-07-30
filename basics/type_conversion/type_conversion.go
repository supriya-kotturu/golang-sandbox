package conversion

import (
	"fmt"
	"unsafe"
)

func GetSizesOfInt() {
	fmt.Println("SIZES OF THE DATATYPES")
	var num_int8 int8 = 127
	var num_int16 int16 = 32767
	var num_int32 int32 = 2147483647
	var num_int64 int64 = 9223372036854775807

	fmt.Println("SIGNED INT")
	fmt.Printf("int_8 %d - size %d\n", num_int8, unsafe.Sizeof(num_int8))
	fmt.Printf("int_16 %d - size %d\n", num_int16, unsafe.Sizeof(num_int16))
	fmt.Printf("int_32 %d - size %d\n", num_int32, unsafe.Sizeof(num_int32))
	fmt.Printf("int_64 %d - size %d\n", num_int64, unsafe.Sizeof(num_int64))

	var num_uint8 uint8 = 127
	var num_uint16 uint16 = 32767
	var num_uint32 uint32 = 2147483647
	var num_uint64 uint64 = 9223372036854775807

	fmt.Println("SIGNED INT")
	fmt.Printf("uint_8 %d - size %d\n", num_uint8, unsafe.Sizeof(num_uint8))
	fmt.Printf("uint_16 %d - size %d\n", num_uint16, unsafe.Sizeof(num_uint16))
	fmt.Printf("uint_32 %d - size %d\n", num_uint32, unsafe.Sizeof(num_uint32))
	fmt.Printf("uint_64 %d - size %d\n", num_uint64, unsafe.Sizeof(num_uint64))
}

func GetSizeOfFloat() {
	var num_float32 float32 = 3.4028235e+38
	var num_float64 float64 = 1.7976931348623157e+308

	fmt.Println("SIZES OF FLOATS")
	fmt.Printf("float_32 %f - size %d\n", num_float32, unsafe.Sizeof(num_float32))
	fmt.Printf("float_64 %f - size %d\n", num_float64, unsafe.Sizeof(num_float64))
}

func GetSizeOfString() {
	var str string = "Hello, World!"
	var ru rune = 'r'

	fmt.Println("SIZE OF STRING")
	fmt.Printf("string %s - size %d\n", str, unsafe.Sizeof(str))

	fmt.Println("SIZE OF RUNE")
	fmt.Printf("rune %c - size %d\n", ru, unsafe.Sizeof(ru))
}

func TypeConversion() {
	// int conversion
	var a_int8 int = 127
	fmt.Printf("int_8 %d to int_16 %d - size %d\n", a_int8, int16(a_int8), unsafe.Sizeof(int16(a_int8)))
	fmt.Printf("int_8 %d to int_32 %d - size %d\n", a_int8, int32(a_int8), unsafe.Sizeof(int32(a_int8)))
	fmt.Printf("int_8 %d to int_64 %d- size %d\n", a_int8, int64(a_int8), unsafe.Sizeof(int64(a_int8)))
	fmt.Printf("int_8 %d to uint_8 %d - size %d\n", a_int8, uint8(a_int8), unsafe.Sizeof(uint8(a_int8)))
	fmt.Printf("int_8 %d to uint_16 %d - size %d\n", a_int8, uint16(a_int8), unsafe.Sizeof(uint16(a_int8)))
	fmt.Printf("int_8 %d to uint_32 %d - size %d\n", a_int8, uint32(a_int8), unsafe.Sizeof(uint32(a_int8)))
	fmt.Printf("int_8 %d to uint_64 %d - size %d\n", a_int8, uint64(a_int8), unsafe.Sizeof(uint64(a_int8)))

	// float to int
	var a_float32 float32 = 3.4028235e+38
	fmt.Printf("float_32 %f to int_32 %d - size %d\n", a_float32, int32(a_float32), unsafe.Sizeof(int32(a_float32)))

	// int to float
	var a_int32 int32 = 2147483647
	fmt.Printf("int_32 %d to float_32 %f - size %d\n", a_int32, float32(a_int32), unsafe.Sizeof(float32(a_int32)))

	// rune to int
	var a_rune rune = 'r'
	fmt.Printf("rune %c to int_32 %d - size %d\n", a_rune, int32(a_rune), unsafe.Sizeof(int32(a_rune)))

	// int to rune
	var a_int16 int16 = 32767
	fmt.Printf("int %d to rune %c - size %d\n", a_int16, rune(a_int16), unsafe.Sizeof(rune(a_int16)))

	// string to []byte
	var a_str string = "Hello, World!"
	fmt.Printf("string %s to []byte %v - size %d\n", a_str, []byte(a_str), unsafe.Sizeof([]byte(a_str)))

	// []byte to string
	var a_byte []byte = []byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	fmt.Printf("[]byte %v to string %s - size %d\n", a_byte, string(a_byte), unsafe.Sizeof(string(a_byte)))
}
