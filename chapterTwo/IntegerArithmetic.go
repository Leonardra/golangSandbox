package main

import "fmt"

func main() {
	//age := 56
	//var addedAge int = 45
	//newAge := age + addedAge
	//fmt.Println(newAge)
	//fmt.Println("1 + 1 = ", 1+1)

	//loops
	//increaseNumber()
	//fizzBuzz()
	//divisibleByThree()

	numbers := []int{1,2,4,5}

	//functionWithoutPointer(numbers)
	functionWithPointer(&numbers)
	fmt.Println(numbers)
}

