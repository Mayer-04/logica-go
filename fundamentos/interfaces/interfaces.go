package main

import "fmt"

/*
* Interfaces en Go:
Las interfaces en Go son una forma de definir un conjunto de métodos que un tipo debe implementar.
Su objetivo principal es crear `abstracciones` en el código, solo debes crear interfaces cuando realmente las necesitas.

IMPORTANTE: Go utiliza un sistema de interfaces implícitas, lo que significa que no es necesario declarar
que un tipo "implementa" una interfaz; simplemente,
si un tipo posee `todos` los métodos definidos por una interfaz, entonces ese tipo satisface esa interfaz.

- Las interfaces pueden ser implementadas por cualquier tipo.
- Un tipo puede implementar múltiples interfaces.
- Cualquier struct que implemente un método Error() automáticamente implementa la `interfaz Error`.
- En Go, es común que las interfaces tengan un solo método.
Esto se conoce como interfaces mínimas, lo que permite una mayor flexibilidad y reusabilidad.
- Es una buena práctica nombrar las interfaces con un sufijo `-er`, como `Printer` en lugar de `Print`.
- Las interfaces pueden ser compuestas, es decir, una interfaz puede incluir otras interfaces.
- IMPORTANTE: Puedes implementar interfaces definidas en otros paquetes de Go, no solo las que tú defines.

* Características adicionales:
- Una interfaz nula no tiene un valor ni un tipo concreto.
Una interfaz nula se crea simplemente al declarar una interfaz sin asignarle un valor.
- Una interfaz vacía tiene cero métodos y puede ser utilizada para manejar datos de cualquier tipo.
Una interfaz vacía no es considerada una expresión.

* ¿Cuándo usar interfaces?
- Comportamiento común: Usar interfaces cuando varios tipos implementan un comportamiento común (Polimorfismo).
*Ejemplo:
Imagina que tienes varios tipos (estructuras) en tu programa que comparten un comportamiento común.
En lugar de duplicar código, puedes definir una interface que capture ese comportamiento.

- Desacoplamiento: El desacoplamiento es cuando queremos que nuestro código sea más flexible y no esté atado
a implementaciones específicas, es decir, que cada componente sea independiente y pueda cambiarse sin afectar a otros componentes.
* Ejemplo:
Si estas creando una aplicación que necesita guardar datos e inicialmente decides guardar los datos
en un archivo y deseas cambiar la implementación para que los datos se guarden en una base de datos,
puedes definir un interface que te permita hacer estos cambios de manera sencilla.

- Restringir el comportamiento: Cuando queremos asegurarnos de que un tipo específico cumple
con ciertos requisitos o restricciones.
*/

// Printer es una interfaz mínima que define un solo método Print.
type Printer interface {
	Print()
}

// Logger es una interfaz que define un solo método Log.
type Logger interface {
	Log(message string)
}

// PrinterLogger es una interfaz compuesta que incluye las interfaces Printer y Logger.
// Cualquier tipo que implemente ambos métodos (Print y Log) satisfará la interfaz PrinterLogger.
type PrinterLogger interface {
	Printer
	Logger
}

// MyError es una estructura personalizada que implementara la interfaz `error`.
// Contiene un campo `Msg` que almacenará un mensaje de error.
type MyError struct {
	Msg string
}

// Error es un método que hace que MyError cumpla con la interfaz `error`.
// La interfaz `error` es una interfaz nativa de Go que solo requiere la implementación
// de un método: `Error() string`. Cualquier tipo que implemente este método puede
// ser considerado un error en Go.
// El método Error de MyError devuelve el mensaje de error almacenado en el campo `Msg`.
func (e MyError) Error() string {
	return e.Msg
}

// Person es una estructura que contiene el nombre y la edad de una persona.
type Person struct {
	Name string
	Age  int
}

// Print es un método de Person que cumple con la interfaz `Printer`.
// Esto significa que cualquier instancia de Person puede ser asignada a una variable de tipo Printer.
func (p Person) Print() {
	fmt.Printf("Hola, mi nombre es %s y tengo %d años.\n", p.Name, p.Age)
}

// Log es un método de Person que cumple con la interfaz Logger.
// Esto significa que cualquier instancia de Person puede ser asignada a una variable de tipo Logger.
func (p Person) Log(message string) {
	fmt.Printf("[LOG]: %s\n", message)
}

func main() {
	// Crear una instancia de Person y asignarla a una variable de tipo Printer.
	// Esto es posible porque Person implementa el método Print, que es el único
	// requisito para satisfacer la interfaz Printer.
	var p Printer = Person{
		Name: "Mayer",
		Age:  24,
	}
	p.Print() // Hola, mi nombre es Mayer y tengo 24 años.

	// Crear otra instancia de Person y usar el método Print directamente.
	// Aquí no estamos usando la interfaz, sino la estructura directamente.
	// El método `Print` de Person ya implementa implícitamente la interfaz Printer.
	robert := Person{
		Name: "Robert",
		Age:  59,
	}
	robert.Print() // Hola, mi nombre es Robert y tengo 59 años.

	// Crear una instancia de Person y asignarla a una variable de tipo Logger.
	// Esto es posible porque Person implementa el método Log.
	var cristiano Logger = Person{
		Name: "Cristiano",
		Age:  39,
	}
	cristiano.Log("Soy el bicho...SIUUU!") // [LOG]: Soy el bicho...SIUUU!

	// Crear una instancia de Person y asignarla a una variable de tipo PrinterLogger.
	// Esto es posible porque Person implementa ambos métodos: Print y Log,
	// cumpliendo así con la interfaz compuesta PrinterLogger.
	var person PrinterLogger = Person{
		Name: "Messi",
		Age:  37,
	}
	person.Print()                 // Hola, mi nombre es Messi y tengo 37 años.
	person.Log("Hola, soy Messi!") // [LOG]: Hola, soy Messi!

	//* Creamos una instancia de MyError con un mensaje de error.
	err := MyError{Msg: "Ocurrió un error"}

	// Dado que MyError implementa el método Error, se puede asignar a una variable de tipo `error`.
	var genericError error = err

	// Imprimimos el mensaje de error.
	fmt.Println("Error personalizado:", genericError.Error()) // Ocurrió un error

	// Ejemplo de interfaz nula.
	// No ha sido inicializada con ningún valor. Debido a esto, su valor es nil.
	var i interface{}
	fmt.Println("Interfaz nula:", i) // Interfaz nula: <nil>

	// Esta interfaz vacía contiene en su interior un valor de tipo string.
	e := interface{}("Go")
	fmt.Println("Interfaz vacía con un valor de tipo string:", e) // Output: Go

	// Ejemplo de interfaz vacía.
	var emptyInterface interface{}
	emptyInterface = 42
	fmt.Println("Interfaz vacía con un int:", emptyInterface) // Interfaz vacía con un int: 42

	emptyInterface = "Hello, World!"
	fmt.Println("Interfaz vacía con un string:", emptyInterface) // Interfaz vacía con un string: Hello, World!

	// Type assertion - Verificar el tipo de emptyInterface.
	// Validar el tipo concreto que tiene la interfaz vacía y si es ese tipo obtener el valor que esta dentro de ella.
	// Si emptyInterface.(type)
	switch emptyInterface.(type) {
	case int:
		fmt.Println("Es un int")
	case string:
		fmt.Println("Es un string")
	default:
		fmt.Println("No es ni un int ni un string")
	}
}
