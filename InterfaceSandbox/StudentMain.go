package main

import (
	"fmt"
)

func main() {
	var firstName string
	var lastName string
	var age uint32
	fmt.Println("Enter first name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter last name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter age name:")
	fmt.Scan(&age)


	studentService := createService()
	student := new(Student)
	student.id = 1
	student.firstName = firstName
	student.lastName = lastName
	student.age = age

	fmt.Println(student.firstName)

	studentService.addStudent(student)

}
