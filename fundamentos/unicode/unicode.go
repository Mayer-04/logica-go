package main

import (
	"fmt"
	"unicode"
)

/*
* Paquete unicode

El paquete `unicode` en Go proporciona utilidades para clasificar y manipular caracteres Unicode
(también conocidos como runes).
Esto incluye funciones para determinar si un rune es una letra, dígito, espacio en blanco, etc.,
así como para convertir entre mayúsculas y minúsculas.

* Algunos rangos Unicode comunes disponibles en el paquete:

- unicode.Latin    → caracteres latinos (Ej. A–Z, a–z, ñ, etc.)
- unicode.Han      → caracteres chinos
- unicode.Cyrillic → caracteres cirílicos

* Subpaquetes relacionados:

- `unicode/utf8`  → Para codificar y decodificar texto en UTF-8.
- `unicode/utf16` → Para codificar y decodificar texto en UTF-16.
*/

func main() {
	// * unicode.IsLetter.
	// Determina si un rune representa una letra en cualquier idioma o alfabeto.
	fmt.Printf("¿'A' es una letra? %t\n", unicode.IsLetter('A')) // true
	fmt.Printf("¿'ñ' es una letra? %t\n", unicode.IsLetter('ñ')) // true
	fmt.Printf("¿'5' es una letra? %t\n", unicode.IsLetter('5')) // false

	// * unicode.IsDigit.
	// Determina si un rune es un dígito decimal (0-9).
	fmt.Printf("¿'5' es un dígito? %t\n", unicode.IsDigit('5')) // true
	fmt.Printf("¿'A' es un dígito? %t\n", unicode.IsDigit('A')) // false

	// * unicode.IsSpace.
	// Determina si un rune es un carácter de espacio en blanco, incluyendo tabulaciones, saltos de línea, etc.
	fmt.Printf("¿' ' es un espacio? %t\n", unicode.IsSpace(' '))    // true
	fmt.Printf("¿'\\t' es un espacio? %t\n", unicode.IsSpace('\t')) // true
	fmt.Printf("¿'\\n' es un espacio? %t\n", unicode.IsSpace('\n')) // true

	// * unicode.IsUpper.
	// Verifica si un rune está en mayúsculas.
	fmt.Printf("¿'A' está en mayúsculas? %t\n", unicode.IsUpper('A')) // true
	fmt.Printf("¿'a' está en mayúsculas? %t\n", unicode.IsUpper('a')) // false
	fmt.Printf("¿'Ñ' está en mayúsculas? %t\n", unicode.IsUpper('Ñ')) // true

	// * unicode.IsLower.
	// Verifica si un rune está en minúsculas.
	fmt.Printf("¿'a' está en minúsculas? %t\n", unicode.IsLower('a')) // true
	fmt.Printf("¿'A' está en minúsculas? %t\n", unicode.IsLower('A')) // false

	// * unicode.IsControl.
	// Determina si un rune es un carácter de control (no imprimible), como tabuladores o saltos de línea.
	fmt.Printf("¿'\\n' es un carácter de control? %t\n", unicode.IsControl('\n')) // true
	fmt.Printf("¿'\\t' es un carácter de control? %t\n", unicode.IsControl('\t')) // true
	fmt.Printf("¿'A' es un carácter de control? %t\n", unicode.IsControl('A'))    // false

	// * unicode.IsPunct.
	// Verifica si un rune es un signo de puntuación (.,;!? etc.).
	fmt.Printf("¿'.' es un signo de puntuación? %t\n", unicode.IsPunct('.')) // true
	fmt.Printf("¿',' es un signo de puntuación? %t\n", unicode.IsPunct(',')) // true
	fmt.Printf("¿'!' es un signo de puntuación? %t\n", unicode.IsPunct('!')) // true

	// * unicode.ToUpper.
	// Convierte un rune a su equivalente en mayúsculas.
	fmt.Printf("'a' en mayúsculas: %c\n", unicode.ToUpper('a')) // 'A'
	fmt.Printf("'ñ' en mayúsculas: %c\n", unicode.ToUpper('ñ')) // 'Ñ'

	// * unicode.ToLower.
	// Convierte un rune a su equivalente en minúsculas.
	fmt.Printf("'A' en minúsculas: %c\n", unicode.ToLower('A')) // 'a'
	fmt.Printf("'Ñ' en minúsculas: %c\n", unicode.ToLower('Ñ')) // 'ñ'

	// * unicode.Is.
	// Sirve para saber si un carácter (rune) pertenece a un grupo de letras específico.
	// Por ejemplo: letras latinas (como la A) o caracteres chinos (como 漢).
	fmt.Printf("¿'A' pertenece al grupo Latin? %t\n", unicode.Is(unicode.Latin, 'A'))     // true
	fmt.Printf("¿'漢' pertenece al grupo Han (chino)? %t\n", unicode.Is(unicode.Han, '漢')) // true

	// * unicode.SimpleFold.
	// Nos ayuda a comparar letras sin importar si están en mayúsculas o minúsculas.
	// Va cambiando entre variantes de una misma letra: por ejemplo 'A' → 'a' → 'A'.
	// Es útil cuando queremos comparar letras ignorando el caso.
	folded := unicode.SimpleFold('A')
	fmt.Printf("¿unicode.SimpleFold('A') devuelve 'a'? %t\n", folded == 'a') // true

	// * unicode.To.
	// Convierte un carácter (rune) al tipo de letra que elijamos:
	// - unicode.UpperCase → convierte a mayúscula
	// - unicode.LowerCase → convierte a minúscula
	// - unicode.TitleCase → convierte a mayúscula de título (inicial en mayúscula)
	converted := unicode.To(unicode.UpperCase, 'ñ')
	fmt.Printf("unicode.To(unicode.UpperCase, 'ñ') = %c\n", converted) // 'Ñ'
}
