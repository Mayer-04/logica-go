package main

import (
	"fmt"
	"math"
)

/*
* Built-in Functions - Funciones incorporadas en Go

Go proporciona una serie de funciones predefinidas que se pueden usar sin necesidad de importarlas.

* Principales funciones incorporadas:

- append(): Agrega uno o más elementos al final de un slice, devolviendo un nuevo slice.
- copy(): Copia elementos de un slice a otro. Devuelve el número de elementos copiados.
- cap(): Devuelve la capacidad de un slice, array o canal.
- len(): Devuelve la longitud de un slice, array, mapa, string o canal.
- make(): Crea e inicializa slices, maps o canales.
- new(): Asigna memoria para un tipo y devuelve un puntero al valor con cero inicializado.
- delete(): Elimina un par clave-valor de un mapa.
- close(): Cierra un canal para indicar que no se enviarán más valores.
- panic(): Genera un error en tiempo de ejecución que interrumpe la ejecución.
- recover(): Recupera de un panic dentro de una función `defer`, permitiendo que el programa continúe.
- real(): Devuelve la parte real de un número complejo.
- print() y println(): Imprimen valores en la salida estándar, con y sin salto de línea respectivamente.
*/

func main() {
	// `println()` Imprime un texto con un salto de línea al final.
	println("¡Hola, Go!")

	// `print()` Imprime un texto sin un salto de línea.
	print("Built-in Functions!")

	// `append()` Agrega elementos al final de un slice y modifica el slice original.
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println("append:", slice)

	// `len()` Devuelve el número de elementos en un slice, array, map, string y channel.
	// En slices y maps devuelve el número de elementos en `v`.
	// En strings devuelve el número de bytes en `v`.
	fmt.Println("len:", len(slice))

	// `cap()` Devuelve la capacidad total de un slice, es decir, el número máximo de elementos que puede contener sin reasignación.
	// En arrays, devuelve el número de elementos en `v`. Igual que `len(v)`.
	// En slices, devuelve la longitud máxima que puede alcanzar el slice si decides recortar (es decir, hacer más grande) el slice.
	fmt.Println("cap:", cap(slice))

	// `copy()` Copia elementos de un slice a otro. Devuelve el número de elementos copiados.
	src := []string{"a", "b", "c"}
	dst := make([]string, len(src)) // Crear un slice destino con la misma longitud que el origen.
	n := copy(dst, src)             // Los elementos de `src` se copian en la variable `dst`.
	fmt.Println("copy:", dst, "número de elementos copiados:", n)

	// `close()` Cierra un canal, lo que indica que no se enviarán más valores a través de él.
	//* IMPORTANTE: Leer o enviar a un canal cerrado causará un `panic`.
	ch := make(chan int)
	close(ch)

	// `delete()` Elimina un par clave-valor de un mapa a partir de su clave.
	m := map[string]int{"a": 1, "b": 2}
	delete(m, "a")
	fmt.Println("delete:", m) // Output: delete: map[b:2]

	// `new()` Asigna memoria para un nuevo valor de tipo y retorna un puntero a este.
	// Crear un nuevo puntero.
	newInt := new(int)
	fmt.Println("new:", newInt, "valor:", *newInt) // Output: new: 0xc00000a150 valor: 0

	// `make()` Crea y inicializa un slice, mapa o canal.
	// Crea un slice de enteros con longitud 5.
	newSlice := make([]int, 5)
	fmt.Println("make:", newSlice) // Output: [0 0 0 0 0]

	// `max()` Encuentra el valor más grande de una lista de números o textos.
	// Puede usarse con enteros, flotantes o cadenas.
	fmt.Println("max:", max(1, 2, 4, 6, 7, 8, 9)) // Output: 9

	//* IMPORTANTE: Si NaN está en la lista, `min` y `max` devuelven NaN.
	fmt.Println(max(3.0, math.NaN(), 5.0)) // → NaN
	fmt.Println(min(3.0, math.NaN(), 5.0)) // → NaN

	// `min()` Encuentra el valor más pequeño de una lista de números o textos.
	// Puede usarse con enteros, flotantes o cadenas.
	fmt.Println("min:", min(1, 2, 0, 5, 10, 9))  // Output: 0
	fmt.Println(min("zebra", "apple", "banana")) // Output: apple

	// `real()` Devuelve la parte real de un número complejo.
	complexNum := 1 + 2i
	fmt.Println("real:", real(complexNum)) // Output: 1

	// `panic()` Lanza una excepción en tiempo de ejecución, deteniendo la ejecución de la goroutina actual.
	// `recover()` Se usa para recuperar de un panic, permitiendo que el programa continúe ejecutándose.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se produjo un error:", r)
		}
	}()
	panic("Error de ejecución")
}
