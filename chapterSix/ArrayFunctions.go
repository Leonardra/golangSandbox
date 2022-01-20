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

func appendSlices(){
	firstSlice :=[]int {1,2,3,4}
	secondSlice := []int {5,6,7,8,9}
	thirdSlice := append(firstSlice, secondSlice...)
	fmt.Println(thirdSlice)
}

func findSliceLength(){
	ints := make([]int, 3, 9)
	fmt.Println(ints)
}

func  smallestNumber(){
	numberSeries := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	var smallest = numberSeries[0]
	for _, series := range numberSeries {
		if series < smallest{
			smallest = series
		}
	}
	fmt.Println(smallest)
}


