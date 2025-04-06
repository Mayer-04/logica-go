package main

import "fmt"

/*
Crea una función que reciba un map[int]string y un string.
La función debe eliminar todas las entradas del map que tengan ese string.
*/

func main() {
	ciudades := map[int]string{
		1: "Bogotá",
		2: "Cali",
		3: "Medellin",
		4: "Pasto",
		5: "Bucaramanga",
	}

	ciudad := "Medellin"
	eliminarValor(ciudades, ciudad)
	fmt.Printf("cuidades: %v", ciudades) // cuidades: map[1:Bogotá 2:Cali 4:Pasto 5:Bucaramanga]
}

func eliminarValor(m map[int]string, s string) {
	for key, value := range m {
		if value == s {
			delete(m, key)
		}
	}
}
