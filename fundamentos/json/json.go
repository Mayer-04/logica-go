package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*
* JSON y Etiquetas en Go
- Las etiquetas en las estructuras (struct), como `json:"name"`, son instrucciones que le indican a Go cómo
codificar y decodificar (serializar/deserializar) las estructuras en formatos como JSON.
- `json.Unmarshal` prefiere una coincidencia exacta de la clave en la estructura. Sin embargo, si encuentra varias
claves en el JSON que solo se diferencian en mayúsculas/minúsculas, selecciona la última clave que coincide
(sensible a mayúsculas/minúsculas).

* Las etiquetas son cadenas que siguen el formato `json:"clave,opciones"`:
- `"clave"`: El nombre de la clave en el JSON.
- `"-"`: Indica que el campo debe ser ignorado en la codificación y decodificación JSON.
- `"omitempty"`: El campo solo se incluirá si su valor no es el valor cero del tipo
(por ejemplo, 0 para enteros, "" para cadenas, nil para punteros, etc.).
- `"string"`: Convierte automáticamente un campo numérico o booleano a una cadena al serializarlo en JSON.
- Por defecto, si no se especifica una etiqueta, Go usará el nombre del campo de la estructura
tal como está para el nombre de la clave en JSON.

IMPORTANTE: Puedes usar el paquete `reflect` para inspeccionar las etiquetas de los campos en tiempo de ejecución.

* Deserializar (Unmarshal):
- Es el proceso de convertir una secuencia de bytes (JSON) en una estructura de datos de Go.
* Serializar (Marshal):
- Es convertir una estructura de datos de Go en una secuencia de bytes, típicamente en formato JSON.
* Decodificar:
- Se refiere a la deserialización (como el proceso de convertir JSON a estructuras de Go), aunque en el contexto
de JSON, la palabra "decodificar" y "deserializar" suelen usarse de manera intercambiable.
*/

type Person struct {
	// El campo `Name` se convertirá en la clave "name" en JSON.
	Name string `json:"name"`

	// La etiqueta "-" indica que este campo se omitirá en JSON.
	Lastname string `json:"-"`

	// El campo `Age` se convertirá en la clave "age" en JSON.
	Age int `json:"age"`

	// La etiqueta "omitempty" indica que el campo solo se incluirá en el JSON si su valor no es el valor cero.
	Active bool `json:"active,omitempty"`

	// La etiqueta "string" convierte campos numéricos o booleanos en cadenas en JSON.
	Play bool `json:"play,string"`
}

type Country struct {
	// `Nombre` se serializa como "nombre" en el JSON.
	Nombre string `json:"nombre"`

	// `Habitantes` se serializa como "habitantes".
	Habitantes int `json:"habitantes"`

	// `Capital` se serializa como "capital".
	Capital string `json:"capital"`

	// `Idiomas` se serializa como un array de cadenas en JSON.
	Idiomas []string `json:"idiomas"`
}

type Person2 struct {
	// Se espera que coincida con "name" o "Name" en el JSON, pero Go toma la última coincidencia encontrada.
	Name string `json:"name"`
}

func main() {
	// Crear una instancia de Person con algunos campos.
	mayer := Person{
		Name:     "Mayer",
		Lastname: "Chaves",
		Age:      24,
		Active:   false, // El campo Active no será incluido en el JSON porque es false y tiene la etiqueta "omitempty".
		Play:     false, // El valor de Play será representado como "false" en JSON (gracias a la etiqueta `string`).
	}

	// Convertir la estructura Person a JSON (serialización).
	jsonData, err := json.Marshal(mayer)
	if err != nil {
		panic(err)
	}

	// Convertir el slice de bytes a una cadena para imprimirlo.
	str := string(jsonData)
	fmt.Printf("JSON de Person: %s\n", str) // Output: {"name":"Mayer","age":24,"play":"false"}

	// Ejemplo de JSON para deserialización (Unmarshal).
	exampleJson := `{"nombre":"Canada","habitantes":37314442,"capital":"Ottawa","idiomas":["Inglés","Francés"]}`
	// Crear una instancia vacía de Country.
	var countrys Country

	// Convertir la cadena JSON a una instancia de Country (deserialización).
	if err := json.Unmarshal([]byte(exampleJson), &countrys); err != nil {
		panic(err)
	}

	// Imprimir la estructura deserializada.
	fmt.Printf("estructura Country: %+v\n", countrys) // Output: {Nombre:Canada Habitantes:37314442 Capital:Ottawa Idiomas:[Inglés Francés]}

	// Usar `reflect` para inspeccionar etiquetas de los campos en la estructura Country.
	t := reflect.TypeOf(Country{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		typeName := field.Type.Name()
		// Imprimir el nombre del campo, el tipo y la etiqueta JSON asociada.
		fmt.Printf("field %s (%s), tag: %q\n", field.Name, typeName, tag)
	}

	// JSON con claves similares pero diferentes en mayúsculas/minúsculas.
	jsonStr := `{"name": "Diego", "Name": "Lucas"}`

	var p Person2
	// Deserializar el JSON en la estructura Person2.
	if err := json.Unmarshal([]byte(jsonStr), &p); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// En este caso, Go preferirá la clave "Name" ya que es la última coincidencia.
	fmt.Printf("Resultado deserializado: %+v\n", p) // Output: {Name:Lucas}
	fmt.Printf("Nombre: %s\n", p.Name)              // Output: Lucas
}