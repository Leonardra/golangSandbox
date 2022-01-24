package main

func namedReturnValue(firstNumber int, secondNumber int) (sum int, product int, difference int){
	sum = firstNumber + secondNumber
	difference = firstNumber -secondNumber
	product = firstNumber * secondNumber
	return
}

