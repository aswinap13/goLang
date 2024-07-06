package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("===============")
	for i := range days {
		fmt.Printf(days[i])
	}
	fmt.Println("===============")
	for index, value := range days {
		fmt.Printf("index : %v, value : %v\n", index, value)
	}
	//COMMA okay syntax
	for _, value := range days {
		fmt.Printf("value : %v\n", value)
	}

	//Label
labell:
	//action
	goto labell

}
