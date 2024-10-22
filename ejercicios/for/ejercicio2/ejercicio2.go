package main

import (
	"fmt"
)

/*
* Buscar el primer número primo:
Usa un bucle anidado con break para optimizar la búsqueda de números primos.
*/

/*
el operador && solo se puede aplicar a expresiones booleanas (que devuelven true o false)
, y parece que en tu código lo estás utilizando con valores de tipo int.
*/

func main() {

	// 1. Mayor que 1
	// 2. El número no se considera ni primo ni compuesto
	// 3. Tiene solo dos divisores positivos: el mismo número y el 1
	// 4. Se entiende por divisible a que el resultado sea un número entero, es decir, no decimal.

	/*
		- Si numero % i === 0, significa que numero es divisible por i (no deja residuo).
		Esto indica que numero tiene un divisor diferente a 1 y a sí mismo,
		lo cual contradice la definición de un número primo.

		- Si encontramos un divisor, inmediatamente retornamos false, ya que esto prueba que el número no es primo
		y no hace falta seguir revisando otros divisores.

	*/

	for i := 2; i <= 50; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}

		// fmt.Println(i)
	}
}
