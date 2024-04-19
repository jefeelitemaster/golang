package main

import "fmt"

func returnCube(n int) int {

	return n * n * n
}

func modify(s *string) { //passing by reference in functions
	*s = "world"
}

func modifyslice(s []int) { //slices are passed by reference, by default
	s[0] = 100
}

func modifyMap(m map[string]int) { //maps, as well, are passed by reference, by default
	m["K"] = 75
}

func main() {

	fmt.Println("\n Slice")
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modifyslice(slice)
	fmt.Println(slice)
	fmt.Println("\n Map")
	ascii_codes := make(map[string]int)
	ascii_codes["A"] = 65
	ascii_codes["F"] = 70
	fmt.Println(ascii_codes)
	modifyMap(ascii_codes)
	fmt.Println(ascii_codes)
	a := "hello"
	fmt.Println(a)
	modify(&a) //dereference operator
	fmt.Println(a)
	//fmt.Println("Functions")
}
