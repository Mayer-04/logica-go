package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
* Strings en Go: Secuencia de bytes.
- Los strings en Go son inmutables y codificados en UTF-8.
- Para definir un string, utilizamos comillas dobles ("").
- Caracteres especiales (saltos de línea, tabulaciones y espacios) forman parte del string.
- Se almacenan como una secuencia de bytes, pero no son equivalentes a `[]byte`.
- Un carácter ASCII ocupa 1 byte, mientras que los caracteres Unicode pueden ocupar entre 1 y 4 bytes en UTF-8.
- Para concatenar, se usa `+`, `fmt.Sprintf()` o `strings.Builder` (más eficiente para múltiples concatenaciones).
- Los caracteres entre comillas simples ('') son `rune` (alias de `int32`), representando un solo carácter Unicode.
- Cuando necesitas trabajar con caracteres individuales de una cadena, debes usar el tipo `rune`.
- Al iterar sobre una cadena usando el bucle `for range`,
Go automáticamente descompone la cadena en runas (no en bytes).

* Inmutabilidad:
Significa que una vez que se crea, no se puede modificar su contenido.

* Substrings:
- Un substring es una secuencia consecutiva de caracteres dentro de un string.
- Los substrings son inmutables: cualquier modificación genera un nuevo string.

* Substrings únicos:
Son substrings que no se repiten.

* DEFINICIONES CLAVE 📖:

* Secuencia:
Conjunto ordenado de elementos.

* Bit:
Es la unidad más pequeña de información que usa una computadora.
- Un bit solo puede tener dos valores: 0 o 1, como un interruptor que está apagado (0) o encendido (1).
- Los bits se combinan para formar datos más complejos como letras, imágenes o sonidos.
- Cuando se agrupan 8 bits, forman un `byte`, que ya puede representar información útil como una letra o un color.

* Byte:
Es un grupo de 8 bits.
- Al juntar 8 valores de 0 o 1, se pueden crear 256 combinaciones diferentes (de 0 a 255).
Un byte se usa para guardar información como una letra, un número pequeño, o una parte de un color en una imagen.
Por ejemplo, la letra "A" ocupa 1 byte.


* Codificación:
Es el proceso de `transformar` datos o información a un formato específico que pueda ser procesado
o entendido por una máquina o computadora (como UTF-8).

* ASCII:
Conjunto de caracteres que usa `7 bits` para representar 128 caracteres (0-127),
incluyendo letras del alfabeto inglés, números y signos de puntuación.
- Es un subconjunto de UTF-8.

* Unicode:
Estándar internacional que asigna un número único (punto de código) a cada carácter en casi todos los idiomas
del mundo, símbolos matemáticos, emojis y otros símbolos.

* Punto de código:
Valor numérico que representa un carácter específico en Unicode.
Por ejemplo, la letra "A" es el punto de código U+0041.

* UTF-8:
Es un tipo de codificación que se usa para representar caracteres de texto (Unicode) en una computadora.
Puede usar entre 1 y 4 bytes por carácter: 1 byte para caracteres ASCII y más bytes para otros caracteres.
- UTF-8 simplemente extiende ASCII para incluir caracteres adicionales.
- En UTF-8, los códigos de 0 a 127 son exactamente los mismos que en ASCII.

* UTF-16:
Es una codificación de Unicode que usa 2 bytes por carácter.
- Usa mínimo 2 bytes para cada carácter, incluso para caracteres ASCII simples.
- Es más eficiente con idiomas que utilizan muchos caracteres como (japonés, coreano o chino).

*/

func main() {
	// Declaración de un string.
	name := "Mayer"
	fmt.Printf("tipo: %T - valor: %s\n", name, name) // Output: tipo: string - valor: "Mayer"

	// Uso de salto de línea con `\n`.
	secondName := "Andres"
	fmt.Printf("tipo: %T - valor: %v\n", secondName, secondName)

	//* Interpolación de strings con `fmt.Sprintf()`.
	// Permite insertar (incrustar) valores dentro de una cadena como variables o expresiones.
	fullName := fmt.Sprintf("Mi nombre es %s %s", name, secondName)
	fmt.Printf("interpolación: %s\n", fullName) // Output: "Mi nombre es Mayer Andres"

	// Modificar un string (convertir a slice de bytes).
	// Los strings en Go son inmutables, pero se pueden convertir a un slice de bytes para hacer modificaciones.
	// NO se recomienda hacer este proceso, ya que no es eficiente y puede afectar el rendimiento.
	cadena := "Hola"
	bytes := []byte(cadena)                      // Convierte el string en un slice de bytes
	bytes[0] = 'M'                               // Modifica el primer carácter
	nuevoString := string(bytes)                 // Convierte de nuevo a string
	fmt.Printf("nuevoString: %q\n", nuevoString) // "Mola"

	// La función `len()` devuelve el número de bytes, no el número de caracteres reales.
	fmt.Println("longitud de name:", len(name)) // Output: 5
	nuevaCadena := "Mayer, 世界"
	fmt.Println("longitud de nuevaCadena:", len(nuevaCadena)) // Output: 13
	fmt.Println(utf8.RuneCountInString(nuevaCadena))          // Output: 9 caracteres

	// Concatenar dos strings utilizando el operador +.
	nameAndSecondName := name + " " + secondName
	fmt.Println("concatenar strings:", nameAndSecondName) // Output: "Mayer Andres"

	//* Iterar sobre un string con `range`.
	// Go no itera sobre los bytes de un string, sino sobre sus runas.
	for i, char := range name {
		fmt.Printf("index: %d, char: %d\n", i, char) // Output: index: 0, char: 77, etc ...
	}

	//* Paquete "strings": Métodos útiles para manipular cadenas.

	// Verifica si una cadena contiene una subcadena.
	// Devuelve true o false.
	str := strings.Contains(name, "Mayer")
	fmt.Println("contains:", str) // Output: true

	// Obtener el índice de la primera aparición de una subcadena.
	// Devuelve el índice y -1 si no esta presente.
	str2 := strings.Index(name, "y")
	fmt.Println("index:", str2) // Output: 2

	// Convierte todos los caracteres de una cadena a minúsculas.
	str3 := strings.ToLower(name)
	fmt.Println("toLower:", str3) // Output: "mayer"

	// Convierte todos los caracteres de una cadena a mayúsculas.
	uppercase := strings.ToUpper(name)
	fmt.Println("toUpper:", uppercase) // Output: "MAYER"

	// Función `Replace()`: Reemplaza una subcadena por otra dentro de la cadena original.
	// s: Es la cadena original.
	// old: Es la cadena que quieres reemplazar.
	// new: La nueva subcadena que va a reemplazar a `old`.
	// n: indica cuántas veces realizar el reemplazo; si es -1, se reemplazan todas las apariciones.
	fruta := "coco"
	str4 := strings.Replace(fruta, "o", "a", -1)
	fmt.Printf("replace: %q\n", str4) // Output: "caca"

	// Divide una cadena en partes usando un separador.
	// Devuelve un slice de strings.
	str5 := strings.Split(name, "")
	fmt.Printf("split: %q\n", str5) // Output: ["M" "a" "y" "e" "r"]

	// Verifica si una cadena termina con una subcadena.
	// Devuelve true o false.
	str6 := strings.HasSuffix(name, "er")
	fmt.Println("hasSuffix:", str6) // Output: true

	// Verifica si una cadena comienza con una subcadena.
	// Devuelve true o false.
	str7 := strings.HasPrefix(name, "ma")
	fmt.Println("hasPrefix:", str7) // Output: false

	// Elimina los espacios en blanco del inicio y final de una cadena.
	str8 := strings.TrimSpace(name)
	fmt.Println("trimSpace:", str8)

	// Elimina caracteres específicos del inicio y final de una cadena.
	lastname := "!!Chaves¡¡"
	str9 := strings.Trim(lastname, "!¡")
	fmt.Println("trim:", str9) // Output: Chaves

	// Une los elementos de un `slice` en una sola cadena, usando un separador.
	// Devuelve una nueva cadena con los elementos unidos con el separador.
	sli := []string{"foo", "bar", "baz"}
	fmt.Printf("join: %q\n", strings.Join(sli, ", ")) // Output: "foo, bar, baz"

	// Función `Clone` (Introducida en Go 1.18).
	// Devuelve una copia de la cadena.
	// Útil para reducir el uso de memoria en algunos casos.
	animal := "León"
	clone := strings.Clone(animal)
	fmt.Println("clone:", clone) // Output: "León"

	// Función `Cut` (Go 1.18).
	// Divide una cadena en dos partes usando un separador.
	// Retorna la parte antes del separador, la parte después y un booleano que indica si se encontró el separador.
	s := "Hola mundo"
	sep := " "
	before, after, found := strings.Cut(s, sep)
	fmt.Println("Antes del separador:", before)   // Output: "Hola"
	fmt.Println("Después del separador:", after)  // Output: "mundo"
	fmt.Println("¿Separador encontrado?:", found) // Output: true

	// Función `CutPrefix` (Go 1.20).
	// Elimina un prefijo específico del inicio de una cadena.
	// Prefijo: Una cadena que aparece al principio de otra.
	// Retorna la cadena sin el prefijo y un booleano que indica si el prefijo fue encontrado.
	gopher := "gopher"
	prefix := "go"
	after, ok := strings.CutPrefix(gopher, prefix)
	fmt.Println("Después del prefijo:", after) // Output: "opher"
	fmt.Println("¿Prefijo encontrado?:", ok)   // Output: true

	// Función `CutSuffix` (Go 1.20).
	// Elimina un sufijo específico del final de una cadena.
	// Sufijo: Una cadena que aparece al final de otra.
	// Retorna la cadena sin el sufijo y un booleano que indica si el sufijo fue encontrado.
	document := "documento.txt"
	suffix := ".txt"
	newS, ok := strings.CutSuffix(document, suffix)
	fmt.Println("Después del sufijo:", newS) // Output: "documento"
	fmt.Println("¿Sufijo encontrado?:", ok)  // Output: true

	//* Paquete "strconv": Conversión entre strings y tipos numéricos.

	// Convierte un string que representa un número entero a un tipo `int`.
	age := "23"
	age2, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tipo: %T - valor: %v\n", age2, age2) // Output: tipo: int - valor: 23

	// Convierte un string que representa un número flotante a un `float64`.
	weight := "70.5"
	weight2, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tipo: %T - valor: %v\n", weight2, weight2) // Output: tipo: float64 - valor: 70.5

	//* Uso de `strings.Builder` para concatenar cadenas de manera eficiente.
	// Trabaja con un `buffer` interno que minimiza la creación de nuevos strings, mejorando el rendimiento.
	var sb strings.Builder // sb := new(strings.Builder)

	// `WriteString()` agrega el contenido de la cadena al buffer interno de strings.Builder.
	sb.WriteString("Hola")
	sb.WriteString(", ")
	sb.WriteString("mundo!")

	// El método `String()` convierte el contenido interno del buffer de `strings.Builder` en una `cadena`.
	resultado := sb.String()
	fmt.Println("resultado:", resultado) // Output: Hola, mundo!
}
