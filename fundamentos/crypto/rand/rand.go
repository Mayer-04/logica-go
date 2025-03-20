package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	// La función Text() devuelve un string aleatorio de 32 bytes.
	// El resultado contiene al menos 128 bits de aleatoriedad.
	text := rand.Text()
	fmt.Println(text) // Output: P3MS3QPOZYJVSCO2MCOY
}
