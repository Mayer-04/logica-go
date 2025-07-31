package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
)

/*
* Servidor HTTP en Go
- El paquete 'net/http' permite crear y gestionar servidores HTTP de manera nativa en Go.
- El paquete 'encoding/json' facilita el trabajo con datos en formato JSON,
proporcionando funciones para su codificación y decodificación.

* Multiplexor de solicitudes (Mux):
- Se utiliza para organizar y gestionar cómo se manejan diferentes rutas en un servidor HTTP.
- Permite definir rutas específicas que corresponden a diferentes manejadores (handlers).
- Cuando llega una solicitud HTTP, el multiplexor verifica la ruta solicitada y la asigna al manejador
correspondiente para procesar la petición y responder al cliente.

* Manejadores (Handlers):
- Los manejadores son funciones que responden a las solicitudes HTTP,
procesando la información recibida y generando una respuesta para el cliente.
- Cada manejador toma una solicitud HTTP, realiza la lógica necesaria
(como acceder a una base de datos, procesar datos, etc.), y envía una respuesta al cliente.

* Rutas (Routes):
- Las rutas son patrones de URL que dirigen las solicitudes HTTP a los manejadores correspondientes.
- Cada ruta se asocia a un manejador específico, lo que permite que el servidor responda de manera
diferente según la URL solicitada.
- Por ejemplo, la ruta `/usuarios` puede listar todos los usuarios.


* Algunos conceptos a tener en cuenta:
- Escribir datos: Consiste en enviar o transferir información desde un programa (como un servidor)
a un cliente HTTP, una conexión de red, un archivo o un buffer.

- fmt.Fprintf(): Permite escribir datos formateados a un destino, que puede ser un archivo, una conexión de red o un buffer.
El primer argumento es un 'io.Writer', que determina el destino de la escritura.

- http.Error(): Envía un mensaje de error al cliente como parte de la respuesta HTTP,
junto con el código de estado correspondiente.

- w.Write(): Escribe datos en una respuesta HTTP que se enviará al cliente. El argumento que toma es un []byte,
que generalmente representa datos binarios o texto. Estos datos se enviarán en el cuerpo de la respuesta HTTP.

- io.WriteString(): Se utiliza para escribir (s) el contenido de una cadena de texto en un destino (w)
que puede aceptar datos, como un archivo, la salida estándar o una respuesta HTTP.
*/

// Definimos una estructura 'User' para representar un usuario.
// Las etiquetas `json:""` especifican cómo los campos serán representados en JSON.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// * Simulación de una base de datos en memoria.
// La variable 'users' almacena los usuarios en un mapa.
// La clave es un entero que representa el ID del usuario, y el valor es un estructura `User`.
var (
	users = make(map[int]User)
	mu    sync.RWMutex
)

func main() {
	// Crea un servidor HTTP.
	// Nuevo multiplexor de solicitudes HTTP.
	mux := http.NewServeMux()

	// Define la ruta y el manejo de la petición.
	mux.HandleFunc("/", handleRoot)

	// La solicitud debe ser de tipo 'POST' y la ruta debe ser '/users'.
	// Esta sintaxis es solo válida a partir de la versión 1.22 de Go.
	mux.HandleFunc("POST /users", createUser)

	// La solicitud debe ser de tipo 'GET' y la ruta debe ser '/users/{id}'.
	// Recupera el parámetro de la ruta '{id}' de la solicitud que captura un valor dinámico.
	mux.HandleFunc("GET /users/{id}", getUsers)

	// La solicitud debe ser de tipo 'DELETE' y la ruta debe ser '/users/{id}'.
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	// Mostramos un mensaje en la consola cuando el servidor se inicia.
	// `slog.Info()` es una función que registra un mensaje en la consola.
	slog.Info("Iniciando servidor", "puerto", ":8080")

	//* Iniciamos el servidor.
	// El primer argumento es la dirección o puerto donde escuchará el servidor.
	// El segundo argumento es el manejador de solicitudes (en este caso, `mux`).
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Error fatal del servidor", "error", err.Error())
		os.Exit(1)
	}
}

// Controlador para manejar la ruta raíz.
// Escribe un mensaje de "Hola Mundo" como respuesta.
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Nos permite enviar cualquier contenido en la respuesta.
	// fmt.Fprintf(w, "Hola Mundo!")

	// Usamos `io.WriteString` para escribir directamente la respuesta como cadena.
	io.WriteString(w, "Hola Mundo!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Creamos una instancia vacía de `User` para almacenar los datos del cuerpo de la solicitud.
	// Go inicializa los campos de la estructura `User` a sus valores cero correspondientes.
	var user User

	// Crea un decodificador basado en el cuerpo (body) de la petición.
	// Decodifica los datos de la petición y los almacena en la variable 'user'.
	// La variable 'user' contendra todos los datos de que se especifique en el body de la petición.
	// Evidentemente el cuerpo de la solicitud debe coincidir con nuestra estructura 'User'.
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		// Le pasamos la respuesta, el error y el estado de la petición en este caso 400.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//* Validaciones básicas.
	// Verificamos que los campos requeridos no estén vacíos.
	if user.Name == "" {
		http.Error(w, "el nombre es requerido", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "el email es requerido", http.StatusBadRequest)
		return
	}

	// El ID del usuario se crea como el siguiente entero mayor al último ID guardado.
	userID := len(users) + 1

	mu.Lock()
	defer mu.Unlock()

	// Guardamos el usuario en la "base de datos".
	users[userID] = user

	// Indicamos que la solicitud fue exitosa con un estado 201.
	w.WriteHeader(http.StatusCreated)

	// Respondemos con un mensaje de confirmación.
	fmt.Fprintf(w, "el usuario %s fue creado exitosamente", user.Name)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Recuperamos el valor del parámetro '{id}' de la ruta de la solicitud.
	// Este método `PathValue` está disponible desde la versión 1.22 de Go.
	userId := r.PathValue("id")

	// Convertimos el ID de usuario de string a entero.
	id, err := strconv.Atoi(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.RLock()
	defer mu.RUnlock()

	// Buscamos el usuario en el mapa `users`.
	user, ok := users[id]

	if !ok {
		// Si el usuario no existe, devolvemos un error 404 (Not Found).
		http.Error(w, "el usuario no fue encontrado", http.StatusNotFound)
		return
	}

	// Configuramos la cabecera de la respuesta para indicar que es de tipo JSON.
	w.Header().Set("Content-Type", "application/json")

	// Transformamos el usuario a JSON para mandarlo como respuesta al cliente.
	data, err := json.Marshal(user)

	if err != nil {
		// Si ocurre un error al codificar el JSON, devolvemos un error 500 (Internal Server Error).
		// Es un error interno del servidor porque hasta este punto el usuario ingreso las datos correctamente.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devuelve la respuesta al cliente con el estado 200.
	w.WriteHeader(http.StatusOK)
	// Devuelve los datos del usuario en formato JSON.
	w.Write(data)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Recuperamos el valor del parámetro '{id}' de la ruta de la solicitud.
	userId := r.PathValue("id")

	// Convertimos el ID de usuario de string a entero.
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Verificamos si el usuario existe en el mapa.
	if _, ok := users[id]; !ok {
		http.Error(w, "el usuario no fue encontrado", http.StatusNotFound)
		return
	}

	// Eliminamos el par clave-valor del mapa por su id.
	delete(users, id)

	// Indicamos que la petición fue exitosa con un estado 204.
	// El código 204 indica que la acción fue exitosa, pero no es necesario devolver ninguna información.
	w.WriteHeader(http.StatusNoContent)
}

// renderJSON es una función que se encarga de enviar una respuesta HTTP con un cuerpo
// codificado en formato JSON.
func RenderJSON[T any](w http.ResponseWriter, data T) {
	// Configura la cabecera de la respuesta para indicar que el contenido es JSON.
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	// Codifica el valor 'T' en formato JSON.
	if err := encoder.Encode(data); err != nil {
		// Si ocurre un error al codificar el JSON, se devuelve un error 500 (Internal Server Error)
		// con el mensaje de error correspondiente.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
