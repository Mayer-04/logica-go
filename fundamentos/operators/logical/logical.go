package main

import "fmt"

/*
 * Operadores lógicos en Go
---------------------------
En Go, los operadores lógicos funcionan exclusivamente con valores booleanos (bool).
Son utilizados para evaluar expresiones lógicas y controlar el flujo de ejecución en programas.
- && (AND lógico): Devuelve true si ambos operandos son true.
- || (OR lógico): Devuelve true si al menos uno de los operandos es true.
- !  (NOT lógico): Invierte el valor de un booleano (true → false, false → true).
*/

func main() {
	a := true
	b := false
	c := true

	fmt.Println(a && b) // false
	fmt.Println(a || b) // true
	fmt.Println(!a)     // false

	// Múltiples expresiones
	fmt.Println("(a && c) || b:", (a && c) || b)   // true
	fmt.Println("!(a && b) || c:", !(a && b) || c) // true

	// Ejemplos de uso en condicionales
	if a && c {
		fmt.Println("Ambos valores son verdaderos")
	}

	if !b {
		fmt.Println("b es falso, por lo que !b es verdadero")
	}

	// Ejemplos de uso en ciclos
	for a || c {
		fmt.Println("a es verdadero o c es verdadero, por lo que se ejecuta el ciclo")
		break
	}

	//* EJEMPLOS DE USO REAL

	// Ejemplo 1: Validación de credenciales - && AND lógico
	user := "admin"
	password := "12345"

	if user == "admin" && password == "12345" {
		fmt.Println("Acceso concedido")
	} else {
		fmt.Println("Acceso denegado")
	}

	// Ejemplo 2: Validación de respuestas - || OR lógico
	response := "sí"

	if response == "sí" || response == "Si" || response == "SI" {
		fmt.Println("Respuesta afirmativa")
	} else {
		fmt.Println("Respuesta negativa")
	}

	// Ejemplo 3: Validación de luz roja - ! NOT lógico
	luzRoja := true

	if !luzRoja {
		fmt.Println("Puedes avanzar")
	} else {
		fmt.Println("Detente, la luz está en rojo")
	}
}
