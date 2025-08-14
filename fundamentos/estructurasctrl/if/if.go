package main

import "fmt"

/*
* Estructura if en Go: Si

La estructura `if` se utiliza para ejecutar un bloque de código solo si una condición es verdadera.

* Características clave:
- La condición debe ser una expresión booleana (`true` o `false`), no se permiten conversiones implícitas.
  Por ejemplo, `if 1` o `if ""` no son válidos.
- No se usan paréntesis alrededor de la condición.
- No existe el operador ternario (`?:`) como en otros lenguajes.
- Las asignaciones dentro de la condición no están permitidas; hacerlo genera un error de compilación.
*/

func main() {
	// Condición simple: se ejecuta si la condición es verdadera.
	if true {
		fmt.Println("¡Hola Mundo!")
	}

	edad := 18

	// Estructura `if-else`: permite ejecutar un bloque si la condición es verdadera,
	// y otro bloque si es falsa.
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// If con declaración corta de variable.
	// Se puede declarar una variable en la misma línea de la condición.
	// La variable solo existe dentro del bloque `if-else`.
	// Se suele utilizar más para errores.
	if mayor := edad >= 18; mayor {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// Estructura if-else if-else.
	// Permite evaluar múltiples condiciones en secuencia (una detrás de la otra).
	// La primera condición verdadera será la que determine el bloque de código a ejecutar.
	// Si ninguna condición es verdadera, se ejecutará el bloque `else` final (si existe).
	if edad < 13 {
		fmt.Println("Eres un niño")
	} else if edad < 18 {
		fmt.Println("Eres un adolescente")
	} else {
		fmt.Println("Eres un adulto")
	}

	// * Aserción de tipo (Type Assertion).
	// Es una operación que permite obtener el `valor oculto` almacenado en una variable de tipo interface.
	// Es como decir: "Oye, sé que esta variable guarda algo, y creo que es un tipo específico,
	// quiero que me lo muestres como ese tipo".
	var i any = 10
	if value, ok := i.(int); ok {
		fmt.Println("i es un int:", value)
	} else {
		fmt.Println("i no es un int")
	}
}
