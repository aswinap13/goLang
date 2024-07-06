package main

import "fmt"

func main() {

	fmt.Println("GO pointers!")

	// var firstInt *int
	// fmt.Println(firstInt)
	// fmt.Printf("type of var : %T\n",firstInt)

	var firstNumber int = 23
	var firstPointer *int = &firstNumber

	fmt.Println("number : ", firstNumber)
	fmt.Println("Pointer Address : ", firstPointer)
	fmt.Println("Pointer Value : ", *firstPointer)

	secondNumber := 24
	secondPointer := &secondNumber

	fmt.Println("number : ", secondNumber)
	fmt.Println("Pointer Address : ", secondPointer)
	fmt.Println("Pointer Value : ", *secondPointer)
}
