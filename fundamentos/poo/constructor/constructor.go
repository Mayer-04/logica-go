package main

import "fmt"

/*
* Constructores en Go
Go no tiene constructores como en otros lenguajes de programación orientado a objetos.
En su lugar, se utilizan funciones que retornan una instancia de una estructura (struct).

- Un constructor es una función que se utiliza para crear e inicializar una nueva instancia de una estructura.
- Por convención, estas funciones suelen llamarse `New` o `NewNombreStruct`.
- Las funciones constructoras pueden incluir validaciones para asegurar que los valores iniciales son correctos.
- Estas funciones pueden devolver punteros o valores según la necesidad.
- Generalmente, se devuelven punteros cuando se espera que la instancia pueda ser modificada después de su creación,
o cuando se desea evitar la copia de grandes estructuras. Se devuelven valores cuando se quiere una copia
independiente del objeto.
*/

// Definimos una estructura llamada `Persona`
type Persona struct {
	Nombre string
	Edad   int
}

//* New crea una nueva instancia de `Persona`.
// Recibe el nombre y la edad como parámetros y retorna un puntero a una nueva instancia de `Persona`.
func New(nombre string, edad int) *Persona {
	// Creamos y retornamos un puntero a `Persona` con los valores iniciales que se pasaron como parámetros.
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}

//* NewPersona crea una nueva instancia de `Persona` con validaciones.
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

//* NewPersonaValor crea una nueva instancia de `Persona` y la retorna por valor.
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
