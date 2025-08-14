package main

import (
	"cmp"
	"fmt"
)

/*
* Generics: Genéricos
Es una característica que nos permite definir funciones, tipos y estructuras que acepten diferentes tipos de datos.

- Introducido en la versión 1.18 de Go.
- La idea de los genéricos es no sacrificar la seguridad que brindan los tipos.
- tipo any: Es un alias para `interface{}`, que permite trabajar con cualquier tipo.
- Permite escribir código con tipos que se pueden especificar más tarde e instanciar cuando sea necesario.
- Go puede deducir automáticamente el tipo de los genéricos a partir de los argumentos en la llamada a la función.

* Restricción:
Una restricción es una `interfaz` que define los tipos permitidos para los parámetros genéricos:
- Puede restringir los tipos permitidos a un conjunto de comportamientos (métodos) o tipos específicos.
- Utiliza la notación `~` para permitir tipos subyacentes equivalentes, es decir, variantes de un tipo base.
Cualquier tipo que creemos y se base en un tipo existente.

* Características de los genéricos en Go:
- [T any]: Permite trabajar con cualquier tipo.
- comparable: Es una restricción que asegura que el tipo puede compararse con == o !=.
- Unión de tipos |: Permiten limitar los tipos válidos a un grupo específico.
- Operador ~ (tipo subyacente): Indica que se acepta cualquier tipo cuyo tipo base sea el especificado.
*/

// * Función genérica.
// T: Es un tipo de parámetro que permite a la función aceptar cualquier tipo de dato.
// value: Es el valor que se va a imprimir, que puede ser de cualquier tipo.
// La función `Print` puede imprimir valores de cualquier tipo.
func Print[T any](value T) {
	fmt.Println(value)
}

// * Tipo genérico - Estructura genérica.
// El operador | indica que el parámetro T puede ser de tipo `int` o de tipo `uint`.
// El campo `Age` puede ser de tipo `int` o de tipo `uint` dependiendo de la instancia.
type Person[T int | uint] struct {
	Name string
	Age  T // int o uint
}

// * Función genérica que suma números.
// Esta función genérica acepta cualquier tipo numérico (int, float, etc.),
// utilizando una restricción más amplia para los tipos numéricos.
func Sum[T int | float64](a, b T) T {
	return a + b
}

// * Definición de una restricción.
// La restricción `Constraint` limita el tipo a `int` o `string`,
// lo que permite crear funciones genéricas que solo acepten estos tipos.
type Constraint interface {
	~int | ~string
}

// * Función genérica que usa la restricción `Constraint`.
// Esta función solo acepta valores de tipos permitidos por la restricción (`int`, `string` o sus variantes).
func PrintConstraint[T Constraint](value T) {
	fmt.Printf("Valor: %v (Tipo: %T)\n", value, value)
}

// Esta función genérica acepta cualquier tipo.
// Puedes usarlo para declarar variables, crear punteros, o construir slices,
// pero no puedes asumir que el tipo soporta operaciones específicas (como == o >).
func AnyUsage[T any](t T) {
	a := t
	b := new(T)
	c := &t
	d := []T{t}

	fmt.Printf("a: %v, b: %v, c: %v, d: %v\n", a, b, c, d)
}

// * Función genérica que compara dos valores.
// Utiliza la restricción `comparable` para asegurar que los valores puedan ser comparados.
// La restricción `comparable` es útil con los mapas, donde las `claves` deben ser de tipo `comparable`.
func Compare[T comparable](a, b T) bool {
	return a == b
}

// Usa la restricción `constraints.Ordered` para comparar valores ordenables (números, strings).
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// PrintSlice imprime los elementos de un slice de cualquier tipo.
// Mostrando el índice y el valor.
func PrintSlice[T any](items []T) {
	for i, v := range items {
		fmt.Printf("[%d]: %v\n", i, v)
	}
}

func main() {
	// Múltiples llamadas a la función genérica `Print` con diferentes tipos.
	Print(10)     // int
	Print("Hola") // string
	Print(true)   // bool

	// AnyUsage es una función genérica que acepta cualquier tipo.
	AnyUsage(10)     // int
	AnyUsage("Hola") // string
	AnyUsage(true)   // bool

	// Instancias de la estructura genérica `Person` con diferentes tipos de `Age`.
	person1 := Person[uint]{Name: "Mayer", Age: 24}
	fmt.Printf("%+v\n", person1)
	person2 := Person[int]{Name: "Mayer", Age: -24}
	fmt.Printf("%+v\n", person2)

	// Uso de la nueva función genérica `Sum` para sumar valores numéricos.
	fmt.Println(Sum(10, 20))   // Suma de enteros
	fmt.Println(Sum(5.5, 4.5)) // Suma de flotantes

	// Llamadas a la función genérica `PrintConstraint` con tipos restringidos.
	PrintConstraint(42)          // Valor: 42 (Tipo: int)
	PrintConstraint("Genéricos") // Valor: Genéricos (Tipo: string)

	// Comparación de valores con la función `Compare`.
	fmt.Println(Compare(10, 10))             // true
	fmt.Println(Compare("Go", "Go"))         // true
	fmt.Println(Compare("Go", "JavaScript")) // false

	// Uso de la función genérica `Max` para obtener el mayor de dos valores.
	fmt.Println("Mayor:", Max(10, 20))
	fmt.Println("Mayor:", Max("Go", "Java"))

	// Uso de la función genérica `PrintSlice` para imprimir un slice de cualquier tipo.
	numbers := []int{1, 2, 3, 4}
	PrintSlice(numbers)

	words := []string{"hola", "genéricos", "go"}
	PrintSlice(words)
}
