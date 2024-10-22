package main

import "fmt"

/*
Crea una función que reciba una string y devuelva un map donde la clave es cada letra única
y el valor es el número de veces que esa letra aparece en el string.
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

	for i := 0; i < len(s); i++ {
		caracter := string(s[i])
		newMap[caracter]++
	}

	return newMap
}
