package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

		fmt.Println(generateRandomString())
}

func generateRandomString() string{
	var bytesLetter = [5]byte{}
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < 5; i++{
		randomNumber := rand.Intn(122 - 65) + 65
		bytesLetter[i] = byte(randomNumber)
	}
	randomString := string(bytesLetter[:])
	return randomString
}

func assignTask(human *Human){
	human.task = "Deputy"
}