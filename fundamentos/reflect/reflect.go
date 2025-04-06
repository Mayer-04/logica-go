package main

import (
	"fmt"
	"reflect"
)

/*
* Paquete reflect en Go
El paquete `reflect` permite trabajar con tipos y valores en tiempo de ejecución.

NOTA 📖
"Trabajar con tipos y valores en tiempo de ejecución" significa que puedes ver qué tipo tiene una variable
y qué valor tiene mientras el programa ya está corriendo, aunque no lo supieras antes.

- Permite inspeccionar y modificar variables, campos de estructuras, acceder a etiquetas (tags)
y cambiar valores dinámicamente.
- Es útil en escenarios como validación, serialización, frameworks o bibliotecas genéricas.
- Debe usarse con precaución, ya que puede afectar el rendimiento y hacer el código más complejo.
- Solo debe usarse `reflect` cuando sea necesario.

* Funciones del paquete reflect:
- `reflect.TypeOf()`: Devuelve el tipo dinámico de una variable.
- `reflect.ValueOf()`: Devuelve el valor dinámico de una variable.

*/

func main() {
	// `reflect.TypeOf()`: Determina el tipo de dato de las variables.
	number := 10
	text := "Hello, World!"
	boolean := true
	floatNumber := 3.14

	fmt.Println("El tipo de 'number' es:", reflect.TypeOf(number))           // int
	fmt.Println("El tipo de 'text' es:", reflect.TypeOf(text))               // string
	fmt.Println("El tipo de 'boolean' es:", reflect.TypeOf(boolean))         // bool
	fmt.Println("El tipo de 'floatNumber' es:", reflect.TypeOf(floatNumber)) // float64

	// `reflect.ValueOf()`: Obtiene el valor de las variables.
	valueOfNumber := reflect.ValueOf(number)
	valueOfText := reflect.ValueOf(text)

	fmt.Println("El valor de 'number' es:", valueOfNumber) // 10
	fmt.Println("El valor de 'text' es:", valueOfText)     // "Hello, World!"

	// Definición de una estructura con etiquetas (tags) para JSON y validación.
	type Person struct {
		Name  string `json:"name" validate:"required"`
		Age   int    `json:"age" validate:"min=0"`
		Email string `json:"email" validate:"email"`
	}

	// Instanciación de la estructura `Person`.
	person := Person{
		Name:  "Andres",
		Age:   24,
		Email: "andres@example.com",
	}

	// Obtener el tipo `reflect.Type` y el valor `reflect.Value` de la estructura.
	personType := reflect.TypeOf(person)
	personValue := reflect.ValueOf(person)

	fmt.Println("El tipo de 'person' es:", personType)   // Output: main.Person
	fmt.Println("El valor de 'person' es:", personValue) // Output: {Andres 24 andres@example.com}

	// Recorrer e inspeccionar los campos de la estructura.
	// `NumField()` devuelve el número de campos.
	// `Field(i)` devuelve la información del campo de la estructura en la posición `i`.
	for i := range personType.NumField() {
		field := personType.Field(i)  // Información del campo.
		value := personValue.Field(i) // Valor del campo.
		fmt.Printf("Campo: %s, Tipo: %s, Valor: %v\n", field.Name, field.Type, value)
	}

	// Acceso a las etiquetas (tags) definidas en los campos de la estructura.
	for i := range personType.NumField() {
		field := personType.Field(i)
		fmt.Printf("Campo: %s, Tag JSON: %s, Tag Validate: %s\n",
			field.Name,                // Nombre del campo.
			field.Tag.Get("json"),     // Valor de la etiqueta `json`.
			field.Tag.Get("validate"), // Valor de la etiqueta `validate`.
		)
	}

	// Modificación dinámica de un campo.
	// Es necesario pasar un puntero a la estructura para poder modificar sus valores.
	// `Elem()` obtiene el valor al que apunta el puntero.
	modifiablePerson := reflect.ValueOf(&person).Elem()

	if modifiablePerson.FieldByName("Name").CanSet() {
		modifiablePerson.FieldByName("Name").SetString("Diego") // Cambia el valor del campo Name a "Diego".
	}

	fmt.Printf("Nuevo valor de 'Name': %q\n", person.Name) // Output: "Diego"
}
