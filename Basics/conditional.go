package main

import "fmt"

func main() {
	//age := 20
	//gender := "Male"

	//normal(age)
	//nornal2(age, gender)

	a := 3
	Switch(a)

}

func normal(age int) {
	if age > 18 {
		fmt.Println("You are eligible")
	} else if age < 18 {
		fmt.Println("You are not eligible")
	} else {
		fmt.Println("You are nothing")
	}
}

func nornal2(age int, gender string) {
	if age > 18 && gender == "Male" {
		fmt.Println("You are eligible")
	} else if age < 18 && gender == "Male" {
		fmt.Println("You are not eligible")
	} else {
		fmt.Println("You are nothing")
	}
}

func Switch(a int) {

	switch a {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Nothing")
	}
}
