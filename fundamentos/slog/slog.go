package main

import (
	"fmt"
	"log/slog"
	"os"
)

/*
* Paquete slog
- Introducido en la versión 1.21 de Go, `slog` es un paquete nativo diseñado para proporcionar una manera
eficiente y flexible de registrar eventos (logs) en tus aplicaciones.
- A diferencia de otros métodos de logging, `slog` permite una personalización avanzada y mejor
manejo del formato de salida, incluyendo soporte para JSON y otros formatos estructurados.
*/

func main() {
	// Utilizando el logger predeterminado de `slog` que imprime en formato de texto simple.

	// `Info()` registra un mensaje informativo. Se utiliza para indicar eventos generales en la ejecución del programa.
	slog.Info("El proceso se inició correctamente.")

	// `Warn()` registra un mensaje de advertencia. Útil para indicar posibles problemas o condiciones inusuales.
	slog.Warn("Se detectó un comportamiento inusual, pero el proceso continuará.")

	// `Error()` registra un mensaje de error. Indica que algo salió mal en el programa.
	slog.Error("Error al conectar con la base de datos.")

	// Configuración de un logger personalizado con nivel de log `Debug`.
	// Esto asegurará que todos los mensajes, incluidos los de nivel `Debug`, sean registrados.
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)

	// Ahora, usando el logger personalizado, podemos registrar los mismos mensajes pero en formato JSON.
	logger.Debug("Mensaje de depuración en formato JSON.")
	logger.Info("Mensaje de información en formato JSON.")
	logger.Warn("Se detectó un comportamiento inusual en formato JSON.")
	logger.Error("Error al conectar con la base de datos en formato JSON.")

	// Ejemplo de cómo agregar datos adicionales (campos) a los logs.
	logger.Info("Usuario autenticado", "userID", 1234, "username", "mayer.andres")

	// Ejemplo de cómo registrar un error con contexto adicional.
	err := os.Remove("/ejemplo.txt")
	if err != nil {
		logger.Error("Error al eliminar el archivo", "error", err)
	}

	logFatalError(logger, "Error fatal")
	slogFatalf(logger, "Error fatal: %s", "algo")
}

func logFatalError(logger *slog.Logger, message string, args ...any) {
	logger.Error(message, args...)
	os.Exit(1)
}

func slogFatalf(logger *slog.Logger, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logger.Error(msg, "fatal", true)
	os.Exit(1)
}
