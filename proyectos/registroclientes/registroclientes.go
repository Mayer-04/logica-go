package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Cliente struct {
	ID        int
	Nombre    string
	Direccion string
	Telefono  int
}

// Función constructora para el cliente
func New(id int, nombre, direccion string, telefono int) *Cliente {
	return &Cliente{
		ID:        id,
		Nombre:    nombre,
		Direccion: direccion,
		Telefono:  telefono,
	}
}

// Tipo personalizado que contendra un slice de clientes
type Clientes []Cliente

const maximoDeClientes = 10

var (
	titulo = "REGISTRO DE CLIENTES"
	menu   = `
	1) Agregar cliente
	2) Mostrar clientes
	3) Eliminar cliente
	4) Salir
	`
)

var listaClientes = Clientes{
	{ID: 1, Nombre: "Mayer", Direccion: "Calle 1", Telefono: 1234},
	{ID: 2, Nombre: "Andres", Direccion: "Calle 2", Telefono: 56789},
}

func input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func main() {
	// Utilizando una etiqueta para salir del bucle
loop:
	for {
		fmt.Println(titulo)
		fmt.Println(menu)
		opcion := input("Ingresa una opción: ")
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
	numeroClientes := len(listaClientes) + 1

	nombre := input("Ingrese el nombre del cliente: ")
	if !validarNombre(nombre) {
		color.Red("Error: el nombre solo puede contener letras y espacios.")
		return
	}

	for _, name := range listaClientes {
		if nombre == name.Nombre {
			color.Red("El nombre ya existe")
			return
		}
	}

	direccion := input("Ingrese la dirección: ")
	if err := validarDireccion(direccion); err != nil {
		fmt.Println(err)
		return
	}

	telefono := input("Ingrese el número de teléfono: ")
	numeroTelefono, err := validarTelefono(telefono)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(listaClientes) >= maximoDeClientes {
		color.Red("Error: Número máximo de clientes alcanzado.")
		return
	}

	nuevoCliente := New(numeroClientes, nombre, direccion, numeroTelefono)
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

	color.Cyan("Número de clientes: %d\n", len(listaClientes))

	// Recorrer todos los clientes de la lista y mostrarlos
	for _, cliente := range listaClientes {
		// fmt.Printf("%+v\n", cliente)
		color.Blue("ID: %d, Nombre: %s, Dirección: %s, Teléfono: %d\n",
			cliente.ID, cliente.Nombre, cliente.Direccion, cliente.Telefono,
		)
	}
}

func eliminarCliente() {
	nombre := input("Ingresa el nombre del cliente a eliminar: ")

	if nombre == "" {
		color.Red("Error: el nombre no puede estar vacío.")
		return
	}

	// Encontrar y eliminar el cliente por nombre
	for i, cliente := range listaClientes {
		if nombre == cliente.Nombre {
			// Eliminar al cliente de la lista
			// listaClientes = append(listaClientes[:i], listaClientes[i+1:]...)
			listaClientes = slices.Delete(listaClientes, i, i+1)
			color.Green("Cliente eliminado: %s", cliente.Nombre)
			return
		}
	}
	color.Red("Cliente '%s' no encontrado en la lista.", nombre)
}

// * VALIDACIONES
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
	numero, err := strconv.Atoi(opcion)
	if err != nil {
		return 0, fmt.Errorf("error: La opción debe ser un número")
	}
	return numero, nil
}
