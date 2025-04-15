package main

import (
	"fmt"
	"reflect"
)

/*
Tienes una función que recibe un objeto como parámetro.
La función debe retornar un array con el nombre de las propiedades que su tipo sea boolean.

Por ejemplo, para el objeto { a: true, b: 42, c: false } la función debe retornar ['a', 'c'],
ya que son las dos propiedades que tienen valores booleanos.
*/

func main() {
	booleanProperties := map[string]any{
		"a": true,
		"b": 42,
		"c": false,
	}
	result := booleanPropertyKeys(booleanProperties)
	fmt.Printf("result: %v\n", result) // ["a", "c"]

	// Solución 2
	result2 := booleanPropertyKeys2(booleanProperties)
	fmt.Printf("result2: %v", result2) // ["a", "c"]
}

// Solución 1 utilizando genericos ✅
func booleanPropertyKeys[K string, V any](obj map[K]V) []K {
	booleanKeys := make([]K, 0, len(obj))

	for key, value := range obj {
		if reflect.TypeOf(value).Kind() == reflect.Bool {
			booleanKeys = append(booleanKeys, key)
		}
	}

	return booleanKeys
}

// Solución 2 ✅
func booleanPropertyKeys2(obj map[string]any) []string {
	var booleanKeys []string

	for key, value := range obj {
		typeof := reflect.TypeOf(value).Kind()
		if typeof == reflect.Bool {
			booleanKeys = append(booleanKeys, key)
		}
	}

	return booleanKeys
}
