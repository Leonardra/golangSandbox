package main

import "fmt"

func sumElementsInArray(){
	var anArray [5]int

	anArray[0]= 65
	anArray[1]= 78
	anArray[2]= 93
	anArray[3]= 100
	anArray[4]= 67

	total := 0
	for i := 0; i < len(anArray); i++ {
		total+=anArray[i]
	}
	fmt.Println(total)
}

func sumElementsInArrayWithForLoop(){
	var anArray [5]int

	anArray[0]= 65
	anArray[1]= 78
	anArray[2]= 93
	anArray[3]= 100
	anArray[4]= 67

	total := 0

	for _, value := range anArray {
		total+=value
	}
	fmt.Println(total)
}

func sumOfArrayLiterals(){
	anArray := [5]int{65, 78, 93, 100, 67}
	total := 0
	for _, value := range anArray{
		total+=value
	}
	fmt.Println(total)
}
