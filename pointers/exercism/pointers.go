//https://exercism.org/tracks/go/concepts/pointers

package main

import "fmt"

func incrementarEdadPedro(m map[string]int) {
	m["Pedro"] += 1
}

func main() {

	var a int

	a = 3

	fmt.Println(a)

	var p *int //'p' contains the memory address of an integer

	fmt.Println(p)

	var b int
	b = 2

	var o *int
	o = &b
	//o = &b //the variable 'o' contains the memory address of 'b'

	fmt.Println("Valor: ", b, "\nDirección de memoria", o)

	var c int
	c = *o //dereferencing

	fmt.Println(c)

	*o = *o + 5 + *o

	fmt.Println(b)

	//var q *int      // q is nil initially
	//fmt.Println(*q) // panic: runtime error: invalid memory address or nil pointer dereference

	type Persona struct {
		Nombre string
		Edad   int
	}

	var pedro Persona
	pedro = Persona{Nombre: "Pedro", Edad: 22}

	var pointerPersona *Persona
	pointerPersona = &pedro

	//When we have a pointer to a struct, we don't need to dereference the pointer before accessing one of the fields:
	fmt.Println(pointerPersona.Nombre)

	//slices and maps are already pointers,  Imagine we have a function that increments the value of a key in a map:
	ages := map[string]int{
		"Pedro": 25,
	}

	incrementarEdadPedro(ages)
	fmt.Println(ages)
}
