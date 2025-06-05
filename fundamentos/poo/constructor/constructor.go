package main

import "fmt"

/*
* Constructores en Go
En Go, creamos nuevos objetos de una manera diferente a otros lenguajes de programación.
En lugar de tener una función especial llamada "constructor", simplemente usamos
funciones normales que crean y devuelven nuevas instancias de nuestras estructuras.

* Características de los constructores en Go:

- Un constructor es simplemente una función que crea un nuevo objeto y lo configura
  con valores iniciales apropiados.

- Por costumbre entre programadores de Go, estas funciones suelen llamarse "New" o
  "NewAlgo" (como "NewUsuario" o "NewCoche").

- Estas funciones pueden incluir comprobaciones para asegurarse de que los datos
  son correctos. Por ejemplo, verificar que una edad no sea negativa.

- Los constructores pueden devolver el objeto de dos maneras:
  • Como valor normal: Cuando queremos una copia independiente del objeto.
  • Como puntero: Cuando queremos poder modificar el objeto original después,
    o cuando el objeto es muy grande y queremos evitar hacer copias.
*/

// Definimos una estructura llamada `Persona`.
type Persona struct {
	Nombre string
	Edad   int
}

// * New crea una nueva instancia de `Persona`.
// Recibe el nombre y la edad como parámetros y retorna un puntero a una nueva instancia de `Persona`.
func New(nombre string, edad int) *Persona {
	// Creamos y retornamos un puntero a `Persona` con los valores iniciales que se pasaron como parámetros.
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}

// * NewPersona crea una nueva instancia de `Persona` con validaciones.
// Si el nombre es vacío o la edad es negativa, retorna un error.
func NewPersona(nombre string, edad int) (*Persona, error) {
	// Validación para nombre vacío
	if nombre == "" {
		return nil, fmt.Errorf("el nombre no puede estar vacío")
	}

	// Validación para edad negativa
	if edad < 0 {
		return nil, fmt.Errorf("la edad no puede ser negativa")
	}

	// Si los datos son válidos, retorna un puntero a `Persona`.
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}, nil
}

// * NewPersonaValor crea una nueva instancia de `Persona` y la retorna por valor.
// Se utiliza cuando no se necesita modificar la instancia después de la creación, o cuando la estructura es pequeña.
func NewPersonaValor(nombre string, edad int) Persona {
	// Retornamos un valor de tipo `Persona` con los valores iniciales.
	return Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}

func main() {
	// Crear una nueva instancia de `Persona` usando el constructor básico `New`.
	persona := New("Mayer", 24)
	// `persona` es un puntero a `Persona`.
	fmt.Printf("persona: %+v\n", persona) // Output: persona: &{Nombre:Mayer Edad:24}
	// Acceder a los valores de `persona` usando el operador de desreferenciación `*`.
	fmt.Printf("persona: %+v\n", *persona)                             // Output: persona: {Nombre:Mayer Edad:24}
	fmt.Printf("nombre: %s, edad: %d\n", persona.Nombre, persona.Edad) // Output: nombre: Mayer, edad: 24

	// Crear una nueva instancia de `Persona` usando el constructor `NewPersona` con validaciones.
	persona2, err := NewPersona("Andres", 24)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Si no hubo errores, se imprimen los valores de `persona2`.
	fmt.Printf("persona2: %+v\n", persona2)
	fmt.Printf("nombre: %s, edad: %d\n", persona2.Nombre, persona2.Edad)

	// Crear una nueva instancia de `Persona` usando el constructor `NewPersonaValor` que retorna un valor.
	persona3 := NewPersonaValor("Robert", 59)
	// `persona3` es una estructura `Persona` por valor, no un puntero.
	fmt.Printf("persona3: %+v\n", persona3) // Output: persona3: {Nombre:Robert Edad:59}
	// Como `persona3` es un valor, no necesitamos desreferenciar para acceder a sus campos.
	fmt.Printf("nombre: %s, edad: %d\n", persona3.Nombre, persona3.Edad) // Output: nombre: Robert, edad: 59
}
