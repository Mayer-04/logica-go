package main

import "fmt"

/*
Dada una cadena de caracteres `s`, queremos calcular su puntuación.

* La puntuación de una cadena se calcula de la siguiente manera:
1. Primero, conviertes cada carácter de la cadena a su valor ASCII
(un número que representa a cada letra en la computadora).
2. Luego, para cada par de caracteres consecutivos en la cadena, calculas la diferencia absoluta
entre sus valores ASCII.
3. Finalmente, sumas todas esas diferencias absolutas para obtener la puntuación total.

* Reglas:
- Si la cadena tiene solo un carácter, la puntuación es 0, ya que no hay ningún par de caracteres
adyacentes para comparar.
- La diferencia absoluta entre dos números es siempre positiva o cero.
Se calcula como: |a - b|, es decir, si a es mayor o menor que b, el resultado siempre es positivo.


* Entrada:
s = "hello"

Explicación:
Los valores ASCII de los caracteres en `s` son:
'h' = 104, 'e' = 101, 'l' = 108, 'l' = 108, 'o' = 111.
Entonces, la puntuación de `s` sería:
|104 - 101| + |101 - 108| + |108 - 108| + |108 - 111| = 3 + 7 + 0 + 3 = 13.

Salida: 13

* Ejemplo 2:
Entrada:
s = "zaz"

Explicación:
Los valores ASCII de los caracteres en `s` son:
'z' = 122, 'a' = 97. Entonces, la puntuación de `s` sería:
|122 - 97| + |97 - 122| = 25 + 25 = 50.

Salida: 50
*/

func main() {
	s := "hello"
	fmt.Println(scoreOfString(s)) // 13

	//* Ejemplo 2
	str := "zaz"
	fmt.Println(scoreOfString(str)) // 50
}

func scoreOfString(s string) int {
	sum := 0
	// En cada iteración del bucle se compara el carácter actual con el carácter anterior.
	// El bucle no puede empezar en i = 0 porque entonces no habría un "carácter anterior" a comparar.
	for i := 1; i < len(s); i++ {
		sum += abs(int(s[i-1]), int(s[i]))
	}
	return sum
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
