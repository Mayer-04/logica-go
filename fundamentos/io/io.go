package main

import (
	"io"
)

/*
* Paquete io
El paquete `io` en Go proporciona interfaces y funciones básicas para realizar operaciones de entrada y salida
(Input/Output, de ahí el nombre "io"). Permite leer y escribir datos desde y hacia diferentes fuentes
(como archivos, conexiones de red, buffers en memoria, etc.), utilizando interfaz comunes.

* Conceptos clave:
- `io.Reader` y `io.Writer` son las interfaces más importantes para realizar operaciones de entrada y salida.
`Reader` se encarga de leer datos desde una fuente, y `Writer` se utiliza para escribir datos en un destino.

- El paquete también incluye utilidades adicionales, como `io.Copy`, `io.CopyN` y `io.WriteString`, que facilitan
la transferencia de datos entre diferentes fuentes y destinos.

* Consideraciones de eficiencia:
- En operaciones con grandes volúmenes de datos, como archivos grandes o respuestas HTTP, es recomendable no
usar `io.ReadAll`, ya que intenta cargar todo el contenido en memoria, lo que puede ser ineficiente. En su lugar,
se debe usar un `buffer` para leer los datos en bloques manejables (por ejemplo, un buffer de 4KB es un tamaño típico).

* Errores comunes:
- `io.EOF`: Este error se devuelve cuando no hay más datos que leer de un `Reader`. No debe considerarse
un error real, sino una señal de que la lectura ha llegado al final del flujo de datos.

* Interfaces clave del paquete io:
- `io.Reader`: Define un método `Read` que permite leer datos desde una fuente y devolver el número de bytes leídos
y un posible error. Todo lo que pueda proporcionar datos, como archivos, conexiones de red o buffers, puede implementar esta interfaz.

- `io.Writer`: Define un método `Write` que permite escribir datos en un destino. Implementan esta interfaz todas
las entidades que pueden recibir datos, como archivos, conexiones de red, etc.

- `io.Closer`: Proporciona un método `Close` que debe ser llamado cuando ya no se necesiten los recursos asociados
con un `Reader` o `Writer`, como cuando se termina de trabajar con un archivo o una conexión de red.

- `io.ReaderFrom`: Define un metodo `ReadFrom` que permite leer datos desde un `Reader` hasta EOF o error.
La función `io.Copy` usa este método si esta disponible.
*/

func main() {
	//* NOTE: Este código es simplemente explicativo y no tiene como salida un resultado real.

	// io.Discard es un `io.Writer` que no realiza ninguna acción.
	// Su función principal es descartar cualquier dato que se le escriba,
	// lo que significa que no almacena ni muestra esos datos en ninguna parte.
	discard := io.Discard

	// io.Copy copia datos de un `io.Reader` a un `io.Writer`.
	// Esto es útil para transferir datos directamente desde una fuente de lectura
	// (como un archivo o una conexión de red) a un destino (como otro archivo, un buffer o la salida estándar).
	io.Copy(io.Discard, nil)

	// io.CopyN copia exactamente `n` bytes de un io.Reader a un io.Writer. Si el `Reader` contiene menos de `n` bytes,
	// la operación devolverá un error `io.EOF` cuando no se pueda leer más.
	io.CopyN(io.Discard, nil, 0)

	// io.CopyBuffer funciona como `io.Copy`, pero permite especificar un buffer manualmente para mejorar el rendimiento
	// en ciertos escenarios donde la gestión manual de buffers es deseable.
	// Si el buffer es nil, utiliza uno por defecto.
	io.CopyBuffer(nil, nil, nil)

	// io.WriteString escribe una cadena de texto en un io.Writer.
	io.WriteString(discard, "Hola Mundo!\n")

	// io.ReadAll lee todo el contenido de un io.Reader en memoria y devuelve los datos como un slice de bytes.
	// Esto es conveniente para pequeñas cantidades de datos, pero para grandes volúmenes puede ser ineficiente
	// porque intenta cargar todo el contenido en memoria.
	io.ReadAll(nil)

	// io.LimitReader envuelve un io.Reader y limita la cantidad de datos que se pueden leer.
	// Esto es útil cuando solo se necesita leer una cantidad específica de datos, sin procesar todo el flujo.
	// El ejemplo devuelve un io.Reader que permite leer hasta `n` bytes de un Reader subyacente.
	io.LimitReader(nil, 0)
}
