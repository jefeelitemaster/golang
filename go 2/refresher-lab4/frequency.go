/*In this question, we are building a tool that analyzes the frequency of words in a given text.
You need to implement a function wordFrequency that receives a string and returns a map with
the frequency of each word in the string.


A Go file is located at /root/code/frequency directory.*/

package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	// TODO: implement this function
	//conversión a minúsculas para un análisis insensible a mayúsculas y minúsculas
	//text = strings.ToLower(text)

	// dividir el texto en palabras
	palabras := strings.Fields(text)

	// crear un mapa para almacenar las frecuencias de las palabras
	cuentaPalabras := make(map[string]int)
	//print(cuentaPalabras)
	// Iterar sobre cada palabra y actualizar su recuento en el mapa
	for _, palabra := range palabras {
		fmt.Println(cuentaPalabras[palabra])
		cuentaPalabras[palabra] = cuentaPalabras[palabra] + 1 //Incrementamos el valor asociado a esa clave (que representa el recuento de la palabra) en 1.
		fmt.Println(cuentaPalabras)
	}
	return cuentaPalabras
}

func main() {
	text := "The quick brown fox jumps over the lazy dog dog dog dog The The fox Fox quick Allan Lara Fuentes"
	fmt.Println(wordFrequency(text))

}

/*Expected output:



map[The:1 brown:1 dog:1 fox:1 jumps:1 lazy:1 over:1 quick:1 the:1]
*/
