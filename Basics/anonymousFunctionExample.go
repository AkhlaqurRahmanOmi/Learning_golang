package main

import "fmt"

func addd(a, b int) {
	fmt.Println(a + b)
}
func main() {
	addd(1, 2)

	func(a int, b int) {
		fmt.Println(a + b)
	}(5, 7)
}

func init() {
	fmt.Println("First I am")
}
