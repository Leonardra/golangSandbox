package main

import "fmt"

func getAValueThroughTheKey(){
	var students = map[string]int{"Oluwatobi":14, "Emmanuel": 34,"Janet":70}

	for key, value := range students{
		fmt.Printf("%s is mapped to %d\n", key, value)
	}
}
