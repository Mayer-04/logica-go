package main

import "fmt"

/*
Dado un array de enteros, necesitamos calcular las `proporciones` de elementos que son positivos, negativos y ceros.
Luego, imprimimos estas proporciones en formato decimal, con seis cifras decimales.

Proporciones: Se refiere a la fracción de elementos en el array que cumplen con una cierta condición
(ya sea ser positivo, negativo o cero), en relación con el tamaño total del array.

Supongamos que tenemos un array con 10 elementos y que 3 de esos elementos son positivos.
La proporción de números positivos sería:

Proporción positivas = número de positivos / total de elementos del array = 3 / 10 = 0.3

Paso 1: Contar cada tipo de número.
Paso 2: Calcular las proporciones.
Paso 3: Imprimir las proporciones con seis decimales.

Entrada: [-4, 3, -9, 0, 4, 1]

Salida:
0.500000
0.333333
0.166667
*/

func main() {
	elements := []int{-4, 3, -9, 0, 4, 1}

	var pos, neg, zero int
	for _, element := range elements {
		switch {
		case element == 0:
			zero++
		case element < 0:
			neg++
		case element > 0:
			pos++
		}
	}

	size := float64(len(elements))
	fmt.Printf("%.6f\n%.6f\n%.6f\n", float64(pos)/size, float64(neg)/size, float64(zero)/size)
}
