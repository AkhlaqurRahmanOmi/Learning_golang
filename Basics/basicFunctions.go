package main

import "fmt"

func main() {

	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(add(getNumber(a, b)))
}

func add(a int, b int) int {
	//getNumber(a, b)
	return a + b
}
func getNumber(a int, b int) (int, int) {
	return a, b
}
