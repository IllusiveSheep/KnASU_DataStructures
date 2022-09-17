package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var uint1 uint8
	var uint2 uint16
	var uint3 uint32
	var uint4 uint64

	var int1 int8
	var int2 int16
	var int3 int32
	var int4 int64

	var float1 float32
	var float2 float64

	var complex1 complex64
	var complex2 complex128

	var byte1 byte
	var rune1 rune

	var string1 string

	auto1 := "super"
	auto2 := 123

	array1 := [3]string{"Конфетка", "Train", "世界"}
	array2 := []int{1, 2, 3, 4, 5}

	type struct1 struct {
		excel1 string
		excel2 int
	}

	var structREAL struct1 = struct1{"tesla", 3}

	dict1 := make(map[string]string)
	dict1["alex"] = "tennis"
	dict1["maxim"] = "football"

	set := make(map[string]bool)
	set["Fee"] = true
	set["Fi"] = false
	delete(set, "Fee")
	size := len(set)
	exists := set["Foo"]

	fmt.Printf("\n", reflect.TypeOf(uint1))
	fmt.Printf("\n", reflect.TypeOf(uint2))
	fmt.Printf("\n", reflect.TypeOf(uint3))
	fmt.Printf("\n", reflect.TypeOf(uint4))

	fmt.Printf("\n", reflect.TypeOf(int1))
	fmt.Printf("\n", reflect.TypeOf(int2))
	fmt.Printf("\n", reflect.TypeOf(int3))
	fmt.Printf("\n", reflect.TypeOf(int4))

	fmt.Printf("\n", reflect.TypeOf(float1))
	fmt.Printf("\n", reflect.TypeOf(float2))

	fmt.Printf("\n", reflect.TypeOf(complex1))
	fmt.Printf("\n", reflect.TypeOf(complex2))

	fmt.Printf("\n", reflect.TypeOf(byte1))
	fmt.Printf("\n", reflect.TypeOf(rune1))

	fmt.Printf("\n", reflect.TypeOf(string1))

	fmt.Printf("\n", reflect.TypeOf(auto1))
	fmt.Printf("\n", reflect.TypeOf(auto2))

	fmt.Printf("\n", reflect.TypeOf(array1))
	fmt.Printf("\n", reflect.TypeOf(array2))

	fmt.Printf("\n", reflect.TypeOf(structREAL))

	fmt.Printf("\n")

	integer := 10242444
	fmt.Printf("%08b\n", integer)

	fmt.Printf("\n", &integer)
	fmt.Printf("\n", *&integer)
	fmt.Print("\n")

	fmt.Printf("Offsetof %v \n", unsafe.Offsetof(structREAL.excel1))

	for k := range set {
		fmt.Println(k)
	}

	fmt.Println(size)
	fmt.Println(exists)

	fmt.Println(array1)
	emoji := []rune("привет😀")
	fmt.Println(emoji)
	for i := 0; i < len(emoji); i++ {
		fmt.Println(emoji[i], string(emoji[i]))
	}

}
