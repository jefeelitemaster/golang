package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	//the datatype will be internally determined by the compiler

	var pointer_name = &i
	fmt.Printf("%T, %v \n", pointer_name, pointer_name)

	//we have to use the shorthand declaration operator
	var mi_cadena = "Hola!!"
	apuntador_a_mi_cadena := &mi_cadena
	fmt.Printf("%T %v \n", apuntador_a_mi_cadena, apuntador_a_mi_cadena)

	var b *string = &mi_cadena
	fmt.Println(b)
	fmt.Println(*b)

	*b = "Hola como estás!!"
	fmt.Println(*b)
	fmt.Println(mi_cadena)
}
