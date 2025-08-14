package main

import (
	"fmt"
)

/*
* Longitud de la palabra más larga

Crea una función que encuentre la palabra más larga en una cadena de texto y
devuelva su longitud como número entero.

Ejemplo: En "El zorro salta sobre el perro", la palabra más larga
es "sobre" (5 caracteres), por lo que la función debe retornar 5.
*/

func main() {
	word := "El zorro salta sobre el perro"
	result := wordLength(word)
	fmt.Println(result)
}

func wordLength(s string) int {

	palabras := make([]string, 0)

	for _, word := range s {
		runa := string(word)
		palabras = append(palabras, runa)
	}

	fmt.Println("palabras", palabras)

	for _, pala := range palabras {
		fmt.Println("palabra", pala)
	}

	// palabras := strings.Split(s, " ")
	// fmt.Println(palabras)

	return 0
}
