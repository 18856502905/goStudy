package main

import (
	"fmt"
	"os/exec"
)

func main231() { //go中的字符有2中uint8 以及rune
	var ch byte = 'A' //定义一个字符类型
	fmt.Println(ch)

}

func main232() {
	var a rune = '一'
	fmt.Printf("%T %v", a, a)
}

func main233() {
	var cmd string
	fmt.Scanf("%s", &cmd)
	fmt.Println("执行的命令:", cmd)
	result := exec.Command(cmd)
	result.Run()
}

func main() {
	var cmd string = "dinghao"
	for i := 0; i < len(cmd); i++ {
		fmt.Println(cmd[i])
	}
	for i, ch := range cmd {
		fmt.Printf("%d, %c\n", i, ch)
	}
}
