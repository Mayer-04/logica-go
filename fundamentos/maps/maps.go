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

IMPORTANTE:
- Para evitar problemas al agregar elementos a un map, asegúrese de crearlo usando la función `make()`
para inicializarlo correctamente.
- El valor cero de un map no inicializado es `nil`, lo que significa que no apunta a ninguna estructura de datos.
- los mapas no son seguros para la concurrencia. Si tienes múltiples goroutines accediendo y modificando
el mapa simultáneamente, puede haber problemas de integridad de datos.
- Si deseas utilizar un mapa en múltiples goroutines, deberas usar un mutex, canales o sync.Map para coordinar el acceso.
- sync.Map es una versión concurrente de los mapas que maneja de manera segura múltiples accesos.
*/

func main() {
	// Creando un map literal - Declarando e inicializando el map con valores predeterminados.
	// NOTA: Es obligatorio incluir una coma al final de cada elemento en un mapa literal.
	var edades = map[string]int{
		"Mayer":  24,
		"Andres": 23,
		"Luis":   25,
		"Maria":  26,
	}
	fmt.Println("Edades:", edades)

	// Creando un map literal con valores iniciales en una sola línea.
	nuevoMapa := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("nuevoMapa:", nuevoMapa)

	// Inicializando un map vacío utilizando un map literal.
	// Es común usar esta técnica para evitar un nil map, el cual no permite agregar elementos.
	mapaVacio := map[int]string{}
	fmt.Println("Mapa vacío:", mapaVacio) // Output: map[]

	// Creando e inicializando un mapa vacío utilizando `make()`.
	// Podemos especificar la capacidad inicial como segundo argumento del map para optimizar el rendimiento.
	edades2 := make(map[string]int, 2)

	//* Agregando nuevos elementos al map.
	edades2["Mayer"] = 24
	edades2["Andres"] = 23
	fmt.Println("Edades2:", edades2)

	//* Accediendo a el valor un elemento del map.
	// Si la clave no existe, se devuelve el valor cero del tipo de dato correspondiente.
	fmt.Println("Valor de la llave 'a':", nuevoMapa["a"])
	fmt.Println("Valor de la llave 'd':", nuevoMapa["d"]) // Output: Valor de la llave 'd': 0

	// La función incorporada `delete()` elimina el elemento con la clave especificada del mapa.
	delete(edades2, "Mayer") // Eliminando del map edades2 la clave "Mayer" con su valor.
	fmt.Println("Edades2:", edades2)

	//* Verificando si una clave existe en el mapa y obteniendo su valor.
	// "value" es el valor de la clave "a". Si la clave no existe, "value" será el valor cero.
	// "ok" es un booleano que indica si la clave existe en el map.
	value, ok := nuevoMapa["a"]
	fmt.Printf("Valor: %v - Existe: %t\n", value, ok)

	// Comprobando si la clave existe con un if.
	// Las claves del mapa deben tener un valor `booleano` para hacer esta comprobación.
	mapBool := map[string]bool{"a": true, "b": false}
	if mapBool["a"] {
		fmt.Printf("La clave 'a' existe en el map.")
	}

	// Eliminando todos los elementos del map usando la función `clear()`.
	// `clear()` es una función que simplifica la eliminación de todos los elementos de un map.
	clear(edades)
	fmt.Println("Edades después de clear():", edades) // Output: map[]

	// Recorriendo un map con un bucle `for range`.
	// El orden de iteración de los elementos en un map es aleatorio y puede cambiar entre ejecuciones del programa.
	// La primera variable son las "claves", la segunda variable son los "valores" del map.
	for nombre, edad := range edades2 {
		fmt.Printf("Nombre: %q, Edad: %d\n", nombre, edad)
	}

	//* Ejemplo de un mapa donde no pueden haber dos claves iguales.
	// Declarando un mapa con la función `make`.
	m := make(map[string]int)

	// Agregando un valor.
	m["a"] = 1
	fmt.Println(m) // Output: map[a:1]

	// Agregar la misma clave con un valor diferente.
	// Se actualiza el valor de la clave "a" a 2. El valor 1 se reemplaza por 2.
	m["a"] = 2
	fmt.Println(m) // Output: map[a:2]

	// Agregar una clave diferente.
	m["b"] = 3
	fmt.Println(m) // Output: map[a:2 b:3]

	// `Equal()`.
	// Verifica si dos mapas contienen los mismos pares clave/valor.
	// Los valores de cada par se comparan utilizando el operador ==.
	equal := maps.Equal(edades, edades2)
	fmt.Println("equal:", equal) // Output: false

	// `Copy()`.
	// Copia todos los pares clave/valor del mapa 'src' al mapa 'dst'.
	// Si una clave de 'src' ya existe en 'dst', su valor será reemplazado por el valor de 'src'.
	newMap := make(map[string]int)
	maps.Copy(newMap, edades2)
	fmt.Println("copy:", newMap)

	// `Clone()`.
	// Crea y devuelve una copia del mapa proporcionado.
	clone := maps.Clone(newMap)
	fmt.Println("clone:", clone)
}
