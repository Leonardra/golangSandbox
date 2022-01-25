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

		boy := new(Boy)
		boy.firstName = "Tosin"
		boy.lastName = "Saba"
		boy.accountNumber = "12345"
		fmt.Println(boy)

}

func assignTask(human *Human){
	human.task = "Deputy"
}