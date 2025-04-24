package main

import "fmt"

/*
Crea una función que tome un slice de enteros y devuelva un mapa.
Donde la clave es el entero y el valor es la cantidad de veces que aparece en el slice.
*/

func main() {
	enteros := []int{1, 2, 3, 1, 2, 3, 5, 6, 7, 8, 9, 10}
	resultado := contarCaracteresSli(enteros)
	fmt.Printf("resultado: %#v", resultado) // map[int]int{1:2, 2:2, 3:2, 5:1, 6:1, 7:1, 8:1, 9:1, 10:1}
}

func contarCaracteresSli(sli []int) map[int]int {
	newMap := make(map[int]int)

	for _, entero := range sli {
		// Si la clave 'entero' no existe en el mapa, Go automáticamente inicializa su valor a 0.
		// En este caso, newMap[entero] devolverá 0 inicialmente `newMap[entero] = 0`.
		// Luego, incrementamos ese valor en 1 cada vez que encontramos el mismo entero en el slice.
		newMap[entero] = newMap[entero] + 1
	}

	return newMap
}
