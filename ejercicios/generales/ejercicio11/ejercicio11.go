package main

import (
	"fmt"
	"strings"
)

/*
* Ejercicio 2: Palindromo Detector
Descripción: Crea una función que determine si una cadena es un palíndromo
(se lee igual de izquierda a derecha que de derecha a izquierda), ignorando espacios,
signos de puntuación y diferencias entre mayúsculas y minúsculas.

Input: Una cadena de texto.
Output: Un valor booleano (true si es palíndromo, false si no lo es).
*/

func main() {
	// cadena := "  oso.  "
	nombre := "Mayer"
	fmt.Println(detectorPalindrome(nombre))
}

func detectorPalindrome(cadena string) bool {
	// cadena == ""
	if len(cadena) == 0 {
		return false
	}

	var nuevaCadena string
	var originalLimpia string

	for i := len(cadena) - 1; i >= 0; i-- {
		// if string(cadena[i]) != " " && string(cadena[i]) != "." {
		// 	nuevaCadena += string(cadena[i])
		// }
		if string(cadena[i]) == " " || string(cadena[i]) == "." {
			continue
		}

		nuevaCadena += string(cadena[i])
	}

	for i := range len(cadena) {
		// if string(cadena[i]) != " " && string(cadena[i]) != "." {
		// 	originalLimpia += string(cadena[i])
		// }

		if string(cadena[i]) == " " || string(cadena[i]) == "." {
			continue
		}
		nuevaCadena += string(cadena[i])
	}

	nuevaMinuscula := strings.ToLower(nuevaCadena)
	originalMinuscula := strings.ToLower(originalLimpia)

	return nuevaMinuscula == originalMinuscula
}
