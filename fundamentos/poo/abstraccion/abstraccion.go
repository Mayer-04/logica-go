package main

import "fmt"

/*
* Abstracción en Go

La abstracción en Go consiste en modelar objetos del mundo real de forma simplificada,
enfocándonos solo en los atributos y comportamientos relevantes para nuestro programa.
Esto se logra principalmente mediante el uso de estructuras (structs) y métodos asociados.

La abstracción nos permite:
- Representar entidades reales (como un auto, una persona o un libro) en el código.
- Definir solo las características y acciones importantes, ocultando detalles innecesarios.
- Manipular instancias concretas de esas entidades de manera sencilla y segura.

* NOTAS:
- Usamos el operador punto (.) para acceder o modificar los campos de una estructura.
- Los métodos en Go se pueden asociar a cualquier tipo definido por el usuario.
- Es recomendable ser consistente en el uso de receptores (valor o puntero) en los métodos de una estructura.
- Los métodos se definen fuera de la declaración de la estructura.

* CONCEPTOS CLAVE:

- Abstracción:
	Crear un modelo simplificado de algo real, enfocándonos en lo esencial y ocultando la complejidad.
	Ejemplo: Para un auto, modelamos marca, modelo y color, ignorando detalles como cada tornillo.

- Entidad:
	Cualquier objeto o concepto que queremos representar en el programa.
	Ejemplos: Auto, Persona, Libro, CuentaBancaria.
	Cada entidad tiene atributos (color, nombre, saldo) y puede realizar acciones (avanzar, saludar, depositar).

- Instancia:
	Un ejemplar concreto de una entidad.
	Ejemplo: Si "Auto" es la estructura, "miAuto" (un Honda Civic azul) es una instancia específica.
*/

// Definimos una estructura `car` que abstrae las propiedades de un vehículo.
type car struct {
	model      string
	year       int
	price      float64
	components []string
}

// Método asociado a la estructura `car` que imprime los componentes del vehículo.
// El receptor `c` (de tipo valor) permite acceder a los campos de la instancia.
func (c car) printComponents() {
	message := "Los componentes del vehículo son: "
	for _, component := range c.components {
		message += component + ", "
	}
	// Eliminamos la última coma y espacio para que el mensaje se vea mejor.
	fmt.Println(message[:len(message)-2])
}

// Método que modifica el campo `year` de la estructura `car`.
// Se utiliza un receptor de tipo puntero para modificar el valor original del campo.
func (c *car) changeYear(year int) {
	c.year = year
}

func main() {
	// Creamos una instancia de `car` y la asignamos a la variable `nissan`.
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

	// Si un campo no se inicializa, se le asigna su valor por defecto (zero value) del tipo del campo.
	toyota := car{
		model: "Toyota",
		price: 100000.00,
	}
	fmt.Printf("toyota: %+v\n", toyota) // Output: {model:Toyota year:0 price:100000 components:[]}

	// Asignamos un valor al campo `year` usando el operador punto.
	toyota.year = 2022
	fmt.Printf("toyota: %+v\n", toyota) // Output: {model:Toyota year:2022 price:100000 components:[]}

	// Llamamos al método `printComponents` para imprimir los componentes del vehículo.
	nissan.printComponents() // Output: Los componentes del vehículo son: Engine, Wheels, Brakes

	// Cambiamos el campo `year` usando el método `changeYear`.
	// Se puede usar con un puntero explícito, pero no es necesario.
	(&nissan).changeYear(2000) // Válido pero no recomendado.
	nissan.changeYear(2000)    // Recomendado: Go lo convierte automáticamente a puntero.
	// El año se ha cambiado de 2024 a 2000.
	fmt.Printf("nissan: %+v\n", nissan) // {model:Nissan year:2000 price:200000 components:[Engine Wheels Brakes]}
}
