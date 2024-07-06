package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hi Welcome to our Pizza Shop"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the rating :")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("The type of input : %T", input)
}
