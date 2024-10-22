package main

import "fmt"

/*
* Switch: Conmutador-Selector
Una alternativa más limpia y eficiente a múltiples if-else.
- En Go, cada caso (case) en un switch incluye implícitamente un `break`,
lo que significa que una vez que se encuentra una coincidencia, el programa no ejecutará los casos siguientes.
- Go facilita el trabajo con switch al eliminar la necesidad de break explícitos en cada caso.
- Evita la necesidad de múltiples if-else cuando se tiene una sola variable a comparar.
- La palabra clave `default` se utiliza para indicar que el caso predeterminado se debe ejecutar
cuando no se encuentra una coincidencia.
*/

/*
* Si deseas continuar ejecutando los casos siguientes después de encontrar una coincidencia
(similar al comportamiento por defecto en otros lenguajes),
Go te permite hacerlo utilizando la palabra clave `fallthrough`.
*/

func main() {
	// Declarando una estructura switch con múltiples casos.
	edad := 18
	switch {
	case edad < 13:
		fmt.Println("Eres un niño")
	case edad < 18:
		fmt.Println("Eres un adolescente")
	default:
		fmt.Println("Eres un adulto")
	}

	// * Uso de etiquetas: evita la repetición de valores constantes.
	// También se pueden utilizar etiquetas en switch para manejar diferentes casos con el mismo código.
	switch edad := edad; {
	case edad < 13:
		fmt.Println("Eres un niño")
	case edad < 18:
		fmt.Println("Eres un adolescente")
	default:
		fmt.Println("Eres un adulto")
	}

	// * Switch de Tipos (Type Switches).
	// Se puede utilizar para el caso de que la variable sea de otro tipo.
	var x interface{} = true
	switch x.(type) {
	case int:
		fmt.Println("Es un entero")
	case float64:
		fmt.Println("Es un flotante")
	case string:
		fmt.Println("Es una cadena")
	default:
		fmt.Println("Es de otro tipo") // Se imprime 'Es de otro tipo'
	}

	//* Switch con la palabra clave `fallthrough`.
	// La palabra clave 'fallthrough' indica que el 'case' actual debe ser evaluado,
	// incluso si se cumple una condición posterior.
	switch 1 {
	case 1:
		fmt.Println("Play selected") // Se imprime el primer 'case'
		fallthrough
	case 2:
		fmt.Println("Settings selected") // Se imprime el segundo 'case'
		fallthrough
	case 3:
		fmt.Println("Help selected") // Se imprime el tercer 'case'
	default:
		fmt.Println("Invalid selection")
	}

	//* Uso de break con etiquetas - En Go, un break dentro de un switch normalmente solo termina el caso actual.
	// Si estás dentro de un bucle y deseas salir del bucle basándote en una condición evaluada en un switch,
	// necesitarás usar etiquetas. En este ejemplo, break loop termina el bucle for y no solo el switch.
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

	//* Switch true o Switch condicional.
	// No se coloca ninguna expresión después de switch, lo que implica que cada condición en los `case`,
	// es evaluada de manera individual como una expresión booleana.
	numberOfDays := 28
	// La expresión 'switch true' es equivalente a 'switch true {  }'.
	switch {
	case numberOfDays == 28:
		fmt.Println("It's February")
	default:
		fmt.Println("Not February")
	}

	//* Switch que como primer caso maneja un `default`.
	// Cuando colocas el default al principio, el código maneja los casos no válidos inmediatamente.
	// Útil cuando quieres hacer un manejo temprano de errores.
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

	prepararCafe("espresso")  // Output: Preparando un espresso
	prepararCafe("cafe")      // Output: Error: Tipo de café no disponible
	prepararCafe("americano") // Output: Preparando un americano
	prepararCafe("capuccino") // Output: Preparando un capuccino
}
