package main

import "fmt"

/*
* Concatenar slices:
Crea una función que tome dos slices de enteros y los concatene en uno solo.
Asegúrate de que el slice resultante tenga una longitud igual a la suma de las longitudes de los dos slices originales,
y que su capacidad sea igual a la capacidad del primer slice más la capacidad del segundo slice.
*/

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}

	result := concatenarSlices(slice1, slice2)

	fmt.Println("Longitud:", len(result))  // Longitud: 10
	fmt.Println("Capacidad:", cap(result)) // Capacidad: 10
	fmt.Println(result)                    // [1 2 3 4 5 6 7 8 9 10]
}

func concatenarSlices(slice1, slice2 []int) []int {
	sumaCapacidades := cap(slice1) + cap(slice2)

	// Inicializar result con la longitud adecuada.
	result := make([]int, len(slice1), sumaCapacidades)

	// for _, value := range slice1 {
	// 	result = append(result, value)
	// }

	copy(result, slice1)

	result = append(result, slice2...)

	return result
}
