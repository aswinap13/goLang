package main

import "fmt"

func main() {
	fmt.Println("Function in GO")
	greet()

	result := adder(6, 9)
	fmt.Println("the result is", result)

	proresult := proAdder(3, 5, 6, 0, 6, 9)
	fmt.Println("the result is", proresult)
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total = total + value
	}
	return total
}

func adder(first int, second int) int {
	return first + second
}
func greet() {
	fmt.Println("Greeting from GO")
}
