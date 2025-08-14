package main

import "fmt"

/*
* Sets: Conjuntos en Go
- En Go no existen estructuras de conjunto (set) por defecto.
- Podemos simular conjuntos utilizando mapas, donde las claves son los elementos del conjunto y los valores son booleanos.
- Solo se permiten tipos comparables como claves (int, string, etc.),
por lo que tipos como funciones, mapas o slices no pueden ser usados.
- En un conjunto implementado como un mapa, solo las claves son relevantes; los valores son meramente `marcadores`.
- Un mapa en Go NO puede contener claves duplicadas, lo que lo convierte en una estructura adecuada para
representar un conjunto.

* Recomendaciones de paquetes:
- https://github.com/emirpasic/gods
- https://github.com/deckarep/golang-set
*/

func main() {
	// Creando un set inicial con algunos elementos.
	set := map[string]bool{
		"a": true,
		"b": true,
	}
	fmt.Printf("Conjunto inicial: %#v\n", set)

	// Creando un set utilizando una estructura vacía para mayor eficiencia.
	// Se recomienda más esta forma porque una estructura vacía ocupa '0 bytes' en memoria.
	set2 := map[int]struct{}{
		1: {},
		2: {},
	}
	fmt.Printf("Conjunto con estructura vacía: %#v\n", set2)

	// Añadiendo elementos a un set.
	set3 := make(map[int]struct{})

	set3[1] = struct{}{}
	set3[2] = struct{}{}
	set3[3] = struct{}{}
	// Intentando añadir un elemento duplicado, no tendrá efecto.
	set3[1] = struct{}{}

	fmt.Printf("Conjunto después de añadir elementos: %#v\n", set3)

	// Eliminando un elemento del set.
	delete(set3, 3)
	fmt.Printf("Conjunto después de eliminar un elemento: %#v\n", set3)

	// Comprobando si un elemento existe en el set.
	// El identificador en blanco _ se utiliza para descartar el valor.
	_, exists := set3[2]
	fmt.Printf("¿El elemento 2 existe en el conjunto? %v\n", exists) // Output: true

	// Obtener e imprimir todas las claves del set.
	for key := range set3 {
		fmt.Printf("Valor: %d\n", key)
	}

	// Ejemplo de uso de la función para convertir un slice en un set.
	slice := []int{1, 2, 2, 3, 4, 4}
	setFromSlice := sliceToSet(slice)
	fmt.Printf("Conjunto creado a partir del slice: %#v\n", setFromSlice)

	// Ejemplo de unión de dos conjuntos.
	setA := map[int]struct{}{1: {}, 2: {}}
	setB := map[int]struct{}{2: {}, 3: {}}
	unionSet := union(setA, setB)
	fmt.Printf("Unión de setA y setB: %#v\n", unionSet)
}

// sliceToSet convierte un slice en un conjunto.
func sliceToSet(slice []int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

// union toma dos conjuntos y devuelve su unión.
func union(setA, setB map[int]struct{}) map[int]struct{} {
	totalLen := len(setA) + len(setB)
	unionSet := make(map[int]struct{}, totalLen)

	// Añadiendo elementos del primer conjunto.
	for k := range setA {
		unionSet[k] = struct{}{}
	}

	// Añadiendo elementos del segundo conjunto.
	for k := range setB {
		unionSet[k] = struct{}{}
	}

	return unionSet
}
