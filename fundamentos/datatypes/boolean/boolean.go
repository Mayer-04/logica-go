package main

import "fmt"

/*
* Booleanos en Go
Un tipo de dato booleano solo puede tener dos valores: `true` (verdadero) o `false` (falso).

¿Para qué se usan?
- Representan condiciones lógicas, como "¿Está activo?" o "¿Es mayor de edad?".
- Se usan comúnmente para controlar el flujo del programa en estructuras como:
  - if (condicionales)
  - switch (selección múltiple)
  - for (bucles)
- También se utilizan en operaciones lógicas:
  - `&&` (AND): verdadero si ambas condiciones son verdaderas.
  - `||` (OR): verdadero si al menos una condición es verdadera.
  - `!` (NOT): invierte el valor (true pasa a false y viceversa).

* Características del tipo `bool` en Go:
- Es estricto: No puedes convertir un `bool` a otro tipo como entero o cadena (ni viceversa).
- Go NO hace conversiones automáticas, lo que evita errores inesperados.
- Aunque un booleano solo necesita 1 bit, Go reserva 1 byte (8 bits) por eficiencia,
ya que los procesadores acceden a la memoria por bytes.
*/

func main() {
	// Declarando una variable booleana explícitamente con valor `true`.
	// Si no lo inicializas, su valor por defecto será `false`.
	var active bool = true

	// Imprimimos el valor booleano usando %t (para "true" o "false").
	fmt.Printf("Valor: %t\n", active)

	// Otra variable booleana, ahora con valor false.
	var active2 bool = false

	// Mostramos el tipo y el valor de la variable.
	fmt.Printf("Tipo: %T - Valor: %t\n", active2, active2)

	// También puedes declarar una variable booleana sin indicar el tipo explícitamente.
	// Go infiere automáticamente que es de tipo bool.
	var active3 = true
	fmt.Println("Valor inferido:", active3)
}
