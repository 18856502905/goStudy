package main

import (
	"demo_01/MyPackage"
	"fmt"
)

func main_12() {
	a, b, _ := 1, 2, 3 // _表示匿名变量 匿名变量不占内存也不会分配内存
	fmt.Println(a, b)
}

func returngo() (int, int, int) {
	return 1, 2, 3
}

func main() {
	a, b, _ := returngo()
	MyPackage.Run()
	fmt.Println(a, b)
}
