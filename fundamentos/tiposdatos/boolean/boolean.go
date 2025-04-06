package main

import "fmt"

func main() {
	// declarando una variable de tipo booleana
	var active bool = true

	// Imprime el formato de la variable - %T es para el tipo de dato y %t es para el valor booleano.
	fmt.Printf("Tipo: %T - Valor: %t\n", active, active)

	var active2 bool = false

	fmt.Printf("Tipo: %T - Valor: %t\n", active2, active2)

	// active3: Inferiendo el tipo de dato
	var active3 = true
	fmt.Println(active3)
}
