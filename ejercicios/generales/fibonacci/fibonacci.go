package main

import "fmt"

/*
Escribe una función en Python que calcule los primeros n términos de la secuencia de Fibonacci.
La secuencia de Fibonacci se define de la siguiente manera:

Los dos primeros términos son 0 y 1.
A partir del tercer término, cada número es la suma de los dos anteriores.
Por ejemplo, la secuencia de Fibonacci para n = 10 es:

0, 1, 1, 2, 3, 5, 8, 13, 21, 34
*/

func getFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0}
	}

	if n == 2 {
		return []int{0, 1}
	}

	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1

	a, b := 0, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
		fib[i] = b
	}

	return fib
}

func main() {
	n := 10
	result := getFibonacci(n)
	fmt.Println(result)
}
