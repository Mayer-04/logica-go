package main

import "fmt"

/*
* Struct: Estructura
- Una Struct es una colección de campos que se pueden definir en un paquete.
- En Go, las estructuras suelen declararse a nivel de paquete y no a nivel de funciones.
- Es similar a una "clase" en otros lenguajes de programación, pero sin métodos asociados directamente.
- Las estructuras en Go no soportan herencia, pero se puede lograr composición mediante la incrustación de structs.
*/

// Persona es una `estructura` que define un objeto con Nombre y Edad.
type Persona struct {
	Nombre string
	Edad   uint8
}

// Rectangulo es una estructura que demuestra cómo definir múltiples campos del mismo tipo en una línea.
type Rectangulo struct {
	Ancho, Alto float64
}

// Dirección es una estructura que se utilizará en la estructura Empleado.
type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
}

// Empleado es una estructura que demuestra la composición mediante la incrustación de la estructura Persona.
type Empleado struct {
	Persona   // Composición: incluye todos los campos de Persona
	Puesto    string
	Direccion Direccion // Campo de tipo struct
}

func main() {
	// Creación de una objeto "maria" de tipo Persona con los valores cero para cada campo.
	var maria Persona
	fmt.Println("maria (valores cero):", maria) // Output: {"", 0}

	// Creación e inicialización de una Persona con todos sus campos.
	luis := Persona{
		Nombre: "Luis",
		Edad:   25,
	}
	fmt.Printf("luis: %+v\n", luis) // Output: luis: {Nombre:Luis Edad:25}

	//* Definición de un campo.
	// Los campos no especificados tomarán el valor cero de su tipo de dato correspondiente.
	// No es obligatorio inicializar todos los campos, pero es recomendable hacerlo.
	andres := Persona{
		Nombre: "Andres",
	}
	fmt.Println("andres:", andres)

	// Asignación de valores en orden sin especificar nombres de campos.
	// Aunque no es obligatorio especificar los nombres de los campos, hacerlo aumenta la claridad del código.
	mayer := Persona{"Mayer", 23}
	fmt.Println("mayer:", mayer)

	// Uso del formato campo: valor para inicializar los campos en cualquier orden.
	lucas := Persona{Edad: 40, Nombre: "Lucas"}
	fmt.Println("lucas:", lucas)

	// Acceder a los campos de la estructura utilizando el operador punto (.).
	fmt.Println("Nombre de Luis:", luis.Nombre)
	fmt.Println("Edad de Luis:", luis.Edad)

	// Definición e instanciación de un struct anónimo.
	// Son útiles cuando solo se usan en un contexto específico y no es necesario reutilizarlos.
	objeto := struct {
		A bool
		B int
		C string
	}{
		A: true,
		B: 42,
		C: "Hola",
	}
	fmt.Printf("Struct anónimo: %+v\n", objeto)

	// Uso del operador & para obtener un puntero a un struct.
	// Los cambios realizados a través del puntero afectarán a la estructura original.
	objetoCopy := &objeto
	objetoCopy.B = 24
	fmt.Printf("Struct original modificado a través del puntero: %+v\n", objeto)

	// Comparación de structs.
	// Los structs se pueden comparar directamente si sus campos son comparables.
	otroLuis := Persona{Nombre: "Luis", Edad: 25}
	fmt.Println("¿Son luis y otroLuis iguales?", luis == otroLuis)

	//* Uso de structs con slices y maps.
	personas := []Persona{luis, andres, mayer}
	fmt.Println("Slice de personas:", personas) // Output: [{Luis 25} {Andres 0} {Mayer 23}]

	personasMap := map[string]Persona{
		"luis":   luis,
		"andres": andres,
		"mayer":  mayer,
	}
	fmt.Println("Map de personas:", personasMap) // Output: map[andres:{Andres 0} luis:{Luis 25} mayer:{Mayer 23}]

	// Demostración de composición y estructuras anidadas
	empleado := Empleado{
		Persona: Persona{Nombre: "Maria", Edad: 47},
		Puesto:  "Desarrolladora",
		Direccion: Direccion{
			Calle:  "Calle Principal 123",
			Ciudad: "Medellín",
			CP:     "12345",
		},
	}
	fmt.Printf("Empleado: %+v\n", empleado)
	fmt.Println("Nombre del empleado:", empleado.Nombre) // Acceso directo al campo de la estructura embebida
}
