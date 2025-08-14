package main

import (
	"fmt"
	"strings"
)

/*
* Cuenta vocales:
Escribe un programa que reciba una cadena de texto y cuente cuántas vocales (a, e, i, o, u) contiene.

- El conteo debe ser insensible a mayúsculas y minúsculas.

* Ejemplo:
Entrada: "Hola Mundo"
Salida: 4
*/

func main() {
	str := "banano"
	result := contarVocales(str)
	fmt.Println(result)

	// Ejemplo 2
	result2 := countVowels(str)
	fmt.Println(result2)

	// Ejemplo 3
	fmt.Println(GetCount(str))
}

// En terminos de eficiencia este código no es el mejor.
func contarVocales(s string) int {
	vocales := [5]string{"a", "e", "i", "o", "u"}
	minuscula := strings.ToLower(s)
	contador := 0

	for _, v := range vocales {
		for _, c := range minuscula {
			char := string(c)
			if char == v {
				contador++
			}
		}
	}

	return contador
}

// Ejemplo 2
func countVowels(s string) int {
	// Mapa que contiene las vocales como claves y un valor booleano true.
	// Esto nos permite buscar si un carácter es una vocal de manera eficiente.
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0

	// Convertir la cadena a minúsculas y recorrer cada carácter
	for _, c := range strings.ToLower(s) {
		// Verificamos si el carácter 'c' es una clave en el mapa 'vocales'
		// Si 'c' es una vocal (es decir, está en el mapa), el valor es true y entramos en el if
		if vowels[c] {
			count++
		}
	}

	return count
}

// Ejemplo 3
func GetCount(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
