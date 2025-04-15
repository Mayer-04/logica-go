package main

import "fmt"

/*
* Funciones de impresión en Go
Go proporciona varias funciones para mostrar información en la consola:

- `print()`: Esta es una función nativa que imprime información sin salto de línea.
No es muy común en aplicaciones más complejas, pero puede ser útil para cosas sencillas.

* Paquete fmt: format
Este paquete es muy útil porque nos permite mostrar, formatear y leer información de manera más controlada.

- Formatear: Organizar y presentar la información de una manera específica o legible.
- `fmt.Print()`: Imprime un valor o información en la consola sin un salto de línea.
- `fmt.Println()`: Imprime un valor o información en la consola y agrega un salto de línea al final.
- `fmt.Printf()`: Permite imprimir información formateada usando "especificadores de formato".
- Los especificadores de formato (`%`) son símbolos que indican cómo queremos que se muestren los valores.
  Ejemplos: `%v` (genérico), `%s` (cadena), `%d` (entero).
- `fmt.Scan()` y `fmt.Scanln()`: Se utilizan para leer datos desde la consola.
La diferencia es que `Scan` toma múltiples valores en una sola línea, y `Scanln` solo uno por línea.
*/

type car struct {
	Modelo string
	Año    int
}

func main() {
	// La función `print()` simplemente imprime sin agregar un salto de línea al final.
	// No es muy usada porque `fmt` ofrece más opciones.
	print("Mi primer programa en Go")

	// `fmt.Println()` imprime un valor con un salto de línea al final.
	fmt.Println("Hola Mundo!")

	// `fmt.Printf()` imprime una cadena con un formato específico.
	fmt.Printf("Hola Mundo con formato\n")

	// Usando `%v` para imprimir el valor de una variable de manera genérica, sirve para cualquier tipo de valor.
	nombre := "Mayer"
	fmt.Printf("Mi nombre es: %v\n", nombre)

	// Usando `%s` para imprimir una cadena de texto.
	fmt.Printf("Mi nombre es: %s\n", nombre)

	// Usando `%q` para imprimir una cadena entre comillas.
	cadena := "Hola Mundo!"
	fmt.Printf("Cadena: %q\n", cadena) // Output: "Hola Mundo!"

	// Usando `%d` para imprimir números enteros.
	edad := 23
	fmt.Printf("Mi edad es: %d\n", edad)

	// Usando `%f` es para números con decimales (flotantes).
	// El número se imprime con 6 decimales por defecto, incluso si el valor original no los tiene.
	estatura := 1.75
	fmt.Printf("Mi estatura es: %f\n", estatura) // Output: Mi estatura es: 1.750000

	// Usando `%.3f` para imprimir el número de decimales después del punto.
	// El número se imprime solo con 3 decimales.
	flotante := 1.75554545
	fmt.Printf("Número flotante: %.3f\n", flotante) // Output: Número flotante: 1.756

	// Usando `%t` para imprimir valores booleanos.
	vivo := true
	fmt.Printf("Vivo: %t\n", vivo)

	// Usando `%T` para imprimir el tipo de dato de una variable.
	fmt.Printf("Tipo: %T\n", nombre)

	// Usando `%c` para imprimir valores de caracteres.
	// Funciona con valores de tipo `rune` y `byte`.
	letra := 'B'
	codigoASCII := 67
	caracterUnicode := '☀'

	fmt.Printf("Letra: %c\n", letra)                              // Output: B
	fmt.Printf("Código ASCII %d: %c\n", codigoASCII, codigoASCII) // Output: C
	fmt.Printf("Carácter Unicode: %c\n", caracterUnicode)         // Output: ☀

	// Usando `%x` para imprimir un valor en formato hexadecimal en letras minúsculas.
	hexadecimal := 0xff
	fmt.Printf("Hexadecimal: %x\n", hexadecimal) // Output: ff

	// Usando `%o` para imprimir un valor en formato octal.
	octal := 045
	fmt.Printf("Octal: %o\n", octal) // Output: 45

	// Usando `%b` para imprimir un valor en formato binario.
	binario := 0b1010
	fmt.Printf("Binario: %b\n", binario) // Output: 1010

	// Usando `%p` para imprimir la dirección de memoria de una variable.
	puntero := &edad
	fmt.Printf("Dirección de memoria: %p\n", puntero) // Output: 0xc00000a0e8

	// Usando `%v`, `%+v` y `%#v` para imprimir estructuras.
	toyota := car{"Toyota Corolla", 2019}
	fmt.Printf("Estructura toyota: %v\n", toyota)  // Output: {Toyota Corolla 2019}
	fmt.Printf("Estructura toyota: %+v\n", toyota) // Output: {Modelo:Toyota Corolla Año:2019}
	fmt.Printf("Estructura toyota: %#v\n", toyota) // Output: main.Car{Modelo:"Toyota Corolla", Año:2019}

	//* ENTRADA DE DATOS DESDE LA CONSOLA:

	/*
		* Ejemplo 1: Lectura de múltiples valores con `fmt.Scan`.
		`fmt.Scan` lee valores de la entrada estándar hasta encontrar suficientes elementos
		para llenar todos los argumentos proporcionados. Los valores deben estar separados
		por espacios en blanco (espacios, tabs o saltos de línea).
		Uso típico:
			Entrada: "Juan 25"
			Resultado: name="Juan", age=25
		NOTA: Si el usuario ingresa más valores de los esperados, los adicionales
		se descartan automáticamente.
	*/
	var name string
	var age int
	fmt.Print("Ingrese su nombre y edad (separados por espacio): ")
	fmt.Scan(&name, &age)
	fmt.Printf("Hola, %s. Tienes %d años.\n", name, age)

	/*
		* Ejemplo 2: Lectura de línea completa con `fmt.Scanln`.

		 `fmt.Scanln` lee la entrada hasta encontrar un salto de línea ('\n').
		 Es ideal para:
		 - Capturar texto que puede contener espacios
		 - Asegurar que la lectura se detenga al final de la línea

		 Diferencias clave con `fmt.Scan`:
		 - Se detiene al encontrar un salto de línea
		 - Preserva espacios en blanco dentro del texto
		 - No continúa leyendo en líneas adicionales

		 Ejemplo de uso:
		  Entrada: "Hola mundo desde Go"
		  Resultado: sentence="Hola mundo desde Go"
	*/
	var sentence string
	fmt.Print("Ingrese una oración: ")
	fmt.Scanln(&sentence)
	fmt.Printf("La oración ingresada es: %s\n", sentence)
}
