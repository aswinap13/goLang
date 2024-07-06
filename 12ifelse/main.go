package main

import "fmt"

func main() {
	fmt.Println("IfElse in Go")
	count := 10
	var result string

	if count < 10 {
		result = "LT 10"
	} else {
		result = "Not LT 10"
	}

	if(9%2 == 0) {
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}

	fmt.Println(result)

}
