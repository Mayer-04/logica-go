package main

import "fmt"

/*
* Abstracción en Go
La abstracción en Go se logra mediante el uso de estructuras (structs) y métodos asociados a ellas.
Permite representar entidades complejas del mundo real y asociarles comportamientos específicos.

- El operador punto (.) se utiliza tanto para acceder como para asignar valores a los campos de una struct.
- En Go, los métodos pueden ser implementados para cualquier tipo de dato.
- Los métodos se declaran fuera de las struct.
- Se recomienda que todos los métodos asociados a una struct utilicen receptores de tipo puntero o todos de tipo valor,
pero no ambos, para evitar inconsistencias y errores.
*/

// Definimos una estructura `car` que abstrae las propiedades de un vehículo.
type car struct {
	model      string
	year       int
	price      float64
	components []string
}

// Método asociado a la estructura `car` para imprimir sus componentes.
// El receptor `c` permite acceder a los campos de la instancia de `car`.
func (c car) printComponents() {
	message := "Los componentes del vehículo son: "
	for _, component := range c.components {
		message += component + ", "
	}
	// Eliminamos la última coma y espacio para que el mensaje se vea mejor.
	fmt.Println(message[:len(message)-2])
}

// Método para modificar el año (`year`) de la estructura `car`.
// Se utiliza un receptor de tipo `puntero` para modificar el valor original del campo.
func (c *car) changeYear(year int) {
	c.year = year
}

func main() {
	// Creamos una instancia de `car` asignándola a la variable `nissan`.
	nissan := car{
		model:      "Nissan",
		year:       2024,
		price:      200000.00,
		components: []string{"Engine", "Wheels", "Brakes"},
	}
	fmt.Printf("myCar: %#v\n", nissan)

	// Accedemos a los campos de la estructura utilizando el operador punto (.).
	fmt.Printf("model: %s\n", nissan.model)
	fmt.Printf("year: %d\n", nissan.year)
	fmt.Printf("price: %.2f\n", nissan.price)
	fmt.Printf("components: %v\n", nissan.components)

	// Si un campo no se inicializa explícitamente, se le asigna el valor cero (0) para su tipo de dato.
	toyota := car{
		model: "Toyota",
		price: 100000.00,
	}
	fmt.Printf("toyota: %+v\n", toyota) // Output: {model:Toyota year:0 price:100000 components:[]}

	// Usamos el operador punto (.) para asignar un valor al campo `year`.
	toyota.year = 2022
	fmt.Printf("toyota: %+v\n", toyota) // Output: {model:Toyota year:2022 price:100000 components:[]}

	// Llamamos al método `printComponents` para imprimir los componentes del vehículo.
	nissan.printComponents() // Output: Los componentes del vehículo son: Engine, Wheels, Brakes

	// Llamamos al método `changeYear` para cambiar el valor del campo `year`.
	// Aunque se puede hacer explícitamente con un puntero, la manera recomendada es como sigue.
	(&nissan).changeYear(2000) // No recomendado, pero válido.
	nissan.changeYear(2000)    // Recomendado: Go lo convierte automáticamente en un puntero.
	// El año se ha cambiado de 2024 a 2000.
	fmt.Printf("nissan: %+v\n", nissan) // {model:Nissan year:2000 price:200000 components:[Engine Wheels Brakes]}
}
