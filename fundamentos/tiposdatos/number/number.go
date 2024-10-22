package main

import "fmt"

/*
* Números en Go
- Si definimos `uint` o `int`, el valor dependera del sistema operativo, si es 32 bits o 64 bits.
- Si queremos que los números más grandes sean legibles podemos usar el guion bajo _.
*/

func main() {
	// Números enteros positivos.
	var number uint8 = 10
	fmt.Printf("Tipo: %T - Valor: %d\n", number, number)

	// Utilizamos el guión bajo para leer mejor un número grande.
	bigNumber := 1_000_000
	fmt.Println("número grande:", bigNumber)

	// Números enteros negativos y positivos.
	var number2 int8 = -10
	fmt.Printf("Tipo: %T - Valor: %d\n", number2, number2)

	// Número decimal de 32 bits.
	var number3 float32 = 10.5
	fmt.Printf("Tipo: %T - Valor: %f\n", number3, number3)

	// Número decimales de 64 bits.
	var number4 float64 = 10.5
	fmt.Printf("Tipo: %T - Valor: %f\n", number4, number4)

	//* Casting de números.
	// Convierte un valor de un tipo a otro de manera `explícita`.
	var a int16 = 5000
	// El valor de b no cambia al hacer el casting.
	// Debe ser del mismo tipo de dato para hacer este casting.
	var b uint8 = 10

	//* Convierte un valor de un tipo a otro.
	// El valor de b es de 8 bits y se castea a 16 bits.
	c := int16(b) + a
	fmt.Printf("Casting: %T - Valor: %d", c, c)
}
