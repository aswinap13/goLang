package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi Welcome!")

	fmt.Println("Please enter a rating for our service")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	//Conversion

	actualRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error occured")
	} else {
		fmt.Println("the input is:", actualRating+1)
	}
}
