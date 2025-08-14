package main

import "fmt"

/*
* Agrupar palabras por longitud
Dado un slice de palabras, utiliza un mapa para agrupar las palabras según su longitud.
El mapa debe tener como clave la longitud de la palabra y como valor un slice de palabras con esa longitud.

Ejemplo:

Entrada: ["go", "java", "python", "rust"]
Salida: {2: ["go"], 4: ["java", "rust"], 6: ["python"]}
*/

func main() {
	lenguajes := []string{"go", "java", "python", "rust"}
	fmt.Println(agruparPorLongitud(lenguajes))
}

func agruparPorLongitud(lenguajes []string) map[int][]string {
	nuevoMapa := make(map[int][]string)
	for _, lenguaje := range lenguajes {
		longitud := len(lenguaje)
		if _, ok := nuevoMapa[longitud]; !ok {
			nuevoMapa[longitud] = make([]string, 0, len(lenguajes))
		}
		nuevoMapa[longitud] = append(nuevoMapa[longitud], lenguaje)
	}

	return nuevoMapa
}
