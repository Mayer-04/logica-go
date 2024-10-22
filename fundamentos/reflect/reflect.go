package main

import (
	"fmt"
	"reflect"
)

/*
* Paquete reflect en Go
El paquete "reflect" proporciona métodos para trabajar con tipos y valores en tiempo de ejecución.

- Permite inspeccionar y modificar dinámicamente variables, estructuras y sus campos.
- Inspección y modificación de campos de estructuras.
- Solo aplicar el paquete `reflect` cuando sea necesario.
- El código puede tener menos rendimiento cuando se utiliza el paquete `reflect`.
- Acceso a etiquetas (tags) de los campos en estructuras.
- `reflect.TypeOf()`: Obtiene el tipo de dato de una variable.
- `reflect.ValueOf()`: Obtiene el valor de una variable.
*/

func main() {
	// `reflect.TypeOf()` - Determina el tipo de dato de las variables.
	number := 10
	text := "Hello, World!"
	boolean := true
	floatNumber := 3.14

	fmt.Println("El tipo de 'number' es:", reflect.TypeOf(number))           // int
	fmt.Println("El tipo de 'text' es:", reflect.TypeOf(text))               // string
	fmt.Println("El tipo de 'boolean' es:", reflect.TypeOf(boolean))         // bool
	fmt.Println("El tipo de 'floatNumber' es:", reflect.TypeOf(floatNumber)) // float64

	// `reflect.ValueOf()` - Obtiene el valor de las variables.
	valueOfNumber := reflect.ValueOf(number)
	valueOfText := reflect.ValueOf(text)

	fmt.Println("El valor de 'number' es:", valueOfNumber) // Output: 10
	fmt.Println("El valor de 'text' es:", valueOfText)     // Output: "Hello, World!"

	// Definición de una estructura con etiquetas (tags) para JSON y validación.
	type Person struct {
		Name  string `json:"name" validate:"required"`
		Age   int    `json:"age" validate:"min=0"`
		Email string `json:"email" validate:"email"`
	}

	// Instanciación de la estructura Person.
	person := Person{
		Name:  "Andres",
		Age:   24,
		Email: "andres@example.com",
	}

	// Obtener el tipo (`reflect.Type`) y valor (`reflect.Value`) de la estructura.
	personType := reflect.TypeOf(person)
	personValue := reflect.ValueOf(person)

	// Inspección de campos de la estructura usando `NumField()` y `Field()`.
	// `NumField()` obtiene el número de campos de la estructura.
	// `Field()` obtiene la información de cada campo de la estructura.
	for i := 0; i < personType.NumField(); i++ {
		field := personType.Field(i)  // Obtiene la información del campo.
		value := personValue.Field(i) // Obtiene el valor del campo.
		fmt.Printf("Campo: %s, Tipo: %s, Valor: %v\n", field.Name, field.Type, value)
	}

	// Acceso a las etiquetas (tags) de los campos de la estructura.
	for i := 0; i < personType.NumField(); i++ {
		field := personType.Field(i)
		fmt.Printf("Campo: %s, Tag JSON: %s, Tag Validate: %s\n",
			field.Name,                // Nombre del campo.
			field.Tag.Get("json"),     // Obtener el valor de la etiqueta JSON.
			field.Tag.Get("validate"), // Obtener el valor de la etiqueta Validate.
		)
	}

	// Modificar dinámicamente un campo (debe ser un puntero para ser modificable).
	modifiablePerson := reflect.ValueOf(&person).Elem() // Usar Elem() para obtener el valor subyacente.

	if modifiablePerson.FieldByName("Name").CanSet() {
		modifiablePerson.FieldByName("Name").SetString("Diego") // Cambia el nombre a "Diego".
	}

	fmt.Printf("Nuevo valor de 'Name': %q\n", person.Name) // Output: "Diego"
}
