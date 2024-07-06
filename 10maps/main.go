package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["Py"] = "python"

	fmt.Println("The languages", languages)
	fmt.Println("The JS language :", languages["JS"])

	//DELETE from a map/slice
	delete(languages,"JS")
	fmt.Println("The languages", languages)
	fmt.Println("The JS language :", languages["JS"])

	//ITERATE IN LOOP

for key, value := range(languages){
	fmt.Printf("For the key %v the value is : %v",key,value)
}
}
