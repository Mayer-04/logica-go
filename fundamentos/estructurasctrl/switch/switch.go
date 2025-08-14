package main

import "fmt"

/*
* Switch en Go

El switch es una alternativa más clara y ordenada al uso repetido de estructuras `if-else`.

* Características:
- Go facilita el uso de switch eliminando la necesidad de escribir `break` explícitos en cada `case`,
lo cual reduce el riesgo de errores y hace el código más claro.
- Cada caso (case) en un switch incluye implícitamente un `break`, por lo que una vez que se encuentra
una coincidencia, el programa no ejecutará los casos siguientes de manera automática.
- Evita la necesidad de múltiples if-else cuando se tiene una sola variable a comparar.
- La palabra clave `default` se utiliza para definir un caso predeterminado que se ejecutará si ningún
otro `case` coincide. Este bloque es opcional.
- En Go, el valor evaluado en el switch puede omitirse si se desea evaluar expresiones booleanas
directamente en cada `case`.
*/

/*
* Si deseas continuar ejecutando los casos siguientes después de encontrar una coincidencia
(similar al comportamiento por defecto en otros lenguajes),
Go te permite hacerlo utilizando la palabra clave `fallthrough`.
*/

func main() {
	// Switch simple: compara una variable contra varios valores posibles.
	opcion := 2
	switch opcion {
	case 1:
		fmt.Println("Opción 1: Ver perfil")
	case 2:
		fmt.Println("Opción 2: Editar perfil")
	case 3:
		fmt.Println("Opción 3: Cerrar sesión")
	default:
		fmt.Println("Opción no válida")
	}

	// Switch sin expresión: permite evaluar condiciones booleanas directamente.
	// Cuando no hay expresión después de switch, Go evalúa cada case como una condición booleana.
	edad := 18
	switch {
	case edad < 13:
		fmt.Println("Eres un niño")
	case edad < 18:
		fmt.Println("Eres un adolescente")
	default:
		fmt.Println("Eres un adulto")
	}

	//* Uso de etiquetas: Evita la repetición de valores constantes.
	// También se pueden utilizar etiquetas en switch para manejar diferentes casos con el mismo código.
	switch edad := edad; {
	case edad < 13:
		fmt.Println("Eres un niño")
	case edad < 18:
		fmt.Println("Eres un adolescente")
	default:
		fmt.Println("Eres un adulto")
	}

	// Type switch: permite ejecutar código según el tipo dinámico de una variable `interface{}`.
	// El switch de tipos permite verificar el tipo de una variable de interfaz vacía (interface{}).
	// Es útil cuando necesitamos hacer operaciones distintas dependiendo del tipo real de la variable.
	var x any = true
	switch x.(type) {
	case int:
		fmt.Println("Es un entero")
	case string:
		fmt.Println("Es una cadena")
	case float64:
		fmt.Println("Es un flotante")
	default:
		fmt.Println("Es de otro tipo") // En este caso: "Es de otro tipo"
	}

	//* Switch con la palabra clave `fallthrough`.
	// La palabra clave `fallthrough` fuerza la ejecución del siguiente caso, incluso si
	// la siguiente condición no se cumple. Este comportamiento es único en Go y puede ser útil
	// en casos en que se quiera que el flujo continúe a través de varios casos.
	switch 1 {
	case 1:
		fmt.Println("Play selected") // Imprime este primer `case`
		fallthrough
	case 2:
		fmt.Println("Settings selected") // Imprime este segundo `case` debido a `fallthrough`
		fallthrough
	case 3:
		fmt.Println("Help selected") // También imprime este tercer `case`
	default:
		fmt.Println("Invalid selection") // No se alcanza debido a `fallthrough`
	}

	//* Uso de break con etiquetas: En Go, un break dentro de un switch normalmente solo termina el caso actual.
	// Si estás dentro de un bucle y deseas salir del bucle basándote en una condición evaluada en un switch,
	// necesitarás usar etiquetas. En este ejemplo, `break loop` termina el bucle for y no solo el switch.
loop:
	for {
		switch 1 {
		case 1:
			break loop
		}
	}

	//* Switch con múltiples valores en un solo `case`.
	// Podemos utilizar múltiples valores dentro de un solo bloque de `case`.
	// En tal caso, el bloque `case` se ejecuta si la expresión coincide con cualquiera de los valores.
	dayOfWeek := "Sunday"
	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}

	//* Switch que como primer caso maneja un `default`.
	// Cuando colocas el default al principio, el código maneja los casos no válidos inmediatamente.
	// Útil cuando quieres hacer un manejo temprano de errores o condiciones no válidas.
	// Si el valor no es válido, el switch termina inmediatamente con un return y no evalúa el resto de los casos.
	prepararCafe := func(tipo string) {
		switch tipo {
		default:
			fmt.Println("Error: Tipo de café no disponible")
			return
		case "espresso":
			fmt.Println("Preparando un espresso")
		case "americano":
			fmt.Println("Preparando un americano")
		case "capuccino":
			fmt.Println("Preparando un capuccino")
		}
	}

	prepararCafe("espresso")  // Preparando un espresso
	prepararCafe("cafe")      // Error: Tipo de café no disponible
	prepararCafe("americano") // Preparando un americano
	prepararCafe("capuccino") // Preparando un americano
}
