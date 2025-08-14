package main

import (
	"fmt"
	"unsafe"
)

/*
* Paquete unsafe: Arriesgado
El paquete `unsafe` en Go permite realizar operaciones de bajo nivel que interactúan directamente con la memoria,
saltando las garantías de seguridad y portabilidad del lenguaje. Esto puede ser útil para optimizaciones avanzadas,
aunque se debe usar con precaución.

Estos tipos son parte de las funciones de `unsafe`:
- type ArbitraryType: Se utiliza para indicar que el tipo de dato puede ser cualquiera (flexibilidad de tipos).
- type IntegerType: Se utiliza para indicar que el tipo de dato es un valor entero.
- type PointerType: Representa un puntero de tipo ArbitraryType.
Se utiliza para hacer conversiones entre diferentes tipos de punteros.

* `unsafe.Pointer`:
- Representa un puntero genérico que puede ser convertido a otros tipos de punteros.
- Es un tipo especial de puntero que puede ser usado para realizar conversiones entre diferentes tipos de punteros,
permitiendo manipular la memoria directamente.
- Cuando lo usas, le estás diciendo al compilador Go: “sé lo que estoy haciendo, así que confía en mí”.
* Ejemplo:
Puedes cambiar un puntero *int a un unsafe.Pointer y luego cambiar ese unsafe.Pointer a un puntero *float64.


* `uintptr`:
- Es un entero sin signo lo suficientemente grande para almacenar un puntero.
- Aunque puede contener una dirección de memoria, `uintptr` no es un puntero real,
por lo que no se garantiza que el valor almacenado en él apunte a una dirección válida durante mucho tiempo.
- No se recomienda convertir un `uintptr` a un `unsafe.Pointer`.


* Aritmética de punteros:
- Técnica avanzada que permite manipular directamente direcciones de memoria,
como sumar o restar para moverse entre posiciones de memoria.
- Aunque Go no permite la aritmética de punteros de manera directa,
con `unsafe` es posible emularla utilizando conversiones a `uintptr`.
- Manipula directamente las direcciones de memoria para acceder a diferentes elementos.
*/

func main() {
	// `unsafe.Sizeof()`
	// Devuelve el tamaño en bytes de una variable.
	// Es útil para ver cuánto espacio ocupa un tipo.
	// No considera el contenido dinámico de slices o strings, sino solo su estructura interna.
	text := "Hello, World!"
	number := 10
	life := true
	animals := []string{"cat", "dog", "fish"}

	fmt.Printf("El tamaño de un string es: %d bytes\n", unsafe.Sizeof(text))             // 16 bytes
	fmt.Printf("El tamaño de un int es: %d bytes\n", unsafe.Sizeof(number))              // 8 bytes
	fmt.Printf("El tamaño de un bool es: %d bytes\n", unsafe.Sizeof(life))               // 1 byte
	fmt.Printf("El tamaño de un slice de string es: %d bytes\n", unsafe.Sizeof(animals)) // 24 bytes

	//* `unsafe.Alignof()`.
	// Devuelve el número de bytes que una variable debe estar `alineada` en memoria.
	// La alineación es una propiedad que asegura que los datos estén organizados en memoria de manera eficiente,
	// permitiendo un acceso más rápido por parte del procesador.

	// Por ejemplo, en un sistema de 64 bits:
	// - Los datos de 64 bits (como un `int` o un `float64`) suelen estar alineados a 8 bytes.
	// - Los datos de 32 bits (como un `int32` o un `float32`) suelen estar alineados a 4 bytes.
	// - Los datos de 8 bits (como un `byte` o `bool`) suelen estar alineados a 1 byte.
	fmt.Printf("El alineamiento de un string es: %d bytes\n", unsafe.Alignof(text))             // 8 bytes
	fmt.Printf("El alineamiento de un int es: %d bytes\n", unsafe.Alignof(number))              // 8 bytes
	fmt.Printf("El alineamiento de un bool es: %d bytes\n", unsafe.Alignof(life))               // 1 byte
	fmt.Printf("El alineamiento de un slice de string es: %d bytes\n", unsafe.Alignof(animals)) // 8 bytes

	//* `unsafe.Pointer()`.
	// Permite convertir un puntero de un tipo específico (como `*int`) a un puntero genérico (`unsafe.Pointer`).
	// Un Pointer puede convertirse en un valor puntero de cualquier tipo.
	// Un uintptr puede convertirse en un puntero.
	// Esto es útil cuando necesitas convertir entre diferentes tipos de punteros,
	// lo cual normalmente no es permitido en Go.
	var x int = 42
	p := unsafe.Pointer(&x) // Convertimos el puntero `*int` a `unsafe.Pointer`
	// Convertimos `unsafe.Pointer` a `uintptr`, que es solo un número que representa la dirección de memoria.
	uptr := uintptr(p)

	// Ahora `uptr` es un número que representa la dirección de memoria de `x`, como si fuera solo un número entero.
	fmt.Printf("Dirección de memoria de x en formato numérico: %d\n", uptr)

	// Ejemplo adicional para mostrar cómo se ve un puntero sin inicializar (apunta a `nil` por defecto)
	var pointer *int
	fmt.Printf("El puntero sin inicializar apunta a: %v\n", unsafe.Pointer(pointer))

	//* Manipulación de cadenas y slices usando `unsafe`.

	//* `unsafe.StringData()`.
	// Devuelve un puntero al primer byte de los datos subyacentes de una cadena.
	// Manipular directamente estos datos es peligroso y puede causar comportamientos indefinidos.
	str := "Hello, World!"
	fmt.Printf("Puntero al primer byte de los datos internos de la cadena: %v\n", unsafe.StringData(str))

	//* `unsafe.Slice()`.
	// Crea un nuevo slice de un tipo específico a partir de un puntero y una longitud.
	// Esto permite manipular la memoria como un slice.
	sli := unsafe.Slice(&str, 1)
	fmt.Printf("Nuevo slice creado a partir de un puntero y longitud: %v\n", sli)

	//* `unsafe.SliceData()`.
	// Devuelve un puntero al primer elemento de los datos subyacentes de un slice.
	sliData := unsafe.SliceData(sli)
	fmt.Printf("Puntero al primer elemento de los datos subyacentes de un slice: %v\n", sliData)

	// Convertir un puntero *int64 a *float64.
	var a int64 = 10
	aPtr := unsafe.Pointer(&a)

	b := (*float64)(aPtr)
	fmt.Printf("Nuevo puntero float64: %v", b)
}

// convertStringToBytes Convierte una cadena en un slice de bytes.
//
// El parámetro de entrada `s` es la cadena a convertir.
// Devuelve un slice de bytes que contiene los bytes de la cadena de entrada,
// o nil si la cadena de entrada está vacía.
func StringToBytes(s string) []byte {
	// Verifica si la cadena de entrada no es nil.
	if s == "" {
		return nil
	}

	// Obtener el puntero al primer byte de la cadena.
	// Los bytes devueltos por `StringData` no deben ser modificados.
	ptr := unsafe.StringData(s)

	if ptr == nil {
		return nil
	}

	// Crea un slice de bytes que apunta a la misma posición de memoria que la cadena.
	// Si el puntero es nulo y la longitud es 0, la función `Slice` devuelve nil.
	return unsafe.Slice(ptr, len(s))
}
