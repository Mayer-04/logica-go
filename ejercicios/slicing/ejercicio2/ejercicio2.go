package main

import "fmt"

/*
* Ejercicio 2: Slice sin modificar el original
Tienes un slice de strings con nombres de frutas.
Extrae un nuevo slice con las 3 primeras frutas sin modificar el original.
Imprime ambos para comprobar que el original sigue intacto.

👉 Objetivo: Comprender cómo slicing no copia, sino que referencia.
*/

func main() {
	frutas := []string{"manzana", "pera", "mango", "uvas", "sandia", "banano"}

	// Nuevo slice
	nuevasFrutas := make([]string, 0, 3)
	nuevasFrutas = append(nuevasFrutas, frutas[:3]...)

	fmt.Println("frutas", frutas)
	fmt.Println("nuevasFrutas", nuevasFrutas)
}
