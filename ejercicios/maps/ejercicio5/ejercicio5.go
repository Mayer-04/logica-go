package main

import "fmt"

/*
* Filtrado de un mapa
Crea una función que reciba un mapa y una función de filtro, y devuelva un nuevo mapa que contenga solo los
elementos que cumplan con la condición definida por el filtro.
Parámetros:
- El primer parámetro es un mapa de cadenas a enteros (map[string]int), donde las claves son los nombres de los
elementos y los valores son cantidades o cualquier valor numérico asociado.
- El segundo parámetro es una función de filtro (func(int) bool), que toma un entero como entrada y devuelve true
si el valor cumple con el criterio de filtrado, o false en caso contrario.
*/

func main() {
	inventario := map[string]int{
		"manzanas": 50,
		"plátanos": 20,
		"naranjas": 30,
		"peras":    15,
		"uvas":     40,
	}

	// Ejemplo 1: Filtrar frutas con cantidad mayor a 25
	frutasAbundantes := filtrarMapa(inventario, func(cantidad int) bool {
		return cantidad > 25
	})
	fmt.Println("Frutas con cantidad mayor a 25:")
	fmt.Println(frutasAbundantes)

	// Ejemplo 2: Filtrar frutas con cantidad par
	frutasPares := filtrarMapa(inventario, func(cantidad int) bool {
		return cantidad%2 == 0
	})
	fmt.Println("\nFrutas con cantidad par:")
	fmt.Println(frutasPares)
}

func filtrarMapa(mapa map[string]int, filtro func(int) bool) map[string]int {
	resultado := make(map[string]int)
	for clave, valor := range mapa {
		if filtro(valor) {
			resultado[clave] = valor
		}
	}

	return resultado
}
