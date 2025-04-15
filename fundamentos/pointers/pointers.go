package main

import "fmt"

/*
* Pointers: Punteros
Puntero: Es una variable que almacena la dirección de memoria de otra variable.

* Declarar un puntero
Para declarar un puntero, se utiliza un asterisco (`*`) antes del tipo de dato.

* Operador de desreferenciación
Para obtener el valor del puntero se utiliza el `operador de desreferenciación` utilizando (*) antes
del nombre de la variable.

* Operador de dirección
Para obtener la dirección de memoria de una variable, se usa el operador de dirección (`&`): &number
*/

func main() {
	// Declaración e inicialización de una variable de tipo entero.
	var numero int = 42

	//* Declaración de un puntero de tipo int.
	// En este punto, el puntero no apunta a ninguna dirección de memoria específica (es nil por defecto).
	// Se recomienda despues de declarar la variable asignar su valor en una sola línea,
	// en lugar de declarar la variable primero y luego asignarle un valor en una línea separada.
	var puntero *int

	// Declaramos un variable `puntero` como un puntero a un entero (int), pero no se le asigna memoria,
	// por lo que su valor por defecto es `nil`.
	fmt.Println("Puntero:", puntero) // Output: <nil>

	//* Operador de dirección de punteros `&`.
	// Asignación del puntero a la dirección de memoria de la variable numero usando el operador &.
	puntero = &numero

	// Acceso al valor a través del puntero utilizando el operador de desreferenciación (*).
	fmt.Println("Valor de numero a través del puntero:", *puntero) // Output: 42

	// Modificación del valor almacenado en la dirección de memoria a la que apunta el puntero.
	*puntero = 10

	// El valor de la variable numero se ha actualizado a través del puntero.
	fmt.Println("Nuevo valor de numero:", numero) // Output: Nuevo valor de numero: 10

	// Ejemplo 2: Declaración de una variable string y un puntero a string.
	fruit := "🍎"
	var pointer *string = &fruit

	// Imprimir la dirección de memoria almacenada en el puntero.
	fmt.Println("Pointer:", pointer) // Output: Pointer: 0xc00008a030
	// Imprimir el valor almacenado en la dirección de memoria apuntada por el puntero.
	fmt.Println("Valor apuntado por el puntero:", *pointer) // Output: Valor apuntado por el puntero: 🍎

	// Ejemplo con la función incrementar.
	var x int = 10

	// Pasar un puntero a la función incrementar.
	incrementar(&x)

	// Imprimir el valor de x después de ser incrementado por la función.
	fmt.Printf("Valor de x después de incrementar: %d, Tipo: %T\n", x, x)

	//* Uso de punteros con estructuras (struct).
	type persona struct {
		nombre string
		edad   int
	}

	// Declarar una variable de tipo Persona.
	carlos := persona{nombre: "Carlos", edad: 25}

	// Crear un puntero a la estructura Persona.
	punteroPersona := &carlos

	// Modificar el valor del campo edad a través del puntero.
	punteroPersona.edad = 26

	// Imprimir los valores actualizados de la estructura.
	fmt.Printf("Nombre: %q, Edad: %d\n", carlos.nombre, carlos.edad) // Output: Nombre: Carlos, Edad: 26
}

// La función `incrementar` toma un puntero a un entero y aumenta el valor apuntado en 1.
// Es un ejemplo de cómo modificar una variable fuera de su contexto usando un puntero.
func incrementar(puntero *int) {
	// El operador de desreferenciación * permite acceder al valor al que apunta el puntero.
	*puntero++
}
