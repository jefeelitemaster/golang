/* The process of converting one data type to another is known as Type Casting*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	var fl float64 = 45.89
	var inte int = int(fl)
	fmt.Printf("%v\n", inte)

	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T", err, err)

}
