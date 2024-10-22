package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
* Manejo de Errores en Go
En Go, los errores se tratan de manera `explícita` en lugar de usar excepciones, como en otros lenguajes.
Esto proporciona un manejo de errores más claro y controlado. Algunos puntos clave son:
- Los errores se devuelven como el último valor de retorno en las funciones.
- Para manejar los errores, generalmente usamos la estructura de control `if`.
- El valor por defecto de un error es `nil`, lo que indica que no hubo error.
- No se deben ignorar los errores devueltos, incluso por operaciones con `defer`,
si deseas ignorar un error, utiliza el operador blank `_`.

* Es un error común manejar un mismo error varias veces. Debemos recordar que:
- Manejar un error significa registrarlo (log) o devolverlo (return), no ambas cosas.
- Si manejas un error, es importante hacerlo solo una vez.
- Para envolver un error y mantener su contexto original, usa `fmt.Errorf` con la directiva `%w`,
permitiendo comparar con `errors.Is` o extraer el error con `errors.As`.
- Si no deseas exponer el error original a quien llama a la función, usa `%v` en lugar de `%w`.
*/

// La función 'New()' del paquete "errors" crea un nuevo valor de error con un mensaje personalizado.
var errorDivide = errors.New("cannot divide by zero")

func main() {
	// Intentamos convertir una cadena a entero usando 'strconv.Atoi'.
	value, err := strconv.Atoi("10")
	if err != nil {
		// Imprimimos un mensaje de error y nos salimos del programa.
		fmt.Printf("cannot convert %q to int: %v\n", "10", err)
		return
	}

	// Si no hay error, imprimimos el valor convertido.
	fmt.Println("Valor convertido:", value)

	// Intentamos dividir dos números. Si el divisor es cero, recibiremos un error.
	result, err := divide(6, 0)
	// Comprobamos si hubo un error.
	if err != nil {
		// Comprobamos si el error recibido es el específico de "división por cero".
		if errors.Is(err, errorDivide) {
			fmt.Println("Error en la división:", err)
		}
		return
	}
	// Imprimimos el resultado.
	fmt.Println("Resultado:", result)
}

// Función que realiza una división y devuelve un error si b es igual a 0.
func divide(a, b float64) (float64, error) {
	// Verificamos si el divisor es cero y retornamos un error si es el caso.
	if b == 0 {
		// Devolvemos un valor de error personalizado.
		return 0, errorDivide
	}
	// Retornamos el resultado de la división y nil, indicando que no hubo error.
	return a / b, nil
}
