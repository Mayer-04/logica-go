package main

import "fmt"

/*
* Operadores aritméticos
- Si divides dos enteros (int), Go realiza la división entera.
Esto significa que el resultado siempre será un `número entero`, descartando la parte decimal (truncamiento).

* Ejemplo:
result := 3 / 6
Como 3 y 6 son enteros, Go realizará la división entera, y el resultado será 0 en lugar de 0.5.

- Si quieres obtener un valor con decimales (punto flotante),
necesitas que al menos uno de los operandos sea un número de punto flotante (float32 o float64).

* Orden de precedencia o prioridad de operadores:
1. Paréntesis `()`.
2. Multiplicación `*`, división `/` y resto `%`: Tienen la misma prioridad y se evalúan de izquierda a derecha.
3. Suma `+` y resta `-`: Tienen la misma prioridad y se evalúan de izquierda a derecha.
*/

func main() {
	// Suma de enteros.
	suma := 10 + 10
	fmt.Println("Suma:", suma) // Output: 20

	// Resta de enteros.
	resta := 10 - 5
	fmt.Println("Resta:", resta) // Output: 5

	// Multiplicación de enteros.
	multiplicacion := 10 * 5
	fmt.Println("Multiplicación:", multiplicacion) // Output: 50

	// División de enteros.
	// Como la división es de enteros no devuelve la parte decimal.
	division := 1905 / 100
	fmt.Println("División:", division) // Output: 19

	// División con parte decimal (flotantes).
	// Al menos uno de los operandos debe ser de punto flotante.
	division2 := float64(1905) / 100
	fmt.Println("División:", division2) // Output: 19.05

	// Módulo o Resto de enteros.
	// Devuelve el resto de la división.
	modulo := 10 % 5
	fmt.Println("Modulo:", modulo) // Output: 0

	//* Ejemplo del orden de ejecución de las operaciones.
	// Primero se realizan las operaciones que estan entre parentesis y luego se multiplica por 2.
	ejemplo := (2 + 3) * 2
	fmt.Println("Ejemplo:", ejemplo) // Output: 10
}
