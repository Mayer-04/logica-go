package main

import "fmt"

/*
* Constantes en Go
- Las constantes en Go son valores `inmutables` que se definen con la palabra clave 'const'.
- Al ser inmutables, las constantes deben recibir un valor en el momento de su declaración.
- Las constantes en Go están limitadas a tipos básicos como cadenas de caracteres, enteros, flotantes y booleanos.
- Las constantes en Go deben conocerse completamente antes de que el programa se ejecute, es decir,
en tiempo de compilación.
- No es necesario utilizar todas las constantes definidas para que el programa compile.
- Se recomienda declarar constantes a nivel de paquete cuando se necesitan en varias partes del código.
- Las constantes no tienen un tipo fijo hasta que se usan, lo que permite flexibilidad al asignarlas.

* Identificador iota
'iota' es un identificador predefinido en Go que se usa para generar una secuencia de números para sus constantes.
- El valor de 'iota' comienza en 0 dentro de un bloque de constantes y se incrementa automáticamente con cada línea.
- Es útil para definir enumeraciones o listas secuenciales de valores relacionados.
- Si no necesitas algún valor en la secuencia, puedes usar el identificador blank `_` para omitirlo.
*/

// Constante a nivel de paquete.
const animal = "🐯"

// Constante exportable a nivel de paquete (inicia con mayúscula).
const PI = 3.1416

// Constante no exportable, solo visible dentro del paquete (inicia con minúscula).
const version = "1.0.0"

// Agrupación de constantes relacionadas en un bloque.
// Permite organizar y agrupar valores constantes que están conceptualmente relacionados.
const (
	fruit1 = "🍎"
	fruit2 = "🍐"
)

// Uso de iota para crear secuencias de valores incrementales.
// Ideal para enumeraciones, días de la semana, estados, etc.
const (
	Lunes     = iota // Lunes == 0
	Martes           // Martes == 1 (iota se incrementa automáticamente)
	Miércoles        // Miércoles == 2 (iota se incrementa automáticamente)
	Jueves           // Jueves == 3 (iota se incrementa automáticamente)
	Viernes          // Viernes == 4 (iota se incrementa automáticamente)
	Sábado           // Sábado == 5 (iota se incrementa automáticamente)
	Domingo          // Domingo == 6 (iota se incrementa automáticamente)
)

// REINICIO DE IOTA EN CADA BLOQUE CONST
// iota se reinicia a 0 en cada nueva declaración const
const (
	A = iota // A == 0
	B        // B == 1
)
const (
	C = iota // C == 0 (¡iota se reinicia!)
	D        // D == 1
)

// PROPAGAR VALORES EN CONSTANTES.
// Si no asignas un valor explícitamente a una constante, tomará el valor de la constante anterior.
const (
	Foo = 3.14 // Foo == 3.14
	Bar        // Bar == 3.14 (toma el valor de Foo)
	Baz        // Baz == 3.14 (toma el valor de Bar)
	Qux = 10   // Qux == 10
	Xyz        // Xyz == 10 (toma el valor de Qux)
)

// Usando iota para saltar valores con el operador blank (_).
const (
	_     = iota // Se omite el primer valor (0)
	One          // One == 1
	Two          // Two == 2 (iota se incrementa automáticamente)
	Three        // Three == 3 (iota se incrementa automáticamente)
)

// Definiendo constantes con operaciones matemáticas.
const (
	KB = 1 << (10 * iota) // 1 << 10 (1 KB = 1024 bytes)
	MB                    // 1 << 20 (1 MB = 1024 KB)
	GB                    // 1 << 30 (1 GB = 1024 MB)
	TB                    // 1 << 40 (1 TB = 1024 GB)
)

func main() {
	// Las constantes se declaran con la palabra clave 'const' seguida del nombre y, opcionalmente, el tipo.
	// No podemos utilizar el operador de asignación corta (:=) para definir una constante.
	const gender string = "Masculino"
	fmt.Println("género:", gender)

	// Imprimir los valores de las constantes agrupadas.
	fmt.Printf("frutas: %s, %s\n", fruit1, fruit2)

	// Imprimir los valores de las constantes definidas a nivel de paquete.
	fmt.Println("animal:", animal)
	fmt.Println("versión:", version)

	// Imprimir valores de constantes en un bloque sin asignación explícita.
	fmt.Printf("Foo: %v, Bar: %v, Baz: %v, Qux: %v, Xyz: %v\n", Foo, Bar, Baz, Qux, Xyz)

	// Imprimir valores de constantes que omiten el primer valor de iota.
	fmt.Printf("One: %v, Two: %v, Three: %v\n", One, Two, Three) // Output: One: 1, Two: 2, Three: 3

	// Imprimir valores de constantes definidas con operaciones matemáticas utilizando iota.
	fmt.Printf("1 KB = %v bytes\n", KB) // 1 KB = 1 bytes
	fmt.Printf("1 MB = %v bytes\n", MB) // 1 MB = 1024 bytes
	fmt.Printf("1 GB = %v bytes\n", GB) // 1 GB = 1048576 bytes
	fmt.Printf("1 TB = %v bytes\n", TB) // 1 TB = 1073741824 bytes

	// Se pueden `determinar en tiempo en ejecución`.
	//* NOTA: Tiempo de compilación → Es cuando el programa se traduce a código ejecutable.
	const firstName = "Mayer"
	const lastName = "Chaves"
	// Go puede hacer esta operación antes de ejecutar el programa
	// porque los valores ya están definidos y no cambian.
	const fullName = firstName + " " + lastName
	fmt.Println("fullname:", fullName) // Output: Mayer Chaves
}
