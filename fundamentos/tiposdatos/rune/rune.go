package main

import "fmt"

/*
* Rune: Runa en Go
- En Go, `rune` es un alias para el tipo `int32`. Esto significa que puede almacenar puntos de código Unicode.
- Una `runa` representa un carácter Unicode y puede ocupar más de un byte en la codificación UTF-8.
- El valor de una `runa` es el punto de código Unicode del carácter.
- Es especialmente útil cuando trabajas con texto que incluye caracteres especiales o no latinos.
- Al usar `rune`, Go permite manejar texto Unicode de manera eficiente.

* ASCII vs Unicode:
- Los caracteres ASCII, como las letras 'a', 'n', 'o', 's', y el espacio " ", se representan con un solo byte (0-127).
- Los caracteres fuera de ASCII, como 'ñ' o emojis como 👴🏼, pueden requerir varios bytes.
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

	// Definiendo una variable de tipo rune para la letra 'a'
	// El punto de código Unicode para 'a' es U+0061
	runa := 'a'
	fmt.Printf("Runa 'a': %d\n", runa) // Output: Runa 'a': 97

	// Ejemplo con un emoji 🍎 (Unicode: U+1F34E)
	runa2 := '🍎'
	fmt.Printf("Runa para 🍎: %d\n", runa2) // Output: Runa para 🍎: 127822

	// Cadena con caracteres Unicode
	s := "años, 👴🏼"

	// Convertimos la cadena en una slice de runas (puntos de código Unicode)
	runaSlice := []rune(s)
	fmt.Println("Runa slice:", runaSlice) // Output: [97 241 111 115 44 32 128116 127996]

	/*
	   * Desglose de la slice de runas:
	   97   → 'a'
	   241  → 'ñ'
	   111  → 'o'
	   115  → 's'
	   44   → ','
	   32   → (espacio)
	   128116 → 👴 (anciano)
	   127996 → tono de piel claro (modificador de emoji)
	*/

	// Recorriendo la cadena y mostrando la posición en bytes y las runas
	for i, r := range s {
		fmt.Printf("Posición en bytes %d, Runa: %c\n", i, r)
	}
	/*
	 La posición en el loop es en bytes, no en caracteres. UTF-8 usa múltiples bytes para algunos caracteres,
	 por lo que la posición en bytes de cada runa no siempre coincide con la posición en caracteres.
	*/
}
