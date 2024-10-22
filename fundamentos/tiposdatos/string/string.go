package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
* Strings: Secuencia de bytes.
- Los strings son secuencias inmutables de bytes, codificados en UTF-8.
Esto significa que cada string es una colección de bytes que representan texto.
- Un carácter ASCII ocupa un byte, mientras que los caracteres Unicode pueden necesitar más (hasta 4 bytes en UTF-8).
- Los strings son inmutables: una vez creados, no pueden modificarse directamente.
- Los strings se declaran entre comillas dobles "".
- Los espacios en blanco, saltos de línea y otros caracteres especiales también son parte del string.
- Se pueden concatenar usando el operador `+` o `fmt.Sprintf()`, que formatea cadenas.
- Los caracteres entre comillas simples '' se tratan como runas (`rune`), un alias para el tipo int32.

* Substrings: Son secuencias de caracteres consecutivos que se encuentran dentro del string original.
* Substrings únicos: Son substrings que no se repiten.

* Algunas definiciones:
- Secuencia: Colección ordenada de elementos, uno tras otro.
- Byte: Es la unidad básica de almacenamiento de datos en una computadora. Cada byte contiene 8 bits (donde un bit puede ser 0 o 1).
Un byte puede representar un número entre 0 y 255.
- Codificar: Es el proceso de transformar datos o información a un formato que pueda ser procesado o entendido por una máquina o computadora.
- ASCII: Conjunto de caracteres que usa 7 bits para representar 128 caracteres (0-127), incluyendo letras, números y signos de puntuación.
- Unicode: Estándar que abarca caracteres de casi todos los idiomas y símbolos, usando hasta 4 bytes para representar caracteres complejos.
Asigna un número único (código) a cada carácter en casi todos los idiomas del mundo.
- UTF-8: Es un tipo de codificación que se usa para representar caracteres de texto en una computadora. Incluye
símbolos de otros idiomas, emojis y caracteres especiales. Puede usar entre 1 y 4 bytes para representar cada carácter.
*/

func main() {
	// Declaración de un string
	name := "Mayer"
	fmt.Printf("tipo: %T - valor: %s\n", name, name)

	// Uso de salto de línea con `\n`
	secondName := "Andres"
	fmt.Printf("tipo: %T - valor: %v\n", secondName, secondName)

	// Interpolación de strings con `fmt.Sprintf()`. Permite insertar valores dentro de una cadena.
	fullName := fmt.Sprintf("Mi nombre es %s %s", name, secondName)
	fmt.Printf("interpolación: %s\n", fullName)

	// Modificar un string (convertir a slice de bytes).
	// Los strings en Go son inmutables, pero se pueden convertir a un slice de bytes para hacer modificaciones.
	cadena := "Hola"
	bytes := []byte(cadena)                      // Convierte el string en un slice de bytes
	bytes[0] = 'M'                               // Modifica el primer carácter
	nuevoString := string(bytes)                 // Convierte de nuevo a string
	fmt.Printf("nuevoString: %q\n", nuevoString) // "Mola"

	// Obtener la longitud de un string usando la función `len()`.
	fmt.Println("longitud de name:", len(name))

	// Concatenar dos strings utilizando el operador +.
	nameAndSecondName := name + " " + secondName
	fmt.Println("concatenar strings:", nameAndSecondName)

	//* Paquete "strings": Métodos útiles para manipular cadenas.

	// Verifica si una cadena contiene una subcadena. Retorna true o false.
	str := strings.Contains(name, "Mayer")
	fmt.Println("contains:", str)

	// Obtener el índice de la primera aparición de una subcadena. Retorna el índice como entero.
	str2 := strings.Index(name, "y")
	fmt.Println("index:", str2)

	// Convierte todos los caracteres de una cadena a minúsculas.
	str3 := strings.ToLower(name)
	fmt.Println("toLower:", str3)

	// Función `Replace()` Reemplaza una subcadena por otra dentro de la cadena original.
	// s: Es la cadena original.
	// old: Es la cadena que quieres reemplazar.
	// new: La nueva subcadena que va a reemplazar a `old`.
	// n: indica cuántas veces realizar el reemplazo; si es -1, se reemplazan todas las apariciones.
	fruta := "coco"
	str4 := strings.Replace(fruta, "o", "a", -1)
	fmt.Printf("replace: %q\n", str4) // Output: "caca"

	// Divide una cadena en partes usando un separador. Retorna un slice de strings.
	str5 := strings.Split(name, "")
	fmt.Printf("split: %q\n", str5) // Output: ["M" "a" "y" "e" "r"]

	// Verifica si una cadena termina con una subcadena. Retorna true o false.
	str6 := strings.HasSuffix(name, "er")
	fmt.Println("hasSuffix:", str6) // Output: true

	// Verifica si una cadena comienza con una subcadena. Retorna true o false.
	str7 := strings.HasPrefix(name, "ma")
	fmt.Println("hasPrefix:", str7) // Output: false

	// Elimina los espacios en blanco del inicio y final de una cadena.
	str8 := strings.TrimSpace(name)
	fmt.Println("trimSpace:", str8)

	// Elimina caracteres específicos del inicio y final de una cadena.
	lastname := "!!Chaves¡¡"
	str9 := strings.Trim(lastname, "!¡")
	fmt.Println("trim:", str9)

	// Une los elementos de un slice en una sola cadena, usando un separador.
	sli := []string{"foo", "bar", "baz"}
	fmt.Printf("join: %q\n", strings.Join(sli, ", ")) // Output: "foo, bar, baz"

	// Función `Clone` (Introducida en Go 1.18). Devuelve una copia de la cadena.
	// Útil para reducir el uso de memoria en algunos casos.
	animal := "León"
	clone := strings.Clone(animal)
	fmt.Println("clone:", clone)

	// Función `Cut` (Nueva en Go 1.18). Divide una cadena en dos partes usando un separador.
	// Retorna la parte antes del separador, la parte después y un booleano que indica si se encontró el separador.
	s := "Hola mundo"
	sep := " "
	before, after, found := strings.Cut(s, sep)
	fmt.Println("Antes del separador:", before)   // Output: "Hola"
	fmt.Println("Después del separador:", after)  // Output: "mundo"
	fmt.Println("¿Separador encontrado?:", found) // Output: true

	// Función `CutPrefix` (Nueva en Go 1.20). Elimina un prefijo específico del inicio de una cadena.
	// Prefijo: Una cadena que aparece al principio de otra.
	// Retorna la cadena sin el prefijo y un booleano que indica si el prefijo fue encontrado.
	gopher := "gopher"
	prefix := "go"
	after, ok := strings.CutPrefix(gopher, prefix)
	fmt.Println("Después del prefijo:", after) // Output: "opher"
	fmt.Println("¿Prefijo encontrado?:", ok)   // Output: true

	// Función `CutSuffix` (Nueva en Go 1.20). Elimina un sufijo específico del final de una cadena.
	// Sufijo: Una cadena que aparece al final de otra.
	// Retorna la cadena sin el sufijo y un booleano que indica si el sufijo fue encontrado.
	document := "documento.txt"
	suffix := ".txt"
	newS, ok := strings.CutSuffix(document, suffix)
	fmt.Println("Después del sufijo:", newS) // Output: "documento"
	fmt.Println("¿Sufijo encontrado?:", ok)  // Output: true

	//* Paquete "strconv": Conversión entre strings y tipos numéricos.

	// Convierte un string que representa un número entero a un tipo int.
	age := "23"
	age2, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tipo: %T - valor: %v\n", age2, age2) // Output: tipo: int - valor: 23

	// Convierte un string que representa un número flotante a tipo float64.
	weight := "70.5"
	weight2, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tipo: %T - valor: %v", weight2, weight2) // Output: tipo: float64 - valor: 70.5
}
