package main

import (
	"fmt"
	"unsafe"
)

func main211() {
	result := 3 > 2
	fmt.Println(result)
	fmt.Println(unsafe.Sizeof(result))
}

func main212() {
	var data1 uint = 10 //无符号整数只能用正数赋值， 负值会导致栈溢出
	var data2 int = 0b10
	var data3 int = 0x10
	fmt.Println(data1, data2, data3)
	fmt.Println(unsafe.Sizeof(data1))
}
