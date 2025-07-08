package main

import (
	"fmt"
	"strings"
)

// this is integer to roman converter
func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	var count int = 0

	for i := 0; i < len(values); i++ {
		value := values[i]
		symbol := symbols[i]
		count += 1

		for num >= value {
			result.WriteString(symbol)
			num -= value
		}
		if num == 0 {
			break
		}

	}
	fmt.Println(count)
	return result.String()
}

//with the map

func intToRomanMap(num int) string {
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	// Define the order explicitly
	order := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var result strings.Builder
	for _, value := range order {
		fmt.Println(value)
		symbol := romanMap[value]
		count := num / value
		for i := 0; i < count; i++ {
			result.WriteString(symbol)
		}
		num %= value
		//fmt.Println(count)
		if num == 0 {
			break
		}

	}
	return result.String()
}

func main() {
	//fmt.Println(intToRoman(700))
	fmt.Println(intToRomanMap(700))
}
