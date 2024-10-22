package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	number := 45
	result := multiplesOperations(number)
	fmt.Printf("%#v", result)
}

func multiplesOperations(number int) map[string]any {
	result := make(map[string]any)

	// Convierte el número en un string
	str := strconv.Itoa(number)

	if len(str) != 2 {
		fmt.Println("El número debe tener exactamente dos dígitos.")
		return nil
	}

	// Convierte el string en reverso
	var reversedStr string
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	result["reverse"] = reversedStr

	// Convertimos los caracteres en enteros
	a, err := strconv.Atoi(string(str[0]))
	if err != nil {
		return nil
	}
	b, err := strconv.Atoi(string(str[1]))
	if err != nil {
		return nil
	}

	// Realizamos las operaciones
	// Asignamos los resultados al map
	result["sum"] = a + b
	result["subtraction"] = a - b
	result["multiplication"] = a * b
	// Pendiente validación por si alguno de los dígitos es 0
	result["division"] = float64(a) / float64(b)
	result["pow"] = int(math.Pow(float64(a), float64(b)))
	result["module"] = a % b

	return result
}
