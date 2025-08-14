package main

import "fmt"

/*
 * Operadores lógicos en Go
---------------------------
Los operadores lógicos se utilizan para trabajar con valores booleanos (true o false).
Permiten combinar condiciones y controlar el flujo del programa (por ejemplo, en `if` y `for`).

Estos son los tres operadores lógicos principales:

- && (AND lógico): Devuelve `true` solo si ambas condiciones son verdaderas.
- || (OR lógico): Devuelve `true` si al menos una de las condiciones es verdadera.
- !  (NOT lógico): Invierte el valor booleano. !true es false y !false es true.
*/

func main() {
	a := true
	b := false
	c := true

	// Operaciones básicas con booleanos.
	fmt.Println(a && b) // false, porque b es false
	fmt.Println(a || b) // true, porque a es true
	fmt.Println(!a)     // false, porque a es true y el operador ! lo invierte

	// Evaluación de expresiones compuestas.
	// (true && true) es true, y true || false es true.
	fmt.Println("(a && c) || b:", (a && c) || b) // Output: true
	// (a && b) es false, !false es true, y true || true es true.
	fmt.Println("!(a && b) || c:", !(a && b) || c) // Output: true

	// Uso en estructuras condicionales.
	if a && c {
		fmt.Println("Ambos valores son verdaderos")
	}

	if !b {
		fmt.Println("b es falso, por lo tanto !b es verdadero")
	}

	// Uso en ciclos (bucles).
	for a || c {
		fmt.Println("a es verdadero o c es verdadero, por lo tanto se ejecuta el ciclo")
		break // Evitamos un ciclo infinito
	}

	//* EJEMPLOS DE USO REAL EN SITUACIONES COMUNES.

	// EJEMPLO 1: Validación de credenciales usando AND lógico.
	user := "admin"
	password := "12345"

	if user == "admin" && password == "12345" {
		fmt.Println("Acceso concedido") // Ambas condiciones son verdaderas
	} else {
		fmt.Println("Acceso denegado")
	}

	// EJEMPLO 2: Reconocimiento de respuestas afirmativas usando OR lógico.
	response := "sí"

	if response == "sí" || response == "Si" || response == "SI" {
		fmt.Println("Respuesta afirmativa") // Se acepta cualquier forma de "sí"
	} else {
		fmt.Println("Respuesta negativa")
	}

	// EJEMPLO 3: Control de semáforo usando NOT lógico.
	luzRoja := true

	if !luzRoja {
		fmt.Println("Puedes avanzar") // Solo si la luz NO está en rojo
	} else {
		fmt.Println("Detente, la luz está en rojo")
	}
}
