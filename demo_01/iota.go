package main

import "fmt"

type 学历 int

const ( //相当与C中的枚举类型
	小学 = 1 << iota
	初中
	高中
	大学
)

func main() {
	fmt.Println(小学)
	fmt.Println(大学)
}
