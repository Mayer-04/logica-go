package main

import "fmt"

/*
* Métodos en Structs (Estructuras)
En Go, un método es una función que se asocia a una estructura mediante un `receptor`.

Esto permite que las estructuras tengan comportamientos, similar a los métodos en la programación
orientada a objetos.

* Receptor:
- Es un parámetro especial entre la palabra clave `func` y el nombre del método.
- Permite acceder a los campos y otros métodos de la estructura.
- Por convención, se usa como nombre del receptor la primera letra del nombre de la estructura.
  Ejemplo: `p` para `Person`, `c` para `Car`.

* Tipos de receptor:

1. Receptor por valor (T):
   - Trabaja con una `copia` de la estructura original.
   - Útil cuando el método `no modifica` el estado del receptor.

2. Receptor por puntero (*T):
   - Trabaja con una `referencia` a la estructura original.
   - Permite modificar campos o evitar copiar estructuras grandes.

📌 Buena práctica:
Usa un único tipo de receptor (valor o puntero) de forma consistente en todos los métodos de una estructura.
Si varios métodos modifican la estructura, es mejor usar receptores por puntero para todos.
*/

// Product representa un producto con nombre y precio.
type Product struct {
	Name  string
	Price float64
}

// Person representa una persona con nombre, apellido y edad.
type Person struct {
	Name     string
	Lastname string
	Age      int
}

// saludar es un método asociado a Person con receptor por valor.
// Como no modifica el estado de la persona, se usa una copia.
func (p Person) saludar() string {
	return "Hola " + p.Name
}

// * Punteros en métodos
// `modificar` es un método asociado a la estructura Product.
// El receptor (p *Product) es de tipo puntero, lo que permite modificar el valor del receptor original.
func (p *Product) modificar() {
	p.Name = "iPhone"
}

// * Métodos con funciones adicionales.
// Es posible definir múltiples métodos para una misma estructura.
// Por ejemplo: `aplicarDescuento` es un método para aplicar un descuento al precio del producto.
// NOTE: No se recomienda porque la estructura Product ya tiene un método de tipo valor.
func (p *Product) aplicarDescuento(descuento float64) {
	p.Price -= descuento
}

func main() {
	// Crear una instancia de Person e invocar su método.
	mayer := Person{Name: "Mayer", Lastname: "Prada", Age: 24}
	fmt.Println(mayer.saludar()) // Output: Hola Mayer

	// Crear una instancia de Product e invocar su método modificar.
	celular := Product{Name: "Xiaomi 12", Price: 2340.00}
	celular.modificar()
	fmt.Println("Producto modificado:", celular) // Output: {iPhone 2340}

	// Go toma automáticamente la dirección (&celular) cuando se necesita un puntero.
	// No es necesario hacerlo explícitamente al invocar el método.

	// Crear un producto nuevo y aplicar un descuento.
	celular2 := Product{Name: "Samsung Galaxy", Price: 799.99}
	celular2.aplicarDescuento(10.0)
	fmt.Println("Después de aplicar descuento:", celular2) // Output: {Samsung Galaxy 789.99}
}
