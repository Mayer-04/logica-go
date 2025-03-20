package main

import "fmt"

/*
* Reemplazar caracteres:
Escribe un programa que reciba una cadena de texto y reemplace todos los espacios en blanco por guiones bajos (_).

Ejemplo:
Entrada: "Go is fun"
Salida: "Go_is_fun"
*/

func main() {
	str := "Go is fun"
	result := reemplazarCaracteres(str)
	fmt.Println(result)

	//* Ejemplo 2
	result2 := replaceCharacters(str)
	fmt.Println(result2)
}

func reemplazarCaracteres(s string) string {
	str := ""
	for _, rune := range s {
		char := string(rune)
		if char == " " {
			char = "_"
		}
		str += char
	}

	return str
}

func replaceCharacters(s string) string {
	result := ""
	for _, c := range s {
		if c == ' ' {
			c = '_'
		}
		result += string(c)
	}

	return result
}
