# Convenciones esenciales y características clave del lenguaje Go

## Manejo de errores en Go

### Variables de error con prefijo `Err`

En Go, es una convención que las variables de error comiencen con el prefijo `Err`. Esto facilita la identificación de las variables que representan errores y hace que el código sea más legible.

```go
var ErrInvalidInput = errors.New("entrada inválida")
```

## Nombres de funciones en Go

### Convención camelCase

Las funciones en Go siguen la convención `camelCase`. Las funciones exportadas deben comenzar con mayúscula **(UpperCamelCase)**, mientras que las funciones internas del paquete deben iniciar con minúscula.

```go
// Esta función es exportada porque empieza con mayúscula
func FetchUserDetails() {}

// Esta función es interna del paquete porque comienza con minúscula
func fetchUserDetails() {}
```

### Uso de verbos

Se recomienda que los nombres de las funciones comiencen con verbos para indicar claramente su propósito. Por ejemplo, es preferible GetUser en lugar de User.

```go
// ✅
func GetUser() {}

// ❎
func User() {}
```

## Importaciones de paquetes

### Importación con `_`

Cuando se usa el subrayado `_` al importar un paquete, este se importa únicamente por su código de inicialización. Esto significa que la `función init()` del paquete se ejecutará automáticamente sin necesidad de referenciar el paquete directamente en el código.

```go
import _ "database/sql"
```

### Importación con `.`

El uso del operador `.` permite acceder a las funciones y variables exportadas de un paquete sin anteponer el nombre del paquete. Esto puede hacer el código más corto, pero también puede reducir la claridad. En proyectos grandes, esto puede causar confusión, ya que no es inmediato de qué paquete provienen las funciones o variables.

Por ejemplo, al usar el operador `.` con el paquete `math`, puedes acceder a constantes y funciones directamente, como `Pi` y `Sin()`, sin necesidad de escribir `math.Pi` o `math.Sin()`.

```go
import . "math"

func main() {
    println(Sin(0.5)) // En lugar de math.Sin(0.5)
}
```

## Variables de entorno

### Configuración de compilación

Algunas de las variables de entorno más comunes son:

```bash
GOOS=windows # Sistema operativo.
GOARCH=amd64 # Arquitectura del procesador.
```

### Compilación cruzada en Go (Cross-Compilation)

Go permite la compilación cruzada, lo que significa que puedes compilar tu código en una máquina y generar un ejecutable para otro sistema operativo o arquitectura.

Esto es útil para distribuir tu software en múltiples plataformas sin requerir que los usuarios tengan Go instalado.

Puedes listar sistemas operativos y arquitecturas compatibles con el siguiente comando:

```bash
go tool dist list
```

Este comando te mostrará pares de `GOOS/GOARCH` compatibles, como **linux/amd64**, **windows/arm64**, **darwin/arm64**, **solaris/amd64**, etc.

### Ejemplo de compilación cruzada

Para compilar un programa en Go para diferentes plataformas, debes definir las variables de entorno `GOOS` (Sistema Operativo) y `GOARCH` (Arquitectura) antes de ejecutar `go build`.

```bash
GOOS=windows && GOARCH=amd64 go build -o hello.exe hello.go
GOOS=android && GOARCH=arm64 go build -o hello_android.exe hello.go
GOOS=linux GOARCH=amd64 go build -o hello_linux hello.go
```

### ¿Cómo funciona internamente la compilación cruzada en Go?

Cuando compilas un programa en Go:

1. El compilador genera un *archivo objeto*, que contiene código en ensamblador específico para la plataforma de destino.

2. El *enlazador (linker)* toma este archivo objeto y lo combina con las bibliotecas estándar de Go para producir un binario ejecutable.

3. El ejecutable generado es completamente independiente y no requiere dependencias externas, lo que lo hace fácil de distribuir.

Si ejecutas `go run`, Go crea un binario temporal en un directorio de trabajo *(work)* y lo ejecuta automáticamente. Puedes ver este directorio usando:

```bash
go run --work hello.go
```

Esto te mostrará la ruta del binario temporal antes de ejecutarlo.

## Tamaño de un archivo en linux

Para conocer el `tamaño de un archivo` en Linux, puedes usar los siguientes comandos:

**Muestra el tamaño total del archivo:**

```bash
du -sh prueba.go
```

**Muestra el tamaño de todos los archivos y directorios:**

```bash
du -h prueba.exe
```

> [!NOTE]
> Esta pequeña sección no esta relacionada con Go, es más personal 😁
