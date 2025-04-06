package main

import (
	"context"
	"fmt"
	"time"
)

/*
* Paquete context en Go:
El paquete `context` es usado en Go para manejar la cancelación y establecer límites de tiempo
en operaciones concurrentes, como goroutines.
- Esencial para gestionar y coordinar tareas concurrentes.

* Propagación de contextos:
- Los contextos permiten compartir datos y gestionar la cancelación de tareas concurrentes.
- Puedes crear nuevos contextos a partir de otros para propagar la cancelación o los plazos de tiempo.
- Los contextos también sirven para evitar que las operaciones tarden más de lo necesario mediante plazos (deadlines)
o tiempos de espera (timeouts).
- Importante: No uses el contexto para almacenar datos grandes o sensibles.

* Usos principales:
1. Manejar cancelaciones de tareas concurrentes.
2. Establecer tiempos límite o de espera.
3. Pasar valores entre funciones de forma segura.
*/

// Definir un tipo específico para la clave, evitando colisiones en `context.WithValue`.
type requestIDKeyType string

const requestIDKey requestIDKeyType = "requestID"

func main() {
	//* `context.TODO()` se usa como un "marcador" cuando aún no estás seguro qué tipo de contexto utilizar.
	// Es útil durante el desarrollo, pero no debe usarse en producción.
	// Normalmente, se reemplaza por `context.Background()` o un contexto con características específicas.
	_ = context.TODO()

	// Crear un contexto raíz (padre).
	// Este contexto no puede ser cancelado y no tiene un plazo de tiempo asociado.
	ctx := context.Background()

	//* Crear un contexto derivado de `ctx` que se puede cancelar.
	// El contexto `ctxCancel` hereda de `ctx` y puede ser cancelado para liberar recursos.
	ctxCancel, cancel := context.WithCancel(ctx)

	//* `context.WithoutCancel(ctxCancel)` se añadio en la versión 1.21 de Go.
	// Permite crear una copia de un contexto padre que no se cancela cuando el contexto padre es cancelado.
	// No tiene un canal `Done`, no tiene una fecha límite y no propaga errores si el contexto padre es cancelado.
	noCancelCtx := context.WithoutCancel(ctxCancel)

	// `cancel()` se llama cuando la función `main` termine, para liberar los recursos asociados con `ctxCancel`.
	// Es una buena práctica llamar a `cancel` para evitar fugas de recursos.
	defer cancel()

	//* Crear un contexto derivado de `ctxCancel` con un tiempo de espera (timeout).
	// El contexto `ctxTimeout` se cancelará automáticamente después de 3 segundo o si se llama a `cancel()`.
	ctxTimeout, cancelTimeout := context.WithTimeout(ctxCancel, 3*time.Second)
	defer cancelTimeout() // Liberar los recursos asociados a `ctxTimeout`.

	// Simular una operación que puede respetar el tiempo de espera definido por el contexto.
	select {
	case <-time.After(2 * time.Second): // Simula una operación que dura 2 segundos.

		// Como el tiempo de espera es de 3 segundos y la operación de retraso es de 2 segundos,
		// la operación se completará antes del tiempo de espera.
		fmt.Println("Operación completada.")

		// Done nos devuelve un canal (chan struct{}) al que no es necesario pasar ningún dato a través del canal.
		// Se utiliza  para notificar cuando el contexto se cancela o se agota el tiempo.
	case <-ctxTimeout.Done():
		// Se ejecuta si el contexto se cancela o el tiempo se agota.
		fmt.Println("Operación cancelada o tiempo agotado:", ctxTimeout.Err())

	case <-noCancelCtx.Done():
		// Este bloque no se ejecutará porque `noCancelCtx` no se cancela.
		fmt.Println("El contexto sin cancelación ha sido cancelado")
	}

	// * ctxDeadline es un contexto derivado de `ctx`.
	// Crea un contexto con una fecha límite (deadline) específica que se alcanzará en 500 ms.
	deadline := time.Now().Add(500 * time.Millisecond)
	ctxDeadline, cancelDeadline := context.WithDeadline(ctx, deadline)
	defer cancelDeadline()

	//* Pasar valores a través de contextos:
	// Los valores pueden ser almacenados en un contexto y ser accesibles en funciones descendientes.
	// Esto es útil para pasar información como IDs de usuario, tokens de autenticación, etc.
	// No usar `string` literales como claves para evitar colisiones.
	ctxValue := context.WithValue(ctxDeadline, requestIDKey, "12345")

	// Simular operaciones concurrentes que usan los contextos anteriores
	go processRequest(ctxValue)   // Output: Operación completada con RequestID: 12345
	go processRequest(ctxTimeout) // Output: Operación completada en el request

	// Esperar 1 segundos para permitir que las goroutines terminen.
	time.Sleep(time.Second)
}

// processRequest Simula una función que procesa una solicitud respetando el contexto proporcionado.
func processRequest(ctx context.Context) {
	select {
	case <-time.After(400 * time.Millisecond): // Simula una operación que dura 400ms.
		// `Value()` obtiene el valor con la clave `requestIDKey` almacenado en el contexto, si existe.
		if requestID := ctx.Value(requestIDKey); requestID != nil {
			fmt.Println("Operación completada con RequestID:", requestID)
			return
		}
		// Si no hay valor en el contexto, indicar que se completó la operación.
		fmt.Println("Operación completada en el request.")
	case <-ctx.Done():
		// Se ejecuta si el contexto se cancela o se agota el tiempo.
		fmt.Println("Operación cancelada", ctx.Err())
	}
}
