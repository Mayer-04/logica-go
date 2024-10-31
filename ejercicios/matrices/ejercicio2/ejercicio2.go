package main

import "fmt"

/*
Recibimos una matriz de cadenas de texto.
Tenemos que localizar la posición de la palabra "Go" en la matriz y devolver su posición como un array
de dos elementos: el índice de la fila y el índice de la columna.

Si no encuentra la palabra debe devolver [-1, -1].

languages := [][]string{
	{"HTML", "CSS", "JavaScript"},
	{"Java", "C++", "Python"},
	{"Ruby", "Go", "Rust"},
}
*/

func main() {
	languages := [][]string{
		{"HTML", "CSS", "JavaScript"},
		{"Java", "C++", "Python"},
		{"Ruby", "Go", "Rust"},
	}
	fmt.Println(findGo(languages)) // [2, 1]
}

func findGo(languages [][]string) []int {
	for i, row := range languages {
		for j, language := range row {
			if language == "Go" {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
