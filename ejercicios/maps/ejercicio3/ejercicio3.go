package main

import "fmt"

/*
Crea un programa que reciba un map donde las claves son cadenas y los valores enteros.
El programa debe devolver un nuevo map donde los valores sean las claves y las claves sean los valores.
Por ejemplo, si el map es {"a": 1, "b": 2}, el nuevo map debería ser {1: "a", 2: "b"}.
*/

func main() {
	m := map[string]int{"a": 1, "b": 2}
	result := invertirMapa(m)
	fmt.Println("mapa original:", m)
	fmt.Println("mapa invertido:", result)
}

func invertirMapa(m map[string]int) map[int]string {
	nuevoMapa := make(map[int]string, len(m))
	for k, v := range m {
		nuevoMapa[v] = k
	}
	return nuevoMapa
}
