package main

func functionWithoutPointer(numbers []int){
	numbers = append(numbers, 56)
}

func functionWithPointer(numbers *[]int){
	*numbers = append(*numbers, 56)
}
