package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in GO")

	//1.First way
	var sliceOne = []string{"apple"}
	fmt.Println("The slice is : ", sliceOne)
	fmt.Println("THe length of slice : ", len(sliceOne))
	fmt.Printf("The type %T\n", sliceOne)

	//APPENDING NEW VALUES
	sliceOne = append(sliceOne, "Banana", "sugar", "new")
	fmt.Println("after appending ======")
	fmt.Println("The slice is : ", sliceOne)
	fmt.Println("THe length of slice : ", len(sliceOne))
	fmt.Printf("The type %T\n", sliceOne)

	//Inclusive of 1, excluded of 3, append to array : so remove 0th and elements after 3(inclusive)
	// sliceOne = append(sliceOne[1:3])
	fmt.Println("after appending ======")
	fmt.Println("The slice is : ", sliceOne)

	//2.The second way for slices
	highScores := make([]int, 4) //allocates integer array of 4 int length
	highScores[0] = 0
	highScores[1] = 1
	highScores[2] = 2
	highScores[3] = 3

	// highScores[4] = 0 //This will throw, index out of bound

	highScores = append(highScores, 5, 4, 6) //This will allocate new memory for slice

	fmt.Println("the highScores slice : ", highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)


	cars := []string{}
	
	cars = append(cars, "audi","bmw","benz","porsche","vw")
	fmt.Println(cars)
	index := 2; //Index whose value to be removed
	cars = append(cars[:index],cars[index+1:]...) //"..." : has to be added, as append method expects more params
	fmt.Println(cars)

}
