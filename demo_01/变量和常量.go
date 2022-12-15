package main

import "fmt"

func main_1() {
	var N int = 10
	var J = 100 //自动推断
	fmt.Printf("%d %d\n", N, J)
}
func main_2() {
	/*var (
		friend  = "qq"
		friend1 = "mm"
		friend2 = "ll"
	)*/
	friend, friend1, friend2 := "qq", "mm", "ll"
	fmt.Println(friend + friend1 + friend2)
}

func main() {
	var a = 100
	var b = 10
	a, b = b, a
}
