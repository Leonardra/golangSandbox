package main

import "fmt"

func increaseNumber(){
	counter := 1
	for counter < 10 {
		if counter == 7 {
			fmt.Println("This is", counter)
		}
		fmt.Println(counter)
		counter++
	}
}

func fizzBuzz(){
	for counter := 0; counter < 100; counter++{
		if counter % 3 == 0 {
			fmt.Println("Fizz")
		}else if counter % 5 == 0 {
			fmt.Println("Buzz")
		}else if (counter % 3 == 0) && (counter % 5 == 0){
			fmt.Println("FizzBuzz")
		}else{
			fmt.Println(counter)
		}
	}
}

func divisibleByThree(){
	for i := 0; i < 100; i++ {
		if i % 3 == 0{
			fmt.Println(i)
		}
	}
}

