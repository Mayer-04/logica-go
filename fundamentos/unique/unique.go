package main

import (
	"fmt"
	"unique"
)

/*
* Paquete unique: Unico
El paquete `unique` es una nueva característica de Go que se integró a partir de la versión 1.23.
- Se enfoca en la optimización de la memoria y la eficiencia, aprovechando una técnica conocida como "interning".
- Ahorro de memoria: Al almacenar una única copia de los valores repetidos, se reduce significativamente el uso de memoria.
- Se utiliza con cualquier valor comparable.
- Se ocupa de la concurrencia de manera automática (como lo hace con `sync.Map`).
- Evita el crecimiento descontrolado de memoria mediante una limpieza automática integrada con el garbage collector (GC) de Go.
- En lugar de comparar los bytes de los valores, compara sus "handles" (punteros), lo que resulta en comparaciones mucho más rápidas.
- El paquete es seguro para ser usado con múltiples goroutines.
- Cuando un valor internado ya no es necesario, el garbage collector lo elimina automaticamente.

* Interning: Internamiento
El interning es una técnica que garantiza que solo se guarde una copia de un valor repetido en memoria.
Si el mismo valor se solicita varias veces, en lugar de crear nuevas copias, se reutiliza la copia existente.
*/

func main() {
	/*
		* La función `Make()` permite internar cualquier tipo de valor comparable.
		- En lugar de guardar el valor directamente, esta función devuelve un "handle" (puntero o referencia)
		que es único para ese valor.
		- En lugar de comparar los valores reales, puedes comparar los handles. Si los handles son iguales,
		sabes que los valores subyacentes también lo son.
		- unique puede internar cualquier tipo de dato comparable.

		func Make[T comparable](value T) Handle[T] { ... }

		* `Handle` es una estructura que actúa como un identificador único o puntero para un valor internado.
		Cuando llamas a unique.Make(), obtienes un `Handle` en lugar del valor original.
		- En lugar de manejar directamente el valor (como una cadena, número, etc.), el paquete unique
		utiliza estos `Handle` para representar los valores.

		type Handle[T comparable] struct {
			value *T
		}
	*/

	// Internamiento de cadenas (string).
	name := unique.Make("Mayer")  // Interna el valor "Mayer"
	name2 := unique.Make("Mayer") // Dado que "Mayer" ya existe, name2 apunta a la misma ubicación en memoria

	// Verificamos que ambos apuntan al mismo "handle" (la referencia interna es la misma).
	// Las comparaciones entre "handles" son eficientes porque comparamos las direcciones de memoria.
	fmt.Println("name == name2:", name == name2) // true, ambas referencias son iguales
	fmt.Println("name:", name)                   // {0xc000086060}
	fmt.Println("name2:", name2)                 // {0xc000086060}

	// Internamiento de otra cadena (string).
	lastname := unique.Make("Prada")                   // Nuevo valor "Prada", distinta referencia
	fmt.Println("name == lastname:", name == lastname) // false, porque son valores diferentes

	// Internamiento de números (int).
	edad := unique.Make(24)                      // Interna el valor 24
	edad2 := unique.Make(24)                     // Reutiliza la referencia de 24
	fmt.Println("edad == edad2:", edad == edad2) // true, ambos apuntan al mismo valor internado
}
