package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * PAQUETE BUFIO: ENTRADA/SALIDA BUFFERIZADA
---------------------------------------------
El paquete `bufio` implementa funciones para entrada/salida bufferizada.
VENTAJAS:
- Mejora significativamente el rendimiento al reducir llamadas al sistema operativo
- Lee/escribe grandes bloques de datos y los almacena en un buffer interno
- Proporciona métodos convenientes para leer líneas, palabras, o caracteres individuales
- Maneja eficientemente diferentes formatos de entrada
COMPONENTES PRINCIPALES:
- Reader: Para operaciones de lectura bufferizada
- Writer: Para operaciones de escritura bufferizada
- Scanner: Para escanear texto de manera eficiente
*/

func main() {
	fmt.Println("=== EJEMPLOS DEL PAQUETE BUFIO ===")

	ejemploNewReader()
	ejemploReadString()
	ejemploReadBytes()
	ejemploReadLine()
	ejemploReadRune()
	ejemploPeek()
	ejemploNewScanner()
}

// ejemploNewReader demuestra cómo crear un nuevo Reader bufferizado
func ejemploNewReader() {
	fmt.Println("\n1. EJEMPLO DE bufio.NewReader")
	fmt.Println("------------------------------")

	/*
	 * bufio.NewReader(rd io.Reader) *bufio.Reader
	 *
	 * Crea un nuevo lector bufferizado que envuelve una fuente de entrada (io.Reader).
	 * - El buffer por defecto es de 4096 bytes
	 * - También existe bufio.NewReaderSize() para especificar un tamaño personalizado
	 */

	// Podemos crear un Reader desde distintas fuentes:
	// 1. Desde una cadena (usando strings.NewReader)
	strReader := bufio.NewReader(strings.NewReader("Esto es un ejemplo de lectura desde una cadena"))
	fmt.Println("Reader creado desde string")

	// 2. Desde la entrada estándar (os.Stdin)
	// Comentado para no bloquear la ejecución
	// stdinReader := bufio.NewReader(os.Stdin)
	fmt.Println("Reader creado desde os.Stdin (comentado)")

	// 3. Desde un archivo (requiere abrir el archivo primero)
	fmt.Println("Reader creado desde archivo (ejemplo conceptual)")

	// Ejemplo de uso básico
	char, _, _ := strReader.ReadRune()
	fmt.Printf("Primer carácter leído: '%c'\n", char)
}

// ejemploReadString demuestra el uso del método ReadString
func ejemploReadString() {
	fmt.Println("\n2. EJEMPLO DE ReadString")
	fmt.Println("------------------------")

	/*
	 * ReadString(delim byte) (string, error)
	 *
	 * Lee hasta encontrar el delimitador especificado (incluido) y devuelve lo leído como string.
	 * - Común para leer líneas usando '\n' como delimitador
	 * - Devuelve error io.EOF cuando llega al final del archivo/entrada
	 */

	texto := "Hola\nMundo\nBufio"
	reader := bufio.NewReader(strings.NewReader(texto))

	// Leer hasta encontrar el carácter '\n'
	linea1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Primera línea: %q\n", linea1) // Incluye el delimitador '\n'
	}

	// Leer otra línea
	linea2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Segunda línea: %q\n", linea2)
	}

	// Leer hasta el final (no hay más '\n', termina con EOF)
	linea3, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Error inesperado:", err)
	} else {
		fmt.Printf("Tercera línea (sin '\\n'): %q\n", linea3)
		if err == io.EOF {
			fmt.Println("Se alcanzó el final del texto (EOF)")
		}
	}
}

// ejemploReadBytes demuestra el uso del método ReadBytes
func ejemploReadBytes() {
	fmt.Println("\n3. EJEMPLO DE ReadBytes")
	fmt.Println("----------------------")

	/*
	 * ReadBytes(delim byte) ([]byte, error)
	 *
	 * Similar a ReadString pero devuelve un slice de bytes en lugar de un string.
	 * - Útil cuando necesitamos los datos en formato []byte
	 * - Funciona igual que ReadString en cuanto al delimitador
	 */

	texto := "123,456,789"
	reader := bufio.NewReader(strings.NewReader(texto))

	// Leer hasta la primera coma
	bytes1, err := reader.ReadBytes(',')
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Primer grupo como bytes: %v\n", bytes1)          // Vista de bytes
		fmt.Printf("Primer grupo como string: %q\n", string(bytes1)) // Convertido a string
	}

	// Leer hasta la segunda coma
	bytes2, err := reader.ReadBytes(',')
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Segundo grupo como string: %q\n", string(bytes2))
	}

	// Leer hasta el final (no hay más comas)
	bytes3, err := reader.ReadBytes(',')
	if err != io.EOF {
		fmt.Println("Error o no se esperaba EOF:", err)
	} else {
		fmt.Printf("Tercer grupo (sin coma, con EOF): %q\n", string(bytes3))
	}
}

// ejemploReadLine demuestra el uso del método ReadLine
func ejemploReadLine() {
	fmt.Println("\n4. EJEMPLO DE ReadLine")
	fmt.Println("---------------------")

	/*
	 * ReadLine() (line []byte, isPrefix bool, err error)
	 *
	 * Lee una línea PERO tiene limitaciones importantes:
	 * - No incluye el delimitador '\n' en el resultado
	 * - Si la línea es más larga que el buffer, devuelve isPrefix=true
	 * - En ese caso, se necesitan llamadas adicionales para leer la línea completa
	 * - Es un método de bajo nivel; generalmente es mejor usar Scanner o ReadString
	 */

	// Texto con líneas cortas y largas
	texto := "Línea corta\nEsta es una línea muy larga que podría exceder el tamaño del buffer interno si creamos un buffer pequeño\nOtra línea"

	// Creamos un reader con buffer pequeño para demostrar isPrefix
	reader := bufio.NewReaderSize(strings.NewReader(texto), 20)

	// Leemos líneas
	var completa []byte
	for {
		linea, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		if isPrefix {
			fmt.Println("Fragmento de línea:", string(linea))
			completa = append(completa, linea...)
		} else {
			if len(completa) > 0 {
				// Completamos la línea anterior
				completa = append(completa, linea...)
				fmt.Println("Línea completa:", string(completa))
				completa = nil
			} else {
				fmt.Println("Línea completa (de una vez):", string(linea))
			}
		}
	}
}

// ejemploReadRune demuestra el uso del método ReadRune
func ejemploReadRune() {
	fmt.Println("\n5. EJEMPLO DE ReadRune")
	fmt.Println("---------------------")

	/*
	 * ReadRune() (r rune, size int, err error)
	 *
	 * Lee un único carácter como tipo rune (punto de código Unicode).
	 * - Maneja correctamente caracteres Unicode multibyte
	 * - Devuelve el tamaño en bytes del rune
	 * - Ideal para procesar texto caracter por caracter
	 */

	// Texto con caracteres especiales Unicode
	texto := "Hola 世界" // "Hola Mundo" con "Mundo" en chino
	reader := bufio.NewReader(strings.NewReader(texto))

	// Leemos caracter por caracter
	for {
		r, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fmt.Printf("Rune: '%c', Código Unicode: %U, Tamaño en bytes: %d\n", r, r, size)
	}
}

// ejemploPeek demuestra el uso del método Peek
func ejemploPeek() {
	fmt.Println("\n7. EJEMPLO DE Peek")
	fmt.Println("----------------")

	/*
	 * Peek(n int) ([]byte, error)
	 *
	 * Devuelve los próximos n bytes sin avanzar el puntero de lectura.
	 * - No consume los datos, solo permite "espiar" lo que viene después
	 * - Útil para determinar qué hacer antes de consumir los datos
	 * - Devuelve ErrBufferFull si n es mayor que el tamaño del buffer
	 */

	texto := "ABC123XYZ"
	reader := bufio.NewReader(strings.NewReader(texto))

	// Miramos los primeros 3 bytes sin consumirlos
	peek, err := reader.Peek(3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Peek (3 bytes): %q\n", peek)
	}

	// Leemos 1 byte (esto avanza el puntero)
	b, _ := reader.ReadByte()
	fmt.Printf("Byte leído: %c\n", b)

	// Miramos de nuevo (ahora vemos desde el segundo caracter)
	peek, _ = reader.Peek(3)
	fmt.Printf("Peek después de leer un byte: %q\n", peek)

	// Intentamos mirar más allá del tamaño disponible
	peek, err = reader.Peek(100)
	if err != nil {
		fmt.Printf("Error al hacer Peek de 100 bytes: %v\n", err)
	}
}

// ejemploNewScanner demuestra el uso de bufio.Scanner.
func ejemploNewScanner() {
	fmt.Println("\n8. EJEMPLO DE bufio.Scanner")
	fmt.Println("-------------------------")

	/*
	 * bufio.Scanner
	 *
	 * Scanner proporciona una interfaz conveniente para leer texto delimitado:
	 * - Por defecto lee líneas (delimitadas por '\n')
	 * - Puede configurarse para leer palabras, tokens u otros delimitadores personalizados
	 * - Muy útil para procesar archivos línea por línea
	 * - Más simple de usar que Reader para ciertos casos
	 */

	texto := "Primera línea\nSegunda línea\nTercera línea"
	scanner := bufio.NewScanner(strings.NewReader(texto))

	// Escanear línea por línea (comportamiento por defecto)
	fmt.Println("Escaneando líneas:")
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("  %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear:", err)
	}

	// Escanear por palabras
	fmt.Println("\nEscaneando palabras:")
	scanner = bufio.NewScanner(strings.NewReader(texto))
	scanner.Split(bufio.ScanWords)

	wordCount := 1
	for scanner.Scan() {
		fmt.Printf("  Palabra %d: %s\n", wordCount, scanner.Text())
		wordCount++
		// Limitamos a 5 palabras para no hacer el ejemplo muy largo
		if wordCount > 5 {
			fmt.Println("  ... (y más palabras)")
			break
		}
	}
}
