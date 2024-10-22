# Características Clave y Convenciones

## Manejo de Errores

### Prefijo “Err”

Es una convención en Go que las variables de error comiencen con el prefijo `Err`. Por ejemplo, se sugiere usar nombres como `ErrInvalidInput`, lo que facilita la identificación de las variables que representan errores.

## Nombres de Funciones

### Convención CamelCase

Las funciones siguen la convención `camelCase`, como `fetchUserDetails`. Las **funciones exportadas** deben comenzar con mayúscula, por ejemplo, `FetchUserDetails`.

### Uso de Verbos

Se recomienda comenzar los nombres de funciones con verbos para indicar claramente su propósito. Por ejemplo, se prefiere usar `GetUser` en lugar de simplemente `User`.

## Importaciones de Paquetes

### Importación con `_`

Cuando se usa el subrayado `_` al importar un paquete, este se importa únicamente por su código de inicialización. Esto significa que la `función init()` del paquete se ejecutará automáticamente sin necesidad de referenciar el paquete directamente en el código.

### Importación con `.`

El uso del operador `.` permite acceder a las funciones y variables exportadas de un paquete sin anteponer el nombre del paquete. Esto puede hacer el código más corto, pero también puede reducir la claridad. En proyectos grandes, esto puede causar confusión, ya que no es inmediato de qué paquete provienen las funciones o variables.

Por ejemplo, al usar el operador `.` con el paquete `math`, puedes acceder a constantes y funciones directamente, como `Pi` y `Sin()`, sin necesidad de escribir `math.Pi` o `math.Sin()`.

## Variables de Entorno

### Configuración de Compilación

Algunas de las variables de entorno más comunes son:

- GOARCH=amd64
- GOOS=windows

### Compilación Cruzada

Go permite la compilación cruzada para diferentes sistemas operativos y arquitecturas. Puedes listar las distribuciones de sistemas operativos compatibles con el siguiente comando:

```bash
go tool dist list
```

Para compilar un archivo para diferentes plataformas, utiliza los siguientes comandos:

```bash
GOOS=windows && GOARCH=amd64 go build -o hello.exe hello.go
GOOS=android && GOARCH=arm64 go build -o hello_android.exe hello.go
```

Durante la compilación, se crea un `archivo objeto` que contiene código ensamblador. Luego, un enlazador _(linker)_ combina este archivo objeto para crear un binario ejecutable. El binario se crea en un directorio temporal, conocido como `work`. Al ejecutar `go run`, el binario se llama desde este directorio.

Por ejemplo, para ver dónde se crea el binario, puedes usar:

```bash
go run --work hello.go
```

## Tamaño de un Archivo

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
