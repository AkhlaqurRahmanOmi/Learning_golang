package main

import "fmt"

var a = 10

func main() {
	/**
	example of variable shadowing
	age := 30

	if age > 18 {
		a := 47
		fmt.Println(a)
	}
	fmt.Println(a)**/
	fmt.Println(a)
}
func init() {
	println(a)
	a = 20
}
