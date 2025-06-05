package main

import (
	"fmt"
)

/*
* Números en Go
En Go, los números se representan con diferentes tipos de datos: `int`, `uint`, `float32`, `float64`, etc.

- Si definimos `uint` o `int`, el tamaño del valor dependerá del sistema operativo: 32 bits o 64 bits.
- Para mejorar la legibilidad de números grandes, podemos utilizar el guion bajo `_`.
- En operaciones con `float`, se pueden obtener valores especiales como `+Inf`, `-Inf` y `NaN`.
*/

func main() {
	// Números enteros positivos (8 bits).
	var number uint8 = 10
	fmt.Printf("Tipo: %T - Valor: %d\n", number, number) // Output: Tipo: uint8 - Valor: 10

	// Uso del guion bajo para mejorar la legibilidad de números grandes.
	bigNumber := 1_000_000
	fmt.Println("Número grande:", bigNumber) // Output: Número grande: 1000000

	// Números enteros con signo (positivos y negativos, 8 bits).
	var number2 int8 = -10
	fmt.Printf("Tipo: %T - Valor: %d\n", number2, number2)

	// Números decimales de 32 bits.
	var number3 float32 = 10.5
	fmt.Printf("Tipo: %T - Valor: %.1f\n", number3, number3) // Output: Tipo: float32 - Valor: 10.5

	// Números decimales de 64 bits.
	var number4 float64 = 10.5555
	fmt.Printf("Tipo: %T - Valor: %f\n", number4, number4) // Output: Tipo: float64 - Valor: 10.555500

	//* Conversión de tipos de datos (casting).
	// Convertimos un valor de un tipo a otro de manera explícita.
	var a int16 = 5000
	var b uint8 = 10

	// Se realiza la conversión de `b` a `int16` antes de realizar la suma.
	var c = int16(b) + a
	fmt.Printf("Casting: %T - Valor: %d\n", c, c) // Output: Casting: int16 - Valor: 5010

	// * Notación científica.
	// Se utiliza para representar números grandes o pequeños.
	var scientificNo int = 1e3                           // 1e3 significa 1 × 10³, es decir, 1000.
	fmt.Printf(("%T: %d\n"), scientificNo, scientificNo) // Output: int: 1000

	var scientificSmall = 5e-3                                   // 5e-3 significa 5 × 10^-3, es decir, 0.005
	fmt.Printf(("%T: %.3f\n"), scientificSmall, scientificSmall) // Output: float64: 0.005

	// Valores especiales en operaciones con flotantes.
	zero := 0.0
	fmt.Println("División de 1 entre 0:", 1/zero)    // Output: +Inf
	fmt.Println("División de -1 entre 0:", -1/zero)  // Output: -Inf
	fmt.Println("División de 0 entre 0:", zero/zero) // Output: NaN
}
