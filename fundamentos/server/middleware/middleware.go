package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// wrappedWriter es un struct que envuelve http.ResponseWriter y nos permite capturar el código de estado HTTP.
type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader sobrescribe el método WriteHeader para capturar el código de estado HTTP.
func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode) // Llama al método original de WriteHeader
	w.statusCode = statusCode                // Guarda el código de estado para utilizarlo después en los logs
}

// Logging es un middleware que registra información sobre cada solicitud HTTP.
// Toma como parámetro un http.Handler (el siguiente handler en la cadena) y devuelve otro http.Handler.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capturar el tiempo al inicio de la solicitud.
		start := time.Now()

		// Creamos un wrappedWriter para interceptar el código de estado HTTP.
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Por defecto, asumimos que la respuesta es 200 OK
		}

		// Llamamos al siguiente handler en la cadena con el wrappedWriter y el request.
		next.ServeHTTP(wrapped, r)

		// Registrar la información relevante después de que la solicitud haya sido procesada.
		log.Printf("Status: %d, Method: %s, Path: %s, Duration: %d\n",
			wrapped.statusCode, r.Method, r.URL.Path, time.Since(start).Milliseconds())
	})
}

// greet es un handler simple que responde con "Hello World!".
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// main configura el servidor HTTP y agrega el middleware de registro a las rutas.
func main() {
	// Crear un nuevo enrutador (multiplexor de rutas).
	router := http.NewServeMux()

	// Registrar la ruta "/example" y asignarle el handler greet.
	// Nota: El método y ruta deben ser ajustados en el patrón adecuado, en Go esto no se coloca en el HandleFunc directamente.
	router.HandleFunc("/example", LoggerMiddleware(log.Default(), greet))

	// Configuramos el servidor HTTP.
	server := &http.Server{
		Addr:    ":8080", // El servidor escucha en el puerto 8080.
		Handler: router,  // Aplicamos el middleware Logging al router.
	}

	// Imprimir mensaje en la consola cuando el servidor esté listo.
	fmt.Printf("Server listening on port %s\n", server.Addr)

	// Iniciar el servidor.
	log.Fatal(server.ListenAndServe())
}

// LoggerMiddleware es otro middleware que usa un registrador personalizado para registrar información de la solicitud.
func LoggerMiddleware(logger *log.Logger, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Capturar el tiempo de inicio de la solicitud.
		receivedTime := time.Now()

		// Registrar la solicitud recibida.
		logger.Printf("Request received: Method=%s, Path=%s\n", r.Method, r.URL.Path)

		// Llamar al siguiente handler en la cadena.
		handler(w, r)

		// Registrar la finalización de la solicitud.
		logger.Printf("Request complete: Method=%s, Path=%s, Duration=%dms\n",
			r.Method, r.URL.Path, time.Since(receivedTime).Milliseconds())
	}
}
