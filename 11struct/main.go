package main

import "fmt"

func main() {
	fmt.Println("Struct in Go")

	student := User{"Aswin", "aswin@ap.com", 13, true}

	fmt.Printf("The type of the student %T\n", student)
	fmt.Println("User : ", student)
	fmt.Printf("The details are %+v\n", student)
	fmt.Printf("The name : %v and Email %v", student.Name, student.Email)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
