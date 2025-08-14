package main

import "fmt"

/*
* Conteo de caracteres únicos
Dado un string, utiliza un mapa para contar cuántas veces aparece cada carácter en la cadena.
Luego imprime cuántos caracteres únicos (que aparecen una sola vez) tiene el string.

Ejemplo:

Entrada: "manzana"
Salida: 2 (m, n) son únicos
*/

func main() {
	fruta := "manzana"
	fmt.Println(contarCaracteresUnicos(fruta))
}

func contarCaracteresUnicos(s string) int {
	nuevoMapa := make(map[rune]int)
	for _, caracter := range s {
		nuevoMapa[caracter]++
	}

	caracteresUnicos := 0
	for _, contador := range nuevoMapa {
		if contador == 1 {
			caracteresUnicos++
		}
	}

	return caracteresUnicos
}
