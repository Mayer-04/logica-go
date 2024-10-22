package main

import "fmt"

/*
* Métodos en Structs
- Se recomienda siempre trabajar o bien con el receptor de tipo valor o puntero.
- No combinar estos dos tipos de receptor en una misma estructura. Esto evita confusión y posibles errores.
*/

// Product representa un producto con un nombre y un precio.
type Product struct {
	Name  string
	Price float64
}

// Person representa a una persona con nombre, apellido y edad.
type Person struct {
	Name     string
	Lastname string
	Age      int
}

//* `saludar` es un método asociado a la estructura Person.
// El receptor (p Person) es de tipo valor, lo que significa que trabaja con una copia del valor del receptor.
// Esto es útil cuando no se necesita modificar el receptor original y se busca evitar efectos secundarios.
func (p Person) saludar() string {
	return "Hola" + " " + p.Name
}

// * Punteros en métodos
// `modificar` es un método asociado a la estructura Product.
// El receptor (p *Product) es de tipo puntero, lo que permite modificar el valor del receptor original.
// Usar un receptor de tipo puntero es eficiente, especialmente para structs grandes,
// ya que evita copiar toda la estructura.
func (p *Product) modificar() {
	p.Name = "iPhone"
}

//* Métodos con funciones adicionales.
// Es posible definir múltiples métodos para una misma estructura.
// Método para aplicar un descuento al precio del producto.
// !NOTE: No se recomienda porque la estructura Product ya tiene un método de tipo valor.
func (p *Product) aplicarDescuento(descuento float64) {
	p.Price -= descuento
}

func main() {
	// Antes de invocar métodos en una estructura, debes crear una instancia de esa estructura
	//* Crear una instancia de la estructura Person e invocar su método saludar.
	mayer := Person{Name: "Mayer", Lastname: "Prada", Age: 24}
	fmt.Println(mayer.saludar()) // Output: Hola Mayer

	// Crear una instancia del struct Product e invocar su método modificar.
	// La función modificar cambiará el nombre del producto a "iPhone".
	celular := Product{Name: "Xiaomi 12", Price: 2.340}
	celular.modificar()
	fmt.Println(celular) // Output: {iPhone 2340}

	//* Nota: Cuando se llama a un método con receptor puntero, Go automáticamente toma la dirección de la variable.
	//* No es necesario hacer una llamada explícita con la dirección (&celular).

	// Crear una instancia del struct Product e invocar su método aplicarDescuento.
	celular2 := Product{Name: "Samsung Galaxy", Price: 799.99}
	celular2.aplicarDescuento(10.0)
	fmt.Println("Después de aplicar descuento:", celular2) // Output: {Samsung Galaxy 789.99}
}
