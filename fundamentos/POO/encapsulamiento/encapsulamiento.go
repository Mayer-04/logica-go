package main

import "fmt"

/*
* Encapsulamiento en Go
En Go, el concepto de encapsulamiento se logra utilizando la visibilidad de los identificadores,
tanto en los campos de las estructuras como en los métodos.

* Campos:
- Si un campo comienza con una letra mayúscula, es público (exportado) y puede ser accedido desde otros paquetes.
- Si un campo comienza con una letra minúscula, es privado (no exportado) y solo puede ser accedido desde el paquete donde fue definido.

* Métodos:
- Si un método comienza con una letra mayúscula, es público y puede ser llamado desde otros paquetes.
- Si un método comienza con una letra minúscula, es privado y solo puede ser llamado desde el paquete donde fue definido.
*/

// La estructura `Person` es pública porque su nombre comienza con la primera letra en mayúscula.
// Esto significa que puede ser utilizada en otros paquetes.
type Person struct {
	Name    string   // Campo público: puede ser accedido desde otros paquetes.
	Age     int      // Campo público: puede ser accedido desde otros paquetes.
	Live    bool     // Campo público: puede ser accedido desde otros paquetes.
	hobbies []string // Campo privado: solo puede ser accedido dentro del paquete donde se definió.
}

// GreetPerson es un método público que imprime un saludo con el nombre de la persona.
func (p Person) GreetPerson() {
	fmt.Println("Hello, my name is", p.Name)
}

// Creando un método `getter`, en Go no se recomienda poner `Get` antes del nombre del método.
// `yourAge` es un método privado que retorna la edad de la persona.
// Solo puede ser llamado dentro del paquete donde se definió.
func (p Person) yourAge() int {
	return p.Age
}

// Creando un método `setter`, se puede utilizar `set` antes del nombre del método.
// `setAge` es un método privado que establece la edad de la persona.
// Utiliza un receptor de tipo `puntero` para modificar el valor original del campo.
// Solo puede ser llamado dentro del paquete donde se definió.
func (p *Person) setAge(age int) {
	p.Age = age
}

func main() {
	// Creamos una nueva instancia de `Person`.
	// En este caso,`hobbies` no puede ser accedido fuera del paquete, pero puede ser inicializado dentro de `main`.
	andres := Person{Name: "Andres", Age: 24, Live: true, hobbies: []string{"programar", "leer"}}

	// Imprimimos los valores de la instancia.
	// Los campos públicos se mostrarán, pero `hobbies` no porque es privado.
	// Este comportamiento es distinto si imprimimos `andres` con `fmt.Printf`.
	fmt.Printf("andres: %+v\n", andres) // Output: andres: {Name:Andres Age:24 Live:true hobbies:[programar leer]}

	// Llamamos al método público `GreetPerson`.
	andres.GreetPerson() // Output: Hello, my name is Andres

	// Llamamos al método privado `yourAge`.
	// Aunque es privado, como estamos dentro del mismo paquete, esto funcionará.
	age := andres.yourAge()
	fmt.Printf("Andres tiene %d años.\n", age) // Output: Andres tiene 24 años.

	// Llamamos al método privado `setAge`.
	andres.setAge(25)
	fmt.Printf("setAge: %+v\n", andres) // Output: setAge: {Name:Andres Age:25 Live:true hobbies:[programar leer]}

	// IMPORTANTE: El campo `Age` del objeto `andres` se ha modificado por completo.
	// Esto es debido a que usamos un receptor de tipo `puntero` para modificar el valor original del campo.
	fmt.Printf("andres: %+v\n", andres) // Output: andres: {Name:Andres Age:25 Live:true hobbies:[programar leer]}
}
