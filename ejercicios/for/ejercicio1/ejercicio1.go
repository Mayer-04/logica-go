package main

import "fmt"

/*
* Suma de números pares hasta que la suma supere 50.
1. Itera a través de los números del 1 al 100.
2. Suma únicamente los números pares.
3. Utilice `continue` para saltar números impares y `break` para detener el bucle cuando la suma supera 50.
4. Al finalizar, imprime el valor de la suma acumulada en la consola.
*/

func main() {
	limite := 100
	suma := 0
	for i := 1; i <= limite; i++ {
		if i%2 != 0 { // Si el número es impar
			continue // Saltamos al siguiente ciclo
		}

		suma += i // Sumamos el número par

		if suma > 50 { // Si la suma supera 50, salimos del bucle
			break
		}
	}
	fmt.Println(suma)
}
