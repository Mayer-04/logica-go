package main

import "fmt"

/*
* Composición en Go
En Go, no existe la herencia como en otros lenguajes orientados a objetos.
En su lugar, se utiliza la composición para reutilizar y extender funcionalidades.

* Composición: Se logra a través de la `incrustación` de tipos.

* Incrustación de tipos.
Consiste en incluir una estructura dentro de otra estructura.
Esto permite que la estructura exterior "adopte" automáticamente todos los campos y métodos
de la estructura interior.

- Ventajas sobre la herencia tradicional:
  • La composición ayuda a evitar la jerarquía de clases (niveles) que se suele ocurrir con la herencia.
  • Hace que el código sea más fácil de entender y mantener.
  • Facilita la creación de pruebas unitarias al trabajar con piezas más pequeñas e independientes.

- La estructura exterior puede acceder directamente a los campos y métodos de la estructura
incrustada como si fueran propios, lo que hace el código más limpio y legible.
*/

// Definimos una estructura `Engine` que representa el motor de un vehículo.
type Engine struct {
	Horsepower int // Caballos de fuerza del motor.
}

// Método `Start` asociado a la estructura `Engine` para simular el arranque del motor.
func (e Engine) Start() {
	fmt.Println("Engine started")
}

// La estructura `Car` incrusta `Engine`, lo que significa que `Car` 'hereda' los campos y métodos de `Engine`.
// Además, `Car` tiene su propio campo `Model` que representa el modelo del vehículo.
type Car struct {
	Engine        // Incrustación del tipo `Engine`, lo que permite que `Car` tenga acceso a `Horsepower` y `Start()`.
	Model  string // Modelo del vehículo
}

func main() {
	// Creamos una instancia de `Car`, inicializando el campo incrustado `Engine` y el campo `Model`.
	// Se puede acceder directamente a los campos y métodos de `Engine` desde una instancia de `Car`.
	myCar := Car{
		Engine: Engine{Horsepower: 200},
		Model:  "Toyota",
	}

	// Accedemos directamente al campo `Model` y al campo incrustado `Horsepower`.
	fmt.Println("Model:", myCar.Model) // Output: Model: Toyota
	// Esto es como si fueran campos propios de `Car`.
	fmt.Println("Horsepower:", myCar.Horsepower) // Output: Horsepower: 200

	// Llamamos al método `Start` que `Car` "hereda" de `Engine` a través de la composición.
	myCar.Start() // Output: Engine started
}
