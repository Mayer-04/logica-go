package main

import "fmt"

/*
Dado un año, regresa el siglo  en el que está.
El primer siglo abarca desde el año 1 hasta el año 100, el segundo, desde el año 101 hasta el año 200, etc.
*/

func solution(year int) int {
	century := year / 100
	if year%100 == 0 {
		century--
	}
	return century + 1 // return (year + 99) / 100
}

func main() {
	fmt.Println(solution(1905)) // 20
}
