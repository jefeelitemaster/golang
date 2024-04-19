package main

import "fmt"

type Producto struct {
	nombre   string
	cantidad int
	precio   float64
}

func imprimeNombre(p Producto) {
	fmt.Println(p.nombre)
}

func imprimeCantidad(p Producto) {
	fmt.Println(p.cantidad)
}

func imprimePrecio(p Producto) {
	fmt.Println(p.precio)
}

func main() {
	var p Producto
	p.nombre = "Chair"
	p.cantidad = 5
	p.precio = 700

	fmt.Println("Detalles del producto:")
	imprimeNombre(p)
	defer imprimePrecio(p) //defer es para que espere a imprimirse hasta que se termine la siguiente función
	imprimeCantidad(p)
}
