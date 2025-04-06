package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

/*
* Manejo de Errores en Go

📖 FRASE: Los errores son valores.

En Go, cualquier tipo que implemente la interfaz `error` se considera un error.

Los errores se tratan de manera `explícita` en lugar de usar excepciones, como en otros lenguajes.
Lo que hace que el código sea más claro y controlado.

✅ Un buen manejo de errores hace que tu código sea más robusto y fácil de depurar.

* Puntos clave:
- Las funciones devuelven los errores como `último valor de retorno`.
- Para manejar los errores, generalmente usamos la estructura de control `if`.
- Si no hay error, el valor del error será `nil`.
- No se deben ignorar los errores devueltos, incluso en operaciones con `defer`.
- Si deseas ignorar un error, utiliza el operador blank `_` (No recomendado).

* NOTA:
Como los errores son solo interfaces, puedes definir tus `propios tipos de error` implementando `error`.

* Es un error común manejar un mismo error varias veces. Debemos recordar que:
- Manejar un error significa registrarlo (log) o devolverlo (return), no ambas cosas.
- Si manejas un error, es importante hacerlo solo una vez.
- Para envolver un error y mantener su contexto original, usa `fmt.Errorf` con la directiva `%w`,
permitiendo comparar con `errors.Is` o extraer el error con `errors.As`.
- Si no deseas exponer el error original a quien llama a la función, usa `%v` en lugar de `%w`.

* Errores comunes al manejar errores:

1️⃣ Evita manejar el mismo error varias veces.
   - Manejar un error significa `registrarlo (log)` o `devolverlo (return)`, pero no ambas cosas.
   - Si decides manejar un error, hazlo solo una vez.

2️⃣ Envolver errores correctamente.
   - Usa `fmt.Errorf` con `%w` para `mantener el contexto original` del error.
   - Esto permite comparar errores con `errors.Is` o extraer detalles con `errors.As`.
   - Si no quieres exponer el error original, usa `%v` en lugar de `%w`.
*/

// La función 'New()' del paquete "errors" crea un nuevo valor de error con un mensaje personalizado.
var errorDivide = errors.New("cannot divide by zero")

func main() {
	// Intentamos convertir una cadena a entero usando 'strconv.Atoi'.
	value, err := strconv.Atoi("10")
	if err != nil {
		// Usamos log para imprimir el error y salimos del programa.
		log.Printf("cannot convert %q to int: %v\n", "10", err)
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
			log.Println("Error en la división:", err)
		}
		return
	}
	// Imprimimos el resultado.
	fmt.Println("Resultado:", result)

	// Ejemplo de cómo envolver un error con fmt.Errorf y %w
	_, err = sqrt(-10)
	if err != nil {
		if errors.Is(err, ErrNegativeNumber) {
			log.Println("Error en la raíz cuadrada:", err)
		}
		return
	}
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

// Ejemplo de cómo envolver un error con fmt.Errorf y %w.
var ErrNegativeNumber = errors.New("negative number")

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("sqrt of %v: %w", x, ErrNegativeNumber)
	}
	// Implementación de la raíz cuadrada (simplificada)
	return x * x, nil
}

// Función que maneja múltiples errores y devuelve un error compuesto.
func multipleOperations(a, b float64) (float64, error) {
	result1, err := divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("operation failed: %w", err)
	}

	result2, err := sqrt(result1)
	if err != nil {
		return 0, fmt.Errorf("operation failed: %w", err)
	}

	return result2, nil
}
