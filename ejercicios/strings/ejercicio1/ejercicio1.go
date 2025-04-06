package main

import "fmt"

/*
Escribe una función que reciba una cadena de texto y retorne un mapa en el que:
- Cada clave representa un carácter único de la cadena.
- Cada valor indica la cantidad de veces que ese carácter aparece en la cadena.

* Ejemplo:
Entrada: "golang"
Salida: {'g': 2, 'o': 1, 'l': 1, 'a': 1, 'n': 1}
*/

func main() {
	abc := "abcabc"
	fruta1 := "coco"
	fruta2 := "banana"

	fmt.Printf("abs: %#v\n", contarCaracteres(abc))     // abs: map[string]int{"a":2, "b":2, "c":2}
	fmt.Printf("coco: %#v\n", contarCaracteres(fruta1)) // coco: map[string]int{"c":2, "o":2}
	fmt.Printf("banana: %#v", contarCaracteres(fruta2)) // banana: map[string]int{"a":3, "b":1, "n":2}
}

func contarCaracteres(s string) map[string]int {
	newMap := make(map[string]int)

	for i := range len(s) {
		caracter := string(s[i])
		newMap[caracter]++
	}

	return newMap
}
