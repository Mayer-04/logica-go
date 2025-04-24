package main

import "fmt"

/*
Imprimir un triángulo de asteriscos.
Escribe un programa que, dado un número n, imprima un triángulo de asteriscos de altura n. Por ejemplo, para n=5:

*
**
***
****
*****
*/

func main() {
	n := 5
	buildTriangle(n)
}

func buildTriangle(n int) {
	const defaultHeight = 5
	if n <= 0 {
		n = defaultHeight
	}

	asterisks := ""
	for range n {
		asterisks += "*"
		fmt.Println(asterisks)
	}
}
