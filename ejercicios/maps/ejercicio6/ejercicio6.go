package main

import "fmt"

/*
Tenemos un mapa que almacena los salarios de nuestro equipo:

salaries := map[string]int{
		"John": 100,
		"Ann":  160,
		"Pete": 130,
}

Escribe el código para sumar todos los salarios y almacenar el resultado en la variable suma.
En el ejemplo de arriba nos debería dar 390.

Si salaries está vacío entonces el resultado será 0.
*/

func main() {
	salaries := map[string]int{
		"John": 100,
		"Ann":  160,
		"Pete": 130,
	}
	result := sumarSalarios(salaries)
	fmt.Println(result)
}

func sumarSalarios(m map[string]int) int {
	suma := 0
	for _, valor := range m {
		suma += valor
	}
	return suma
}
