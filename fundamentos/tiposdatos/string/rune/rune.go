package main

import (
	"fmt"
	"unicode/utf8"
)

/*
* Rune: Runa en Go
- En Go, `rune` es un alias para el tipo `int32`. Esto significa que puede almacenar puntos de código Unicode.
- Una `runa` representa un carácter Unicode y puede ocupar más de un byte en la codificación UTF-8.
- El valor de una `runa` es el punto de código Unicode del carácter.
- Es especialmente útil cuando trabajas con texto que incluye caracteres especiales o no latinos.
- Al usar `rune`, Go permite manejar texto Unicode de manera eficiente, ya que `for range` recorre caracteres
y no bytes.

* ASCII vs Unicode:
- Los caracteres ASCII, como las letras 'a', , 'o', 's', y el espacio " ",
se representan con un solo byte (0-127).
- Los caracteres fuera del rango ASCII, como 'ñ' o emojis como 👴🏼,
pueden requerir de 2 a 4 bytes debido a la codificación UTF-8 de longitud variable.

- len(string) cuenta bytes.
- utf8.RuneCountInString(string) cuenta runes (caracteres Unicode).
*/

func main() {
	/*
	   * Conversión de valores Hexadecimales a Decimales (ejemplo con la letra 'a'):
	   - El valor Unicode para 'a' es U+0061, que en hexadecimal es 0x61.
	   - El sistema hexadecimal (base 16) usa los dígitos 0-9 y las letras A-F (donde A = 10, B = 11, ..., F = 15).
	   - Para convertir 0x61 a decimal, se descompone de la siguiente manera:
	       1. 6 * 16^1 = 6 * 16 = 96
	       2. 1 * 16^0 = 1 * 1 = 1
	       3. 96 + 1 = 97
	   - Por lo tanto, el valor decimal de U+0061 es 97.
	*/

	// Definiendo una variable de tipo rune para la letra 'a'.
	// El punto de código Unicode para 'a' es U+0061.
	runa := 'a'
	fmt.Printf("Runa 'a': %d\n", runa) // Output: Runa 'a': 97

	// Ejemplo con un emoji 🍎 (Unicode: U+1F34E).
	runa2 := '🍎'
	fmt.Printf("Runa para 🍎: %d\n", runa2) // Output: Runa para 🍎: 127822

	// Cadena con caracteres Unicode (mezcla de ASCII, acentuados y emoji con modificador).
	s := "años, 👴🏼"

	// Convertimos la cadena en una slice de runas (Unicode code points).
	runaSlice := []rune(s)
	fmt.Println("Runa slice:", runaSlice) // Output: [97 241 111 115 44 32 128116 127996]

	/*
	   * Desglose del slice de runas:
	   97      → 'a'
	   241     → 'ñ'
	   111     → 'o'
	   115     → 's'
	   44      → ','
	   32      → espacio
	   128116  → 👴 (anciano)
	   127996  → modificador de tono de piel claro
	   Nota: 👴🏼 es una combinación de dos runas, no una sola.
	*/

	// Recorriendo la cadena y mostrando la posición en bytes y las runas.
	/*
	 El índice `i` representa la posición inicial en bytes de la runa dentro de la cadena UTF-8.
	 Esto es importante porque caracteres como 'ñ' o emojis ocupan múltiples bytes.
	*/
	for i, r := range s {
		fmt.Printf("Posición en bytes %d, Runa: %c\n", i, r)
	}

	const name = "Mayer"

	// Cada carácter ASCII ocupa exactamente 1 byte en UTF-8.
	// Como "Mayer" tiene 5 caracteres, el resultado será 5 bytes.
	fmt.Printf("Número de bytes de 'name' ('Mayer'): %d\n", len(name)) // Output: 5

	// Como todos los caracteres de "Mayer" son ASCII, la cantidad de runas (caracteres) es la misma.
	fmt.Printf("Número de runas en 'name': %d\n", utf8.RuneCountInString(name)) // Output: 5

	// Ahora usamos una cadena que contiene un carácter especial: 'é'.
	const secondName = "Andrés"

	// En UTF-8, la letra 'é' ocupa 2 bytes, mientras que las demás letras ocupan 1 byte cada una.
	// Como "Andrés" tiene 6 caracteres, pero la 'é' usa 2 bytes, el total de bytes será 7.
	fmt.Printf("Número de bytes de 'secondName' ('Andrés'): %d\n", len(secondName)) // Output: 7

	// Aunque ocupa 7 bytes, sigue teniendo 6 caracteres visibles (runas).
	fmt.Printf("Número de runas en 'secondName': %d\n", utf8.RuneCountInString(secondName)) // Output: 6

	fmt.Printf("Character: %s, Bytes: %d\n", "A", len("A")) // Output: 1 byte
	fmt.Printf("Character: %s, Bytes: %d\n", "ñ", len("ñ")) // Output: 2 bytes
	fmt.Printf("Character: %s, Bytes: %d\n", "字", len("字")) // Output: 3 bytes
	fmt.Printf("Character: %s, Bytes: %d\n", "😎", len("😎")) // Output: 4 bytes
}
