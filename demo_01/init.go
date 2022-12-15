package main

import (
	"fmt"
	"os/exec"
)

func init() {
	fmt.Println("hello world")
}

func main1() {
	cmd := exec.Command("go", "version")
	out, _ := cmd.CombinedOutput() //下划线处理异常和错误的情况
	fmt.Printf("%s", out)
}

func main() {

}
