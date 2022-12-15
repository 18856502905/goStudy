package MyPackage

import "fmt"

func init() {
	fmt.Println("hello mypackage")
}

func Run() { //大写外部可以使用
	fmt.Println("mypackage run")
}
