package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Estructura que contiene los campos del cliente
type Cliente struct {
	Nombre    string
	Direccion string
	Telefono  int
}

// Función constructora para el cliente
func New(nombre, direccion string, telefono int) *Cliente {
	return &Cliente{
		Nombre:    nombre,
		Direccion: direccion,
		Telefono:  telefono,
	}
}

// Tipo personalizado que contendra un slice de clientes
type Clientes []Cliente

const maximoDeClientes = 10

var titulo = "REGISTRO DE CLIENTES"

var menu = `
	1) Agregar cliente
	2) Mostrar clientes
	3) Eliminar cliente
	4) Salir
`

var listaClientes = []Cliente{
	{Nombre: "Mayer", Direccion: "Calle 1", Telefono: 1234},
	{Nombre: "Andres", Direccion: "Calle 2", Telefono: 56789},
}

func main() {
	// Utilizando una etiqueta para salir del bucle
loop:
	for {
		fmt.Println(titulo)
		fmt.Println(menu)
		var opcion string
		fmt.Print("Ingresa una opción: ")
		fmt.Scanln(&opcion)
		validacion, err := validarOpcion(opcion)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch validacion {
		case 1:
			agregarCliente()
		case 2:
			mostrarClientes()
		case 3:
			eliminarCliente()
		case 4:
			break loop
		default:
			color.Red("Opción no válida")
		}
	}
}

func agregarCliente() {
	var nombre, direccion string
	var telefono string

	fmt.Print("Ingrese el nombre del cliente: ")
	fmt.Scanln(&nombre)
	if !validarNombre(nombre) {
		color.Red("Error: el nombre solo puede contener letras y espacios.")
		return
	}

	fmt.Print("Ingrese la dirección: ")
	fmt.Scanln(&direccion)
	if err := validarDireccion(direccion); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Ingrese el número de teléfono: ")
	fmt.Scanln(&telefono)
	numeroTelefono, err := validarTelefono(telefono)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(listaClientes) >= maximoDeClientes {
		color.Red("Error: Número máximo de clientes alcanzado.")
		return
	}

	nuevoCliente := New(nombre, direccion, numeroTelefono)
	listaClientes = append(listaClientes, *nuevoCliente)
	color.Green("Cliente agregado correctamente.")
}

func mostrarClientes() {
	numeroClientes := len(listaClientes)

	// Si en el slice no hay clientes
	if numeroClientes == 0 {
		color.Red("No hay clientes para mostrar.")
		return
	}

	fmt.Printf("Número de clientes: %d\n", len(listaClientes))
	// Recorrer todos los clientes de la lista y mostrarlos
	for _, cliente := range listaClientes {
		// fmt.Printf("%+v\n", cliente)
		color.Blue("Nombre: %s, Dirección: %s, Teléfono: %d\n", cliente.Nombre, cliente.Direccion, cliente.Telefono)
	}
}

func eliminarCliente() {
	var nombre string
	fmt.Print("Ingresa el nombre del cliente a eliminar: ")
	fmt.Scanln(&nombre)

	// Encontrar y eliminar el cliente por nombre
	for i, cliente := range listaClientes {
		if nombre == cliente.Nombre {
			// Eliminar al cliente de la lista
			listaClientes = append(listaClientes[:i], listaClientes[i+1:]...)
			color.Green("Cliente eliminado: ", cliente.Nombre)
			return
		}
	}
	color.Red("Cliente no encontrado.")
}

//* VALIDACIONES

// TODO: Validar que el nombre no este ya en la lista de clientes.
func validarNombre(nombre string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return regex.MatchString(nombre)
}

func validarDireccion(direccion string) error {
	if direccion == "" {
		return errors.New("la dirección no puede estar vacía")
	}
	return nil
}

func validarTelefono(telefono string) (int, error) {
	numero, err := strconv.Atoi(telefono)
	if err != nil {
		return 0, fmt.Errorf("el teléfono debe ser un número")
	}
	if len(telefono) < 3 {
		return 0, fmt.Errorf("el teléfono debe tener al menos 3 dígitos")
	}
	return numero, nil
}

func validarOpcion(opcion string) (int, error) {
	removiendoEspacios := strings.TrimSpace(opcion)
	numero, err := strconv.Atoi(removiendoEspacios)
	if err != nil {
		return 0, fmt.Errorf("error: La opción debe ser un número")
	}
	return numero, nil
}
