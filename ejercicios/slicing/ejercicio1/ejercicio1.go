package main

import "fmt"

/*
* Ejercicio 1: Obtener una parte del slice
Tienes un slice de enteros del 1 al 10.
Extrae un nuevo slice que contenga solo los números del 4 al 7 (inclusive).

👉 Objetivo: Usar slicing para tomar una subparte específica de un slice.
*/

func main() {
	enteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sli := enteros[3:7]
	fmt.Println("sli", sli)

	str := "Hola mundo"
	sli2 := str[:4]
	fmt.Println("sli2", sli2)
}
