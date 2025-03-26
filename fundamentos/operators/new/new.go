package main

import "fmt"

/*
* Operador new
- La función `new(T)` se utiliza para crear una nueva instancia de un tipo `T`, devolviendo un puntero al tipo `T`.
- Esto es útil cuando quieres inicializar una variable y trabajar directamente con su dirección en memoria.
- Cuando se busca maximizar el rendimiento se recomiendautilizar new(T) en vez de &T{}.
- Se usa principalmente con tipos no compuestos como int, float64 o bool.
- El operador `&` (ampersand) es más común cuando se trabaja con tipos compuestos como structs.

* IMPORTANTE
Si el puntero creado con `new()` solo se usa dentro de una función y no "escapa" a otras partes del programa,
Go puede asignarlo en la pila (stack) en lugar del heap.
- Esto hace que la operación sea más rápida y eficiente, ya que la memoria en la pila se libera
automáticamente cuando la función termina.
*/

type Person struct {
	Name     string
	LastName string
	Age      int
	Hoobies  []string
}

func main() {
	// Uso de la función new para crear un puntero a un entero.
	var x = new(int)

	// La dirección de memoria almacenada en x.
	fmt.Println("Dirección de memoria de x:", x) // Output: Dirección de memoria de x: 0xc00000a0d8

	// Asignar un valor a la dirección de memoria apuntada por x.
	*x = 10

	//* Ejemplo 2
	var a *int
	// Asigna memoria para un entero e inicializa el valor con 0 (valor cero por defecto de int).
	var b *int = new(int)

	fmt.Println(a)  // nil
	fmt.Println(b)  // 0xc00000a120
	fmt.Println(*b) // 0 (valor por defecto de int)

	// Acceder al valor almacenado en la dirección de memoria usando el operador de desreferenciación (*).
	fmt.Println("Nuevo valor de x:", *x) // Output: Nuevo valor de x: 10

	//* Crear una nueva instancia de Person usando new.
	// Go inicializa automáticamente los campos del struct con los valores cero de su tipo de dato correspondiente.
	var person = new(Person)

	// Imprimir los valores por defecto de la estructura Person.
	fmt.Println("Valores por defecto de person:", person) // Output: &{"" "" 0 []}

	// Asignar valores a los campos del struct.
	person.Name = "Mayer"
	person.LastName = "Chaves"
	person.Age = 24
	person.Hoobies = []string{"Programar", "Leer"}

	// Imprimir el contenido del struct utilizando el operador de desreferenciación (*).
	fmt.Println("Valores actualizados de person:", *person) // Output: {Mayer Chaves 24 [Programar Leer]}

	// Comparación de punteros creados con new.
	otraPersona := new(Person)
	// Verificar si dos punteros creados con new apuntan a diferentes ubicaciones de memoria.
	fmt.Println("¿Apuntan a la misma dirección de memoria?", person == otraPersona) // Output: false

	// Uso de new() con tipos básicos como `float64`.
	var pi = new(float64)
	*pi = 3.14159
	// Output: Valor de pi: 3.14159, Dirección de memoria: 0xc00000c0a8
	fmt.Printf("Valor de pi: %.5f, Dirección de memoria: %p\n", *pi, pi)
}
