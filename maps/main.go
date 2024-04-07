package main

import "fmt"

func main() {

	/*codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	fmt.Println(codes["en"])
	fmt.Println(codes["fr"])
	fmt.Println(codes["hi"])*/

	codes := map[string]int{"en": 1, "fr": 2, "hi": 3}
	value, found := codes["en"]
	fmt.Println(found, value)
	value, found = codes["hh"]
	fmt.Println(found, value)
	valor, encontrado := codes["hi"]
	fmt.Println(encontrado, valor)

	//truncate map
	// codes = make(map[string]string)

}
