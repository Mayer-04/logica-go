package main

import "fmt"

/*
* Funciones Incorporadas en Go
Go proporciona una serie de funciones `predefinidas` que se pueden utilizar directamente sin necesidad de importarlas.

* Lista de algunas funciones incorporadas en Go:
- append(): Agrega uno o más elementos al final de un slice, devolviendo un nuevo slice.
- copy(): Copia elementos de un slice a otro y retorna el número de elementos copiados.
- cap(): Retorna la capacidad (capacidad total posible) de un slice.
- close(): Cierra un canal, lo que indica que no se enviarán más valores por él.
- delete(): Elimina un par clave-valor de un map.
- len(): Retorna la longitud (número de elementos) de un slice, map, canal, array o string.
- new(): Asigna memoria para un nuevo valor de tipo y retorna un puntero a este.
- panic(): Lanza una excepción en tiempo de ejecución, deteniendo la ejecución de la goroutina actual.
- recover(): Recupera de un panic, permitiendo que el programa continúe ejecutándose.
- make(): Asigna e inicializa slices, maps o canales.
- real(): Retorna la parte real de un número complejo.
- print() y println(): Imprimen valores en la salida estándar, con y sin salto de línea, respectivamente.
*/

func main() {
	// `print()` imprime un texto sin un salto de línea al final.
	print("Built-in Functions!")

	// `println()` imprime un texto con un salto de línea al final.
	println("¡Hola, Go!")

	// `append()` agrega elementos al final de un slice y modifica el slice original.
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println("append:", slice)

	// `len()` devuelve el número de elementos en un slice, array, map, string y channel.
	// En slices y maps devuelve el número de elementos en `v`.
	// En strings devuelve el número de bytes en `v`.
	fmt.Println("len:", len(slice))

	// `cap()` devuelve la capacidad total de un slice, es decir, el número máximo de elementos que puede contener sin reasignación.
	// En arrays, devuelve el número de elementos en `v`. Igual que `len(v)`.
	// En slices, devuelve la longitud máxima que puede alcanzar el slice si decides recortar (es decir, hacer más grande) el slice.
	fmt.Println("cap:", cap(slice))

	// `copy()` copia elementos de un slice a otro. Devuelve el número de elementos copiados.
	src := []string{"a", "b", "c"}
	dst := make([]string, len(src)) // Crear un slice destino con la misma longitud que el origen.
	n := copy(dst, src)             // Los elementos de `src` se copian en la variable `dst`.
	fmt.Println("copy:", dst, "número de elementos copiados:", n)

	// `close()` cierra un canal, lo que indica que no se enviarán más valores a través de él.
	// IMPORTANTE: Leer o enviar a un canal cerrado causará un panic.
	ch := make(chan int)
	close(ch)

	// `delete()` elimina un par clave-valor de un map a partir de su clave.
	m := map[string]int{"a": 1, "b": 2}
	delete(m, "a")
	fmt.Println("delete:", m) // Output: delete: map[b:2]

	// `new()` asigna memoria para un nuevo valor de tipo y retorna un puntero a este.
	newInt := new(int)
	fmt.Println("new:", newInt, "valor:", *newInt) // Output: new: 0xc00000a150 valor: 0

	// `make()` crea y inicializa un slice, map o canal.
	newSlice := make([]int, 5) // Crea un slice de enteros con longitud 5
	fmt.Println("make:", newSlice)

	// `max()` retorna el número mayor o el más grande de varios argumentos numéricos.
	fmt.Println("max:", max(1, 2, 4, 6, 7, 8, 9)) // Output: 9

	// `min()` retorna el número menor o el más pequeño de varios argumentos numéricos.
	fmt.Println("min:", min(1, 2, 0, 5, 10, 9)) // Output: 0

	// `real()` retorna la parte real de un número complejo.
	complexNum := 1 + 2i
	fmt.Println("real:", real(complexNum))

	// `panic()` lanza una excepción en tiempo de ejecución, deteniendo la ejecución de la goroutina actual.
	// `recover()` se usa para recuperar de un panic, permitiendo que el programa continúe ejecutándose.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se produjo un error:", r)
		}
	}()
	panic("Error de ejecución")
}
