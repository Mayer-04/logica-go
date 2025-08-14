package main

import (
	"fmt"
	"strconv"
)

// obtenerDia recibe un número que representa el día de la semana
// y devuelve la representación en forma de cadena.
func obtenerDia(dia int) string {
	dias := map[int]string{
		1: "Lunes",
		2: "Martes",
		3: "Miércoles",
		4: "Jueves",
		5: "Viernes",
		6: "Sábado",
		7: "Domingo",
	}

	// Verificar si el día está en el mapa, devolver la representación en caso afirmativo.
	// Si no, devolver "Día inválido".
	if nombreDia, existe := dias[dia]; existe {
		return nombreDia
	}

	return "Día inválido"
}

func main() {
	// Solicitar entrada del usuario.
	var input string
	fmt.Print("Ingrese un dígito entre 1 y 7: ")
	fmt.Scan(&input)

	// Convertir la entrada a entero.
	dia, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Entrada inválida. Debe ser un número entero.")
		return
	}

	// Obtener el día y mostrar el resultado.
	mensaje := obtenerDia(dia)
	fmt.Println("Día de la semana:", mensaje)
}
