package main

import "fmt"

func main() {
	fmt.Println("Array in GO")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[2] = "Banana"
	// fruitList[4] = "ma"

	fmt.Println("The fruitList array : ", fruitList)
	fmt.Println("The length of fruitList array : ", len(fruitList))

	//Without Giving the length : This is SLICE
	var carList = []string{"audi", "bmw", "benz"}
	fmt.Println("The carList array : ", carList)
	fmt.Println("The length of carList array : ", len(carList))

	//Without Giving the length : 5
	var carList5 = [5]string{"audi", "bmw", "benz"}
	fmt.Println("The carList5 array : ", carList5)
	fmt.Println("The length of carList5 array : ", len(carList5))
}
