# Paquetes y módulos en Go

## Paquetes en Go

Los paquetes son una forma de organizar el código en Go. Cada `archivo` de Go pertenece a un paquete, y los paquetes permiten dividir el código en **módulos** reutilizables. El paquete principal en Go es `main`, que es donde comienza la ejecución de un programa. Puedes importar paquetes dentro de otros para usar su funcionalidad.

Cuando se usa el paquete `main`, el programa genera un archivo ejecutable independiente. Pero cuando un programa forma parte de un paquete que no es `main`, _Go no genera un archivo binario_.

Cuando utilizas funciones de otros paquetes, necesitas importarlos. Por ejemplo, el `paquete "fmt"` proporciona funciones para imprimir en la consola, y el `paquete "strings"` tiene funciones para trabajar con cadenas de texto.

```go
package main

// Importamos el paquete "fmt" para usar la función "Println".
import "fmt"

func main() {
    fmt.Println("¡Hola, Go!")
}
```

## Módulos en Go

Un módulo es una **colección** de uno o más `paquetes`. Los módulos están definidos por un archivo `go.mod`, que especifica las dependencias del proyecto y la versión de Go utilizada. Los módulos facilitan la administración de dependencias externas **(librerías de terceros)** y su actualización.

El archivo `go.mod` se crea en proyectos que requieran dependencias externas.

**Ejemplo de archivo `go.mod`:**

```go
// Nombre del módulo, generalmente la ruta del repositorio.
module github.com/Mayer-04/Logica-go

// Versión de Go que se está utilizando en el proyecto.
go 1.25.0

// Dependencias requeridas.
require (
    // El paquete jwt de golang-jwt es una dependencia externa.
    // v5.2.1 es la versión específica del paquete que se está utilizando.
    github.com/golang-jwt/jwt/v5 v5.2.1
)
```

## Comandos importantes de Go

### go mod init

Este comando inicializa un nuevo módulo en el directorio actual y crea el archivo `go.mod`. Este archivo es crucial para gestionar las dependencias del proyecto y especifica el nombre del módulo y la versión de Go.

```bash
go mod init <nombre-del-módulo>
```

### go build

El comando `go build` compila el código Go y genera un **archivo ejecutable**. No realiza ninguna instalación en el sistema, solo verifica que el código sea correcto y genera un binario listo para ejecutar.

Si por ejemplo, tienes un archivo `main.go`, este comando generará un ejecutable (llamado `main` o `main.exe` según el sistema operativo).

```bash
go build
```

### go get

Este comando se utiliza para descargar e instalar dependencias externas. También actualiza el archivo `go.mod` con las versiones necesarias de las dependencias. A partir de **Go 1.16**, se recomienda usar `go install` en lugar de `go get` para instalar paquetes ejecutables.

> [!NOTE]
> El flag `-u` actualiza las dependencias a sus últimas versiones disponibles.

```bash
go get -u github.com/gofiber/fiber/v3
```

Este comando descarga el paquete `fiber` en su última versión, una biblioteca para crear servidores web en Go, y lo añade a `go.mod`.

### go install

Compila e instala un paquete Go en el sistema, colocando el ejecutable en el directorio `$GOPATH/bin o en GOBIN`. Es útil para instalar herramientas o crear ejecutables de tus propios proyectos que puedes ejecutar desde cualquier lugar.

```bash
go install <nombre-del-paquete>
```

### go mod tidy

Este comando limpia y optimiza el archivo `go.mod`. Agrega dependencias que faltan y elimina las que no están siendo usadas, asegurando que el archivo `go.mod` esté en orden. No actualiza las versiones de las dependencias, pero garantiza que solo las necesarias estén presentes.

```bash
go mod tidy
```

> [!IMPORTANT]
> Se recomienda usar este comando después de agregar o eliminar dependencias para mantener el proyecto limpio.

### go mod vendor

Este comando descarga las dependencias externas y las guarda en una carpeta llamada `vendor` dentro del proyecto. Es útil cuando se quiere que todas las dependencias estén incluidas en el proyecto, lo que facilita trabajar sin conexión a internet o empaquetar el proyecto completo.

```bash
go mod vendor
```

El comando creará la carpeta `vendor` con todas las dependencias. Si quieres compilar el proyecto usando las dependencias locales en lugar de descargarlas, puedes usar:

```bash
go build -mod=vendor
```

> [!NOTE]
> Esto le indica a Go que utilice las dependencias en la carpeta `vendor` en lugar de volver a descargarlas.
