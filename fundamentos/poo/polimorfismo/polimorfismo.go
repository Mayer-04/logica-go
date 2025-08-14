package main

import "fmt"

/*
* Polimorfismo en Go
El polimorfismo es la capacidad de un mismo código para trabajar con diferentes tipos de datos.
En Go, esto se consigue mediante el uso de interfaces, que son como "contratos" que definen
qué acciones debe poder realizar un tipo.

* Características del polimorfismo en Go:
- Una interfaz define un conjunto de métodos que algo debe poder hacer (como "nadar" o "calcular").
  No dice cómo hacerlo, solo que debe poder hacerse.

- Cualquier tipo que tenga todos los métodos definidos en una interfaz automáticamente
  "cumple el contrato" de esa interfaz, sin necesidad de declararlo explícitamente.
  Por ejemplo, si una interfaz "Nadador" requiere un método "Nadar()", cualquier tipo
  con ese método puede ser tratado como un "Nadador".

- Esto permite escribir funciones que acepten cualquier tipo que sepa hacer ciertas acciones,
  sin importar qué tipo específico sea.

- Lo mejor es que diferentes tipos pueden implementar los mismos métodos de formas distintas,
  pero seguir siendo compatibles con la misma interfaz. Por ejemplo, un "Pato" y un "Humano"
  pueden "Nadar()" de maneras diferentes, pero ambos pueden ser tratados como "Nadadores".
*/

// Definimos una interfaz 'Vehicle' que tiene un método 'Start'.
// Cualquier tipo que implemente este método, automáticamente satisface la interfaz 'Vehicle'.
type Vehicle interface {
	Start() string
}

// Definimos una estructura 'Car' vacía que representa un tipo de vehículo.
type Car struct{}

// Implementamos el método 'Start' para la estructura 'Car'.
// Esto hace que 'Car' satisfaga la interfaz 'Vehicle'.
func (c Car) Start() string {
	return "The Car is starting!"
}

// Definimos una estructura 'Motorcycle' vacía que representa otro tipo de vehículo.
type Motorcycle struct{}

// Implementamos el método 'Start' para la estructura 'Motorcycle'.
// Esto hace que 'Motorcycle' también satisfaga la interfaz 'Vehicle'.
func (m Motorcycle) Start() string {
	return "The Motorcycle is starting!"
}

// Definimos una estructura 'Bicycle' vacía que representa otro tipo de vehículo.
type Bicycle struct{}

// Implementamos el método 'Start' para la estructura 'Bicycle'.
// Esto hace que 'Bicycle' satisfaga la interfaz 'Vehicle'.
func (b Bicycle) Start() string {
	return "The Bicycle is in motion!"
}

// La función 'StartVehicle' toma como parámetro cualquier tipo que implemente la interfaz 'Vehicle'.
// Imprime el resultado del método 'Start' del vehículo que recibe, demostrando el polimorfismo.
func StartVehicle(v Vehicle) {
	fmt.Println(v.Start())
}

func main() {
	// Creación de instancias de diferentes tipos que satisfacen la interfaz 'Vehicle'.
	// Esto demuestra cómo 'StartVehicle' puede trabajar con diferentes tipos de vehículos.
	// Todos estos tipos implementan la misma interfaz 'Vehicle'.
	car := Car{}
	StartVehicle(car) // Output: The Car is starting!

	motorcycle := Motorcycle{}
	StartVehicle(motorcycle) // Output: The Motorcycle is starting!

	bicycle := Bicycle{}
	StartVehicle(bicycle) // Output: The Bicycle is in motion!
}
