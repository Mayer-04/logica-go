package main

import (
	"context"
	"fmt"
	"time"
)

/*
* Paquete context en Go

El paquete `context` se usa para controlar el tiempo de vida de operaciones concurrentes (goroutines).
Es útil para cancelar tareas, establecer límites de tiempo (timeouts)
y pasar información entre funciones de forma segura y controlada.

Su uso es común en servidores, APIs y programas concurrentes.

* Propagación del contexto:

- Puedes crear un nuevo contexto a partir de otro (padre → hijo).
- Esto permite que la cancelación o el vencimiento (timeout/deadline) del contexto padre
  se propague automáticamente a sus contextos hijos.
- Así puedes evitar que tareas corran más tiempo del necesario.

Ejemplo: si un cliente cierra su conexión, puedes cancelar todas las tareas asociadas a esa solicitud.

* Principales usos del paquete context:

1. Cancelar tareas concurrentes cuando ya no son necesarias.
2. Establecer un tiempo máximo para que una tarea se complete (timeout o deadline).
3. Pasar valores pequeños y de solo lectura entre funciones (como IDs de solicitud, tokens, etc.).

✅ Buenas prácticas:
- No uses context para pasar grandes estructuras de datos ni datos sensibles.
- No almacenes el contexto dentro de estructuras, solo pásalo como parámetro explícito.
- Siempre pasa el contexto como primer parámetro en funciones: `func(ctx context.Context, ...)`.
*/

// Definir un tipo específico para la clave, evitando colisiones en `context.WithValue`.
type requestIDKeyType string

const requestIDKey requestIDKeyType = "requestID"

func main() {
	//* `context.TODO()` se usa como un "marcador" cuando aún no estás seguro qué tipo de contexto utilizar.
	// Es útil durante el desarrollo, pero no debe usarse en producción.
	// Normalmente, se reemplaza por `context.Background()` o un contexto con características específicas.
	// Es com decir: "Aquí necesito un contexto, pero aún no sé qué tipo usar".
	_ = context.TODO()

	//* `context.Background()` crea un contexto raíz (padre).
	// Este contexto no puede ser cancelado y no tiene un plazo de tiempo asociado.
	// Normalmente partes de context.Background() y luego creas contextos derivados con cancelación, timeout, etc.
	ctx := context.Background()

	//* Crear un contexto derivado de `ctx` que se puede cancelar.
	// El contexto `ctxCancel` hereda de `ctx` y puede ser cancelado para liberar recursos.
	ctxCancel, cancel := context.WithCancel(ctx)

	//* `context.WithoutCancel()` se añadio en la versión 1.21 de Go.
	// Permite crear una copia de un contexto padre que no se cancela cuando el contexto padre es cancelado.
	// No tiene un canal `Done`, no tiene una fecha límite y no propaga errores si el contexto padre es cancelado.
	noCancelCtx := context.WithoutCancel(ctx)

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
