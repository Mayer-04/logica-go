package main

import "fmt"

/*
* Fusión de mapas
Crea una función que reciba dos mapas y los fusione en uno solo.
Si una clave existe en ambos mapas, el valor del segundo mapa debe prevalecer.
*/

func main() {
	mapa1 := map[string]int{"a": 1, "b": 2, "c": 3}
	mapa2 := map[string]int{"a": 4, "e": 5}
	fmt.Println(fusionMapas(mapa1, mapa2))
}

func fusionMapas(mapa1, mapa2 map[string]int) map[string]int {
	longitudTotal := len(mapa1) + len(mapa2)
	nuevoMapa := make(map[string]int, longitudTotal)

	for c, v := range mapa1 {
		nuevoMapa[c] = v
	}

	for c, v := range mapa2 {
		nuevoMapa[c] = v
	}

	return nuevoMapa
}