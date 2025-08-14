package main

import (
	"fmt"
	"log"
	"net/mail"
)

/*
* Paquete mail
Proporciona herramientas para analizar y formatear mensajes de correo electrónico.
- Leer, analizar y manejar correos electrónicos de manera eficiente.
*/

func main() {
	//* ParseAddress.
	// Analiza una dirección de correo electrónico en formato de texto y devuelve un objeto `mail.Address`.
	// Es útil para validar y extraer la información de una dirección de correo.
	addr, err := mail.ParseAddress("Mayer Andres <mayer@example.com>")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", addr.Name)
	fmt.Printf("addr: %s\n", addr.Address)

	//* ParseAddressList.
	// Permite analizar una lista de direcciones de correo separadas por comas.
	addresses, err := mail.ParseAddressList("mayer@example.com, andres@example.com")
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addresses {
		fmt.Println("address:", addr.Address)
	}
}
