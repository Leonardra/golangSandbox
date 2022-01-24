package main

import "fmt"

func main() {
		firstBoy := new(Human)
		firstBoy.firstName = "Oluwatobi"
		firstBoy.lastName = "Jolayemi"
		firstBoy.age = 40
		assignTask(firstBoy)
		firstBoy.setFirstname("Oluwatosin")
		fmt.Println(firstBoy)
}

func assignTask(human *Human){
	human.task = "Deputy"
}