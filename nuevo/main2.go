package main

import "fmt"

func main() {

	title := "Sir"
	fmt.Println(title)
	var i int
	fmt.Printf("%d\n", i)
	var fl float64
	fmt.Printf("%.2f\n", fl)
	var num1, num2, num3 int = 1, 2, 3
	fmt.Println(num1, num2, num3)
	var name1 string = "Allan"
	fmt.Println("\nMi name1 es: ", name1)

	var name2 string
	fmt.Print("Enter your name2: ")
	fmt.Scanf("%s", &name2)
	fmt.Println("Hey there, ", name2)

	var nombre3 string
	var is_muggle bool

	fmt.Print("Enter your nombre & are you a muggle: ")
	fmt.Scanf("%s %t", &nombre3, &is_muggle)
	fmt.Println(nombre3, is_muggle)

}
