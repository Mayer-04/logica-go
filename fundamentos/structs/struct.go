package main

import (
	"fmt"
	"unsafe"
)

/*
* Struct: Estructura
Una Struct es una colección de campos que se pueden definir en un paquete.
- En Go, las estructuras se suelen declarar a nivel de paquete, no dentro de funciones.
- Es similar a una "clase" en otros lenguajes de programación, pero sin métodos asociados directamente.
- Las estructuras en Go no soportan herencia, pero se puede lograr composición mediante la 'incrustación' de estructuras.
- Internamente, las estructuras se representan un un bloque contiguo en la memoria.
Los campos son colocados uno tras otro en el mismo orden que fueron colocados en la estructura.

* Estructuras anidadas
- Se definen dentro de otra estructura como un campo con nombre explícito.
- Para acceder a sus campos, es necesario hacer referencia a la estructura anidada.
** Usos recomendados:
- Cuando prefieres mantener una separación clara entre la estructura principal y la secundaria.
- Mantener la independencia entre las estructuras.

* Estructuras incrustadas
- Se definen dentro de otra estructura 'sin' un nombre explícito.
- Los campos de la estructura incrustada se 'promueven' a la estructura principal.
- Se puede acceder a estos campos directamente, sin necesidad de hacer referencia a la estructura incrustada.
** Usos recomendados:
- Cuando deseas reutilizar los campos y métodos de otra estructura sin necesidad de referenciarla explícitamente.

* Estructuras vacías
- Es una estrcutura que no contiene ningún campo.
- Una estructura vacía en Go ocupa '0 bytes' en memoria por lo que no necesita reservar espacio.
- Se suelen utilizar en mapas, canales y optimización de memoria.
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

// Empleado es una estructura que demuestra la `composición` mediante la incrustación de la estructura Persona.
// Persona es una estructura 'incrustada' en Empleado.
// El campo Direccion es una estructura 'anidada' dentro de Empleado.
type Empleado struct {
	Persona   // Composición: Incluye todos los campos de 'Persona'
	Puesto    string
	Direccion Direccion // Campo de tipo 'Dirección': anidado
}

func main() {
	// Creación de una objeto "maria" de tipo Persona con los valores cero para cada campo.
	var maria Persona
	fmt.Printf("maria: %+v\n", maria) // Output: {Nombre: Edad:0}

	// Creación e inicialización de una Persona con todos sus campos.
	luis := Persona{
		Nombre: "Luis",
		Edad:   60,
	}
	fmt.Printf("luis: %+v\n", luis) // Output: luis: {Nombre:Luis Edad:60}

	//* Definición de un campo.
	// Los campos no especificados tomarán el valor cero de su tipo de dato correspondiente.
	// No es obligatorio inicializar todos los campos, pero es recomendable hacerlo.
	andres := Persona{
		Nombre: "Andres",
	}
	fmt.Printf("andres: %+v\n", andres) // Output: andres: {Nombre:Andres Edad:0}

	// Asignación de valores en orden sin especificar nombres de campos.
	// Aunque no es obligatorio especificar los nombres de los campos, hacerlo aumenta la claridad del código.
	mayer := Persona{"Mayer", 25}
	fmt.Printf("mayer: %+v\n", mayer) // Output: mayer: {Nombre:Mayer Edad:25}

	// Uso del formato campo: valor para inicializar los campos en cualquier orden.
	lucas := Persona{Edad: 40, Nombre: "Lucas"}
	fmt.Printf("lucas: %+v\n", lucas) // Output: lucas: {Nombre:Lucas Edad:40}

	// Acceder a los campos de la estructura utilizando el operador punto (.).
	fmt.Println("Nombre:", lucas.Nombre) // Output: Lucas
	fmt.Println("Edad:", lucas.Edad)     // Output: 40

	// Definición e instanciación de un estructura anónima.
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
	fmt.Printf("Struct anónimo: %+v\n", objeto) // Output: Struct anónimo: {A:true B:42 C:"Hola"}

	// Uso del operador `&` para obtener un puntero a un struct.
	// Los cambios realizados a través del puntero afectarán a la estructura original.
	objetoCopy := &objeto
	objetoCopy.B = 24
	fmt.Printf("Struct original modificado a través del puntero: %+v\n", objeto) // Output: {A:true B:24 C:Hola}

	// Comparación de structs.
	// Los structs se pueden comparar directamente si sus campos son comparables.
	otroLuis := Persona{Nombre: "Luis", Edad: 60}
	fmt.Println("¿Son luis y otroLuis iguales?", luis == otroLuis) // Output: true

	//* Uso de structs con slices y maps.
	personas := []Persona{luis, andres, mayer}
	fmt.Println("Slice de personas:", personas) // Output: [{Luis 25} {Andres 0} {Mayer 23}]

	personasMap := map[string]Persona{
		"luis":   luis,
		"andres": andres,
		"mayer":  mayer,
	}
	fmt.Println("Map de personas:", personasMap) // Output: map[andres:{Andres 0} luis:{Luis 25} mayer:{Mayer 23}]

	// Demostración de composición (incrustación) y estructuras anidadas.
	empleado := Empleado{
		Persona: Persona{Nombre: "Maria", Edad: 47}, // Composición mediante incrustación
		Puesto:  "Desarrolladora",
		Direccion: Direccion{ // Estructura anidada
			Calle:  "Calle Principal 123",
			Ciudad: "Medellín",
			CP:     "12345",
		},
	}
	fmt.Printf("Empleado: %+v\n", empleado)
	fmt.Println("Nombre del empleado:", empleado.Nombre) // Acceso directo al campo de la estructura embebida
	//* NOTA: Para acceder al campo 'Calle' tenemos que acceder a la estructura anidada 'Direccion' primero.
	fmt.Println("Calle del empleado:", empleado.Direccion.Calle) // Acceso a la estructura anidada

	// Estructura vacía anónima (sin nombre).
	// struct {} es el tipo que es vacío y '{}' es el valor (Una estructura vacía literal).
	empty := struct{}{}
	fmt.Printf("Espacio en memoria: %d bytes, valor: %v\n", unsafe.Sizeof(empty), empty) // Output: 0, {}

	// Estructura anónima con nombre.
	type emptyStruct struct{}
	empty2 := emptyStruct{}
	fmt.Printf("Espacio en memoria: %d bytes, valor: %v\n", unsafe.Sizeof(empty2), empty2) // Output: 0, {}
}
