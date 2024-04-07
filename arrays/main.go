package main

import "fmt"

func main() {

	//var grades [3]int = [3]int{10, 20, 30}

	//grades := [3]int{10, 20, 30}
	//grades := [...]int{10, 20, 30}
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}

	fmt.Println(arr[2][1])
}
