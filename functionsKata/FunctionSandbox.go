package main

import "fmt"

func main() {
	sum, product, difference := namedReturnValue(34, 21)
	fmt.Printf("%d\n%d\n%d\n", sum, product, difference)
}
