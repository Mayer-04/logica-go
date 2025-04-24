package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/*
* Funciones en Go
Go soporta aspectos de programación funcional, tratando funciones como "ciudadanos de primera clase".
- Las funciones pueden ser asignadas a variables, pasadas como argumentos y retornadas desde otras funciones.

* IMPORTANTE: Los parámetros de funciones en Go son pasados por valor, no por referencia.
Se trabaja con una copia del valor, no con el valor original.
- Para tipos como int, float, bool, string, array y struct, pasar por valor significa copiar todo el contenido.
- Los tipos que parecen pasarse por referencia (slices, maps, channels, etc.) en realidad se pasan por valor,
pero el valor que se pasa contiene una referencia a los datos subyacentes.
- Para slices, maps, channels, y pointers, pasar por valor significa copiar la referencia,
no los datos a los que apunta.

* Conceptos clave:
1. Funciones de orden superior: Funciones que aceptan otras funciones como argumentos o las retornan.
2. Closures: Funciones anónimas que pueden acceder y modificar variables de su ámbito externo.
3. Funciones variádicas: Aceptan un número variable de argumentos.
4. Funciones diferidas (defer): Se ejecutan al final de la función que las contiene.
5. Funciones anónimas: Funciones sin nombre, útiles para definiciones inmediatas.
6. Recursión: Funciones que se llaman a sí mismas.
7. Manejo de errores: Funciones que retornan errores como segundo valor de retorno.
8. Métodos: Funciones asociadas a un tipo específico mediante un receptor.
9. Interfaces: Conjuntos de firmas de métodos (que son funciones) que pueden ser implementados por tipos.
10. Panics y Recovery: Manejo de situaciones excepcionales en Go.
*/

// * Métodos y tipos: los métodos son funciones asociadas a un tipo específico.
// * Un método es simplemente una función con un argumento receptor especial.

// Definición de una estructura para demostrar métodos.
type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

// Este es un método de la estructura `Persona`.
// El "método" NombreCompleto es realmente una función con un "receptor" (p *Persona)
// que indica que esta función pertenece al tipo `Persona`.
func (p *Persona) NombreCompleto() string {
	return p.Nombre + " " + p.Apellido
}

// * Interfaces: Una interfaz es un conjunto de firmas de métodos (funciones).
// * Los tipos que implementan todos los métodos de una interfaz, implementan implícitamente la interfaz.

// Definición de una interfaz simple.
type Saludador interface {
	// Saludar es una función que debe ser implementada por los tipos
	// que quieran cumplir con esta interfaz.
	Saludar() string
}

// El tipo `Persona` implementa la interfaz `Saludador`.
// Este es un método (función con receptor) que hace que `Persona` implemente `Saludador`.
func (p Persona) Saludar() string {
	return fmt.Sprintf("¡Hola! Soy %s %s", p.Nombre, p.Apellido)
}

// Función que recibe cualquier tipo que implemente la interfaz `Saludador`.
func SaludarEntidad(s Saludador) {
	fmt.Println(s.Saludar())
}

func main() {
	// Llamada a la función saludar con dos argumentos de tipo string.
	saludar("Mayer", "Chaves")

	// Ejemplo con función por referencia.
	num := 5
	increment(&num)                             // Obtenemos la dirección en memoria de la variable "num".
	fmt.Println("Después de incrementar:", num) // El valor de num se cambia directamente

	//* RETURN: Llamada a una función con retorno de un solo valor.
	result := retornar(1, 3)
	fmt.Println("Resultado de la suma:", result)

	// Múltiples valores de retorno.
	// Podemos ignorar u omitir algún valor utilizando el operador blank (_).
	lower, upper := multiplesRetornos("Mayer")
	fmt.Println("Texto en minúscula:", lower)
	fmt.Println("Texto en mayúscula:", upper)

	//* Función variádica: Acepta un número indefinido de argumentos.
	fmt.Println("Suma variádica:", sumaVariadica(1, 2, 3, 4, 5))

	// Pasando un slice a una función variádica con el operador "..."
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Println("Suma variádica con slice:", sumaVariadica(numeros...))

	// Llamando a Función anonima.
	fmt.Println("Llamando a función anónima preestablecida:")
	anonima()

	// Creando una función anonima.
	// Puede ser definida en el lugar donde se necesite en lugar de ser definida globalmente.
	fmt.Println("Función anónima asignada a variable:")
	salud := func(name string) {
		fmt.Printf("Hola, %s!\n", name)
	}
	salud("Robert")

	//* Función anonima autoinvocada.
	fmt.Println("Función anónima autoinvocada:")
	func(name string) {
		fmt.Println("Hello", name)
	}("mayer")

	// Segunda manera de declarar una función anonima con la declaración de variable corta.
	fmt.Println("Otra forma de función anónima:")
	ejemplo := func() {
		fmt.Println("Soy una función anonima corta")
	}
	ejemplo()

	// Uso de funciones diferidas (defer).
	// La función diferida se ejecutará al final de la función main.
	defer diferida()

	// Ejemplo avanzado de defer con múltiples llamadas (se ejecutan en orden LIFO)
	fmt.Println("Demostración de múltiples defer (se ejecutarán en orden inverso al terminar main):")
	for i := range 3 {
		defer fmt.Printf("Defer #%d\n", i)
	}

	// Llamada a una función recursiva.
	fmt.Println("Factorial de 5:", factorial(5))

	//* Closure: función que captura variables del entorno.
	fmt.Println("Demostración de closure (contador):")
	counter := createCounter()
	fmt.Println("Counter:", counter()) // Output: Counter: 1
	fmt.Println("Counter:", counter()) // Output: Counter: 2

	// Nuevo ejemplo de closure - generador de multiplicador.
	fmt.Println("Closure como generador de funciones:")
	duplicador := createMultiplier(2)
	triplicador := createMultiplier(3)
	fmt.Println("5 duplicado:", duplicador(5))   // Output: 10
	fmt.Println("5 triplicado:", triplicador(5)) // Output: 15

	// Uso de un tipo (type) de función.
	fmt.Println("Uso de tipo de función:")
	var f tipoFuncion = saludar
	f("Mayer", "Andres")

	// Manejo de errores en funciones.
	fmt.Println("Manejo de errores en funciones:")
	resultado, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado de división:", resultado)
	}

	// Provocando un error.
	resultado, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado de división:", resultado)
	}

	// Funciones de orden superior.
	fmt.Println("Funciones de orden superior:")
	numeros2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Números originales:", numeros2)

	// Aplicar una función a cada elemento.
	numerosTransformados := aplicarATodos(numeros2, func(n int) int {
		return n * n
	})
	fmt.Println("Números al cuadrado:", numerosTransformados)

	// Filtrar elementos.
	numerosPares := filtrar(numeros2, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Números pares:", numerosPares)

	// Métodos y tipos.
	fmt.Println("Métodos como funciones con receptor:")
	p := Persona{
		Nombre:   "María",
		Apellido: "González",
		Edad:     28,
	}

	fmt.Println("Nombre completo:", p.NombreCompleto())

	// Interfaces.
	fmt.Println("Interfaces (conjuntos de firmas de métodos):")
	// La persona implementa la interfaz `Saludador`.
	SaludarEntidad(p)

	// Panic y Recovery.
	fmt.Println("Demonstración de panic y recovery:")
	demostracionPanicRecovery()
	fmt.Println("La función main continúa después de recuperarse del panic")

	// Init function.
	fmt.Println("La función init() se ejecutó al inicio, antes que main()")

	// Funciones con timeouts.
	fmt.Println("Función con timeout:")
	resultado2, timeout := operacionConTimeout(2)
	if timeout {
		fmt.Println("La operación tardó demasiado")
	} else {
		fmt.Println("Resultado:", resultado2)
	}

}

// Se ejecuta automáticamente antes de main().
func init() {
	fmt.Println("Inicializando... Esta función init() se ejecuta automáticamente al inicio")
}

//* CARACTERÍSTICAS DE LAS FUNCIONES EN GO.

// Función con parámetros agrupados del mismo tipo.
// Si tenemos varias funciones con parámetros del mismo tipo, podemos agruparlos en un mismo parámetro.
func saludar(name, lastname string) {
	fmt.Printf("Nombre %s con apellido %s", name, lastname)
}

// Función que trabaja con un puntero.
// Para trabajar con parámetros por referencia se deben utilizar "Punteros".
func increment(x *int) {
	// Operador desreferenciación: Accede al valor almacenado en la dirección de memoria.
	*x++ // Incrementa el valor en la dirección de memoria apuntada por x.
}

// Función con valor de retorno: return.
func retornar(num1, num2 int) int {
	return num1 + num2
}

// * Funciones con múltiples retornos.
// * Se debe incluir entre paréntesis los múltiples valores que se quieren retornar.
// Es posible nombrar los parámetros de retorno, aunque no es recomendado en funciones complejas.
func multiplesRetornos(text string) (lower, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)
	return // Retorno implícito
}

// * Funciones variádicas.
// Permiten pasar un número indefinido de argumentos del mismo tipo.
// Se anteponen tres puntos (...) al tipo de dato para indicar que es una función variádica.
func sumaVariadica(nums ...int) int {
	// El parámetro `nums` se convierte en un "slice" de enteros.
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// * Funciones anónimas.
// No tienen un nombre de función.
// Las funciones anónimas pueden asignarse a variables y ser invocadas posteriormente.
var anonima = func() {
	fmt.Println("Soy una función anonima")
}

// * Funciones diferidas: defer().
// Las funciones defer se ejecutan al final de la función que las contiene.
func diferida() {
	fmt.Println("Esta función se ejecutará al final de main")
}

// * Función recursiva.
// Es una función que se llama a sí misma.
// Debe tener una condición de salida.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// * Closures.
// Son funciones que capturan variables del entorno.
func createCounter() func() int {
	counter := 0
	// La función anónima captura la referencia a la variable `counter` de su entorno léxico osea createCounter().
	return func() int {
		counter++ // La función anónima captura y modifica la variable counter.
		return counter
	}
}

// Otro ejemplo de closure: Generador de funciones.
func createMultiplier(factor int) func(int) int {
	// Esta función retorna otra función que multiplica por el factor
	return func(x int) int {
		return x * factor
	}
}

// * Definición de un tipo de función.
// Suele usarse en funciones complejas.
type tipoFuncion func(string, string)

// Función que retorna un error.
// En Go, el manejo de errores es explícito. Por convención, las funciones
// que pueden fallar retornan un valor de error como último valor de retorno.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Creamos y retornamos un error cuando la operación no es válida
		return 0, errors.New("no se puede dividir por cero")
	}
	// Si todo está bien, retornamos el resultado y nil como error
	return a / b, nil
}

// Funciones de orden superior.
// Función que aplica una transformación a cada elemento de un slice.
func aplicarATodos(nums []int, f func(int) int) []int {
	resultado := make([]int, len(nums))
	for i, v := range nums {
		resultado[i] = f(v)
	}
	return resultado
}

// Función que filtra elementos de un slice según un predicado.
func filtrar(nums []int, predicado func(int) bool) []int {
	var resultado []int
	for _, v := range nums {
		if predicado(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

// Panic y Recovery.
// Demostración de panic (situación excepcional) y recovery (manejo de dicha situación).
func demostracionPanicRecovery() {
	// defer se ejecutará incluso después de un panic
	defer func() {
		// recover() captura el valor pasado a panic()
		if r := recover(); r != nil {
			fmt.Println("Recuperado del panic:", r)
		}
	}()

	fmt.Println("A punto de entrar en panic...")
	panic("¡Esto es un panic controlado!")

	// Esta línea nunca se ejecutará
	fmt.Println("Esta línea nunca se ejecuta")
}

// Funciones con timeout.
func operacionConTimeout(segundos int) (string, bool) {
	// Canal para comunicar el resultado
	c := make(chan string)

	// Goroutine para realizar la operación
	go func() {
		// Simulamos una operación que toma tiempo
		time.Sleep(time.Duration(segundos) * time.Second)
		c <- "Operación completada"
	}()

	// Esperamos el resultado o un timeout
	select {
	case resultado := <-c:
		return resultado, false
	case <-time.After(1 * time.Second): // Timeout después de 1 segundo
		return "", true
	}
}
