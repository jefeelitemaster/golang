package main

import "fmt"

func main() {

	var a, b int = 10, 20
	switch {
	case a+b == 30:
		fmt.Println("equal 30")
	case a+b <= 30:
		fmt.Println("less than or equal to 30")
	default:
		fmt.Println("greater than 30")
	}
}

/* This is because Go uses an implicit break statement for each case
This is very different from languages like C or Java, where we have to explicitly write the break keyword*/
