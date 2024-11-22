package main

import (
	"fmt"
	"math/rand/v2"
)

/*

nuevo paquete math/rand/v2 en Go 1.22!

- math/rand/v2: Introduce nuevos generadores, como PCG y ChaCha8,
que son más eficientes y producen una mejor calidad de números aleatorios.
- El paquete math/rand/v2 ha sido diseñado para solucionar las limitaciones de math/rand,
proporcionando mejoras en eficiencia, calidad de aleatoriedad y escalabilidad, especialmente en contextos de concurrencia.

* rand.IntN(n): Devuelve un número entero aleatorio entre 0 y 𝑛−1.
n−1 (inclusive) de tipo int. Es útil para la mayoría de los casos donde necesitas un número entero.

rand.IntN(n) genera un número pseudoaleatorio en el intervalo medio-abierto [0,𝑛)[0,n).
Esto significa que el número generado será un entero que puede ser 0, pero será menor que 𝑛.
La función lanza un pánico si el valor de 𝑛 es menor o igual a 0.



* rand.N(n): Devuelve un número aleatorio en el mismo rango,
pero te permite especificar cualquier tipo de entero, brindando más flexibilidad.

rand.N(n) también genera un número pseudoaleatorio en el intervalo medio-abierto [0,𝑛)[0,n),
pero a diferencia de rand.IntN, puede devolver un número en cualquier tipo de entero (por ejemplo, int, int32, int64).
Esto permite más flexibilidad en cuanto al tipo de valor que quieres recibir.
*/

func main() {
	// Generar un número aleatorio entre 1 y 50
	num1 := rand.IntN(50) + 1 // De 1 a 50
	fmt.Println("Número aleatorio entre 1 y 50:", num1)
	// Generar un número aleatorio entre 0 y 60
	num2 := rand.IntN(61) // De 0 a 60
	fmt.Println("Número aleatorio entre 0 y 60:", num2)

	// rand.N(-100)
}
