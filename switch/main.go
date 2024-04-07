package main

import "fmt"

func main() {

	var i int = 10
	switch i {
	case -5:
		fmt.Println("10")
	case 10:
		fmt.Println("i is 10")
		fallthrough //The fallthrough keyword is used in switch-case to force the execution flow to fall through the succesive case block.
	case 100, 200:
		fmt.Println("i is either 100 or 200")
		fallthrough
	default:
		fmt.Println("i is neither 0, 100 or 200")
	}
}
