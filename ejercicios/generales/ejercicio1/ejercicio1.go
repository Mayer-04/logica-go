package main

import (
	"errors"
	"fmt"
)

//* Suma de enteros pares: Escribe un programa que sume todos los números pares en un rango dado de enteros.

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	result, err := addEvenNumbers(numbers)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func addEvenNumbers(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("no hay números en el slice")
	}

	sum := 0
	for _, value := range nums {
		if value%2 == 0 {
			sum += value
		}
	}

	return sum, nil
}

// NOTA: Esta solución solo es válida en nuestro ejercicio.
func addEvenNumbers2(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("no numbers in slice")
	}

	var sum int
	for i := 1; i < len(numbers); i += 2 {
		sum += numbers[i]
	}

	return sum, nil
}
