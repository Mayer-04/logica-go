package main

import (
	"fmt"
	"strings"
)

/*
* Funciones en Go
- Go soporta aspectos de programación funcional, tratando funciones como "ciudadanos de primera clase".
- Las funciones pueden ser asignadas a variables, pasadas como argumentos y retornadas desde otras funciones.
- IMPORTANTE: Los parámetros de funciones en Go son pasados por valor, no por referencia.
Se trabaja con una copia del valor, no con el valor original.
- Para tipos como int, float, bool, string, array y struct, pasar por valor significa copiar todo el contenido.
- Los tipos que parecen pasarse por referencia (slices, maps, channels, etc.) en realidad se pasan por valor,
pero el valor que se pasa contiene una referencia a los datos subyacentes.
- Para slices, maps, channels, y pointers, pasar por valor significa copiar la referencia, no los datos a los que apunta.

* Conceptos clave:
1. Funciones de orden superior: Funciones que aceptan otras funciones como argumentos o las retornan.
2. Closures: Funciones anónimas que pueden acceder y modificar variables de su ámbito externo.
3. Funciones variádicas: Aceptan un número variable de argumentos.
4. Funciones diferidas (defer): Se ejecutan al final de la función que las contiene.
5. Funciones anónimas: Funciones sin nombre, útiles para definiciones inmediatas.
6. Recursión: Funciones que se llaman a sí mismas.
*/

func main() {
	// Llamada a la función saludar con dos argumentos de tipo string.
	saludar("Mayer", "Chaves")

	// Ejemplo con función por referencia.
	num := 5
	increment(&num)  // Obtenemos la dirección en memoria de la variable "num".
	fmt.Println(num) // El valor de num se cambia directamente

	//* RETURN: Llamada a una función con retorno de un solo valor.
	result := retornar(1, 3)
	fmt.Println(result)

	//* Múltiples valores a retornar.
	// Podemos ignorar u omitir algún valor utilizando el operador blank (_).
	lower, upper := multiplesRetornos("Mayer")
	fmt.Println(lower, upper)

	//* Función variádica: Acepta un número indefinido de argumentos.
	fmt.Println(sumaVariadica(1, 2, 3, 4, 5))

	//* Llamando a Función anonima.
	anonima()

	//* Creando una función anonima.
	// Puede ser definida en el lugar donde se necesite en lugar de ser definida globalmente.
	salud := func(name string) {
		fmt.Printf("Hola, %s!\n", name)
	}
	salud("Robert")

	//* Función anonima autoinvocada.
	func(name string) {
		fmt.Println("Hello", name)
	}("mayer")

	//* Segunda manera de declarar una función anonima con la declaración de variable corta.
	ejemplo := func() {
		fmt.Println("Soy una función anonima corta")
	}
	ejemplo()

	//* Uso de funciones diferidas (defer).
	// La función diferida se ejecutará al final de la función main.
	defer diferida()

	//* Llamada a una función recursiva.
	fmt.Println("Factorial de 5:", factorial(5))

	//* Closure: función que captura variables del entorno.
	counter := createCounter()
	fmt.Println("Counter:", counter()) // Output: Counter: 1
	fmt.Println("Counter:", counter()) // Output: Counter: 2

	//* Uso de un tipo (type) de función.
	var f tipoFuncion = saludar
	f("Mayer", "Andres")
}

//* Características de las funciones.

// Función con parámetros agrupados del mismo tipo.
func saludar(name, lastname string) {
	fmt.Printf("Nombre %s con apellido %s", name, lastname)
}

// Función que trabaja con un puntero.
// Para trabajar con parámetros por referencia se deben utilizar "Punteros".
func increment(x *int) {
	// Operador desreferenciación - Acceder al valor almacenado en la dirección de memoria.
	*x++ // Incrementa el valor en la dirección de memoria apuntada por x
}

// Función con valor de retorno - return.
func retornar(num1, num2 int) int {
	return num1 + num2
}

// * Funciones con múltiples retornos.
// * Se debe incluir entre paréntesis los múltiples valores a retornar.
// Es posible nombrar los parámetros de retorno, aunque no es recomendado en funciones complejas.
func multiplesRetornos(text string) (lower, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)
	return // Retorno implícito
}

// * Funciones variádicas - Permiten pasar un número indefinido de argumentos del mismo tipo.
// Se anteponen tres puntos (...) al tipo de dato para indicar que es una función variádica.
func sumaVariadica(nums ...int) int {
	//El parámetro nums se convierte en un "slice" de enteros.
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// * Funciones anónimas - No tienen un nombre de función.
// Las funciones anónimas se pueden asignar a variables y ser invocadas posteriormente.
var anonima = func() {
	fmt.Println("Soy una función anonima")
}

// * Funciones diferidas - defer().
// Las funciones defer se ejecutan al final de la función que las contiene.
func diferida() {
	fmt.Println("Esta función se ejecutará al final de main")
}

// * Función recursiva - Una función que se llama a sí misma.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// * Closures - Funciones que capturan variables del entorno.
// Cada vez que la función anónima es llamada, accede y modifica la misma instancia de "counter", que existía cuando fue creada.
func createCounter() func() int {
	counter := 0
	// La función anónima captura la referencia a la variable counter de su entorno léxico osea createCounter().
	return func() int {
		counter++ // La función anónima captura y modifica la variable counter.
		return counter
	}
}

// * Definición de un tipo de función.
// Suele usarse en funciones complejas.
type tipoFuncion func(string, string)
