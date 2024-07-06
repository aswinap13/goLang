package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch in Go")

	source := rand.NewSource(time.Now().UnixNano()) //Setting new Source
	rng := rand.New(source)                         //Seeding Random with a source

	diceNumber := rng.Intn(6) + 1
	fmt.Println("Dice value : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
	case 4:
		fmt.Println("Number is 4")
		fallthrough
	case 5:
		fmt.Println("Number is 5")
		fallthrough
	case 6:
		fmt.Println("Number is 6")
	}

}
