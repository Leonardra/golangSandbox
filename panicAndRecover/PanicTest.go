package main

import (
	"fmt"
)

func main() {
	var guessedNumber int
	fmt.Println("Enter number:")
	fmt.Scan(&guessedNumber)

	switch guessedNumber {
	case 1:
		fmt.Println("In range")
	case 2:
		fmt.Println("In range")
		default:
			panic("Not in range")
	}

}
