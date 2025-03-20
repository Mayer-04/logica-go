package main

import "fmt"

/*
Crea un slice que contenga el slice nums duplicado.
Ejemplo: Si nums = [1,2,3], el resultado es [1,2,3,1,2,3].
La capacidad del slice resultante es el doble de la longitud del slice nums.
*/

func main() {
	nums := []int{1, 2, 3}
	result := getConcatenation(nums)
	fmt.Println(result)
}

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n) // Crear el array con tamaño 2n

	// Usamos copy() para copiar nums en la primera y segunda mitad de ans
	copy(ans[0:n], nums)   // Copia la primera mitad
	copy(ans[n:2*n], nums) // Copia la segunda mitad

	return ans
}

// Segunda solución ✅
func GetConcatenation2(nums []int) []int {
	return append(nums, nums...) // Usamos append para concatenar nums consigo mismo
}

// Tercera solución ✅
func GetConcatenation3(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n) // Crear el array ans con tamaño 2n

	// Copiar los valores en la primera y segunda mitad de ans
	for i := range n {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}

	return ans
}
