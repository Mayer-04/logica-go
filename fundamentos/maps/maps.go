package main

import (
	"fmt"
	"maps"
)

/*
* Maps: Mapas
Los Maps en Go son estructuras de datos no ordenadas que almacenan pares clave-valor.

- Una clave puede ser de cualquier tipo comparable (int, uint, string, struct, etc).
- Los valores pueden ser de cualquier tipo.
- Los mapas proporcionan acceso rápido a los datos y son muy útiles para almacenar conjuntos de datos
donde se necesita buscar elementos por clave.
- En un mapa `no` pueden haber dos claves iguales. Cada clave en un mapa debe ser única.
- Si intentas asignar un valor a una clave que ya existe en el mapa, el valor asociado con esa clave se actualizará.

* IMPORTANTE:
- Para evitar problemas al agregar elementos a un map, asegúrese de crearlo usando la función `make()`
para inicializarlo correctamente.
- El valor cero de un map no inicializado es `nil`, lo que significa que no apunta a ninguna estructura de datos.
- los mapas no son seguros para la concurrencia. Si tienes múltiples goroutines accediendo y modificando
el mapa simultáneamente, puede haber problemas de integridad de datos.
- Si deseas utilizar un mapa en múltiples goroutines, deberas usar un mutex, canales o sync.Map.
- `sync.Map` es una versión concurrente de los mapas que maneja de manera segura múltiples accesos.

* NOTA:
Desde la versión 1.24 de Go, los mapas se implementan internamente con `Swiss Tables`.
*/

func main() {
	// Creación de un mapa literal: se declara e inicializa con valores predeterminados.
	//* NOTA: En un mapa literal, es obligatorio colocar una coma al final de cada elemento.
	var edades = map[string]int{
		"Mayer":  24,
		"Andres": 23,
		"Luis":   25,
		"Maria":  26,
	}
	fmt.Println("Edades:", edades)

	// Declaración e inicialización de un mapa literal en una sola línea.
	nuevoMapa := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("nuevoMapa:", nuevoMapa)

	// Inicialización de un mapa vacío.
	// Esta técnica se usa para evitar un map nulo (`nil map`), ya que estos no permiten agregar elementos.
	mapaVacio := map[int]string{}
	fmt.Println("Mapa vacío:", mapaVacio) // Output: map[]

	// Creación e inicialización de un mapa vacío utilizando la función `make()`.
	// Se puede especificar la capacidad inicial como segundo argumento para mejorar el rendimiento.
	edades2 := make(map[string]int, 2)

	//* Agregando nuevos elementos al mapa.
	edades2["Mayer"] = 24
	edades2["Andres"] = 23
	fmt.Println("Edades2:", edades2) // Output: map[Andres:23 Mayer:24]

	//* Acceder a un valor del mapa mediante la clave.
	// Si la clave no existe, devuelve el valor cero del tipo correspondiente.
	fmt.Println("Valor de la llave 'a':", nuevoMapa["a"])
	fmt.Println("Valor de la llave 'd':", nuevoMapa["d"]) // Output: Valor de la llave 'd': 0

	// Uso de la función incorporada `delete()` para eliminar elementos del mapa.
	// Elimina la clave especificada y su valor asociado si existe.
	delete(edades2, "Mayer")         // Elimina la clave "Mayer" con su valor del map edades2.
	fmt.Println("Edades2:", edades2) // Output: map[Andres:23]

	//* Verificando si una clave existe en el mapa y obteniendo su valor.
	// Al acceder a un valor con una clave, se puede usar una segunda variable booleana ("ok")
	// para verificar si la clave existe en el mapa.
	// Si la clave no existe, el valor será el valor cero del tipo y "ok" será false.
	value, ok := nuevoMapa["a"]
	fmt.Printf("Valor: %v - Existe: %t\n", value, ok) // Output: Valor: 1 - Existe: true

	// Comprobando si una clave existe directamente en un condicional `if`.
	// En este caso, se accede a la clave y se evalúa su valor booleano.
	// Solo se ejecutará si el valor asociado a la clave es true.
	mapBool := map[string]bool{"a": true, "b": false}
	if mapBool["a"] {
		fmt.Printf("La clave 'a' existe en el mapa\n") // Esta línea será ejecutada.
	}

	// Eliminando todos los elementos del mapa usando la función `clear()`.
	// `clear()` es una función que simplifica la eliminación de todos los elementos de un mapa.
	clear(edades)
	fmt.Println("Edades después de clear():", edades) // Output: map[]

	// Recorriendo un mapa con un bucle `for range`.
	// El orden de los elementos en un mapa es aleatorio y puede variar en cada ejecución.
	// La primera variable representa la clave y la segunda el valor asociado.
	for nombre, edad := range edades2 {
		fmt.Printf("Nombre: %q, Edad: %d\n", nombre, edad)
	}

	//* Ejemplo que muestra que en un mapa no puede haber claves duplicadas.
	// Se declara un mapa vacío con la función `make()`.
	m := make(map[string]int)

	// Agregando un elemento con clave "a".
	m["a"] = 1
	fmt.Println(m) // Output: map[a:1]

	// Agregando nuevamente la misma clave con un valor diferente.
	// Esto sobrescribe el valor anterior.
	m["a"] = 2
	fmt.Println(m) // Output: map[a:2]

	// Agregando una nueva clave distinta.
	m["b"] = 3
	fmt.Println(m) // Output: map[a:2 b:3]

	// Función `Equal()`.
	// Verifica si dos mapas contienen los mismos pares clave/valor.
	// Los valores de cada par se comparan utilizando el operador ==.
	equal := maps.Equal(edades, edades2)
	fmt.Println("equal:", equal) // Output: false

	// Función `Copy()`.
	// Copia todos los pares clave/valor del mapa 'src' al mapa 'dst'.
	// Si una clave de 'src' ya existe en 'dst', su valor será reemplazado por el valor de 'src'.
	newMap := make(map[string]int)
	maps.Copy(newMap, edades2)
	fmt.Println("copy:", newMap) // Output: map[Andres:23]

	// Función `Clone()`.
	// Crea y devuelve una copia del mapa proporcionado.
	clone := maps.Clone(newMap)
	fmt.Println("clone:", clone) // Output: map[Andres:23]
}
