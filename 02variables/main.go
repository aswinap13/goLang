package main

import "fmt"

func main() {
	var username string = "aswin"
	fmt.Println(username)
	fmt.Printf("the type is : %T \n", username)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("the type is : %T \n", smallInt)

	var largeInt int = 23148
	fmt.Println(largeInt)
	fmt.Printf("the type is : %T \n", largeInt)

	var floatSmall float32 = 34.092323323
	fmt.Println(floatSmall)
	fmt.Printf("the type is : %T \n", floatSmall)

	var floatLarge float64 = 34.092323323
	fmt.Println(floatLarge)
	fmt.Printf("the type is : %T \n", floatLarge)

	//Not initialized
	var undInt string
	fmt.Println(undInt)
	fmt.Printf("the type is : %T \n", undInt)

	//No type
	var bag = "bag"
	fmt.Println(bag)
	fmt.Printf("the type is : %T \n", bag)

	//No Var, No type
	undInts := 9
	fmt.Println(undInts)
	fmt.Printf("the type is : %T \n", undInts)

}
