package main

import "fmt"

func main241() {
	var num1 int = 10
	var num2 int8 = int8(num1)
	fmt.Println(num2) //类型转换的时候不会修改原来的值
}
func main242() { //sprintf 使用
	var str string
	var a int = 100
	var b bool = false
	str = fmt.Sprintf("%d", a)
	str = fmt.Sprintf("%v", b)
	fmt.Println(str)

}

func main243() {
	var string1, string2, string3, string4 = "12", "12.4", "a", "false"
	var a int
	var b float32
	var c bool
	var d byte
	fmt.Sscanf(string1, "%d", &a)
	fmt.Sscanf(string2, "%f", &b)
	fmt.Sscanf(string3, "%t", &d)
	fmt.Sscanf(string4, "%c", &c)
}

func main244() { //类型的别名
	type u_int8 = uint8
	var num u_int8 = 10
	fmt.Println("%d", num)
}
