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
	fmt.Printf("Cadena: %q\n", cadena) // "Hola Mundo!"

	// Usando `%d` para imprimir números enteros.
	edad := 23
	fmt.Printf("Mi edad es: %d\n", edad)

	// Usando `%f` es para números con decimales (flotantes).
	estatura := 1.75
	fmt.Printf("Mi estatura es: %f\n", estatura)

	// Usando `%t` para imprimir valores booleanos.
	vivo := true
	fmt.Printf("Vivo: %t\n", vivo)

	// Usando `%T` para imprimir el tipo de dato de una variable.
	fmt.Printf("Tipo: %T\n", nombre)

	// Usando `%x` para imprimir un valor en formato hexadecimal en letras minúsculas.
	hexadecimal := 0xff
	fmt.Printf("Hexadecimal: %x\n", hexadecimal)

	// Usando `%o` para imprimir un valor en formato octal.
	octal := 045
	fmt.Printf("Octal: %o\n", octal)

	// Usando `%b` para imprimir un valor en formato binario.
	binario := 0b1010
	fmt.Printf("Binario: %b\n", binario)

	//* Entrada de datos del usuario desde la consola:

	// `fmt.Scan()` Lee múltiples valores separados por espacios en una sola línea.
	// Puedes ingresar varios valores en una sola línea y los separará y los asignará a las variables correspondientes.
	var name string
	var age int
	fmt.Print("Ingrese su nombre y edad: ")
	fmt.Scan(&name, &age)
	fmt.Printf("Hola, %s. Tienes %d años.\n", name, age)

	// `fmt.Scanln()` Lee hasta un salto de línea (Enter), capturando toda la entrada.
	// Si intentas ingresar varios valores en una sola línea con `fmt.Scanln()`, solo se tomará el primer valor.
	var sentence string
	fmt.Print("Ingrese una oración: ")
	fmt.Scanln(&sentence)
	fmt.Printf("La oración ingresada es: %s\n", sentence)
}
