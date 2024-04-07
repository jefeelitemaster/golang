package main

import (
	"fmt"
	"reflect"
)

func main() {

	var grades int = 42
	var message string = "hello world"

	fmt.Printf("variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable message='%v' is of type %v \n", message, reflect.TypeOf(message))

	var isCheck bool = true
	var amount float32 = 5466.54

	fmt.Printf("variable isCheck = '%v' is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = '%v' is of type %T \n", amount, amount)
}
