package main

import "fmt"

//%v 值的默认格式表示

func main() {
	cmd := "ak47"
	fmt.Println("%T", "%v", "+v", "%#v", cmd, cmd, cmd, cmd)
}
