# Concatenación de cadenas en Go: Cómo unir múltiples cadenas de manera eficiente

En Go, concatenar cadenas, es decir, unir varias cadenas en una sola, se puede realizar de varias formas.
Una de las más simples es usar el operador `+=`:

```go
package main

import "fmt"

func main() {
 sli := []string{"Hello", "World"}

    // result := ""
 var result string
 for _, val := range sli {
  result += val
 }

 fmt.Println(result) // Output: HelloWorld
}
```

## ¿Por qué usar += es ineficiente?

Usar `+=` para concatenar cadenas es simple, pero no es la forma más eficiente. Esto se debe a cómo Go maneja las cadenas internamente:

- **Inmutabilidad de las cadenas:** En Go, las cadenas son `inmutables`, lo que significa que no se pueden modificar, agregar o eliminar sus caracteres después de ser creadas. Cada vez que usas `+=`, se crea una nueva cadena en memoria que combina la cadena existente con la nueva parte añadida.
- **Asignación de memoria:** Cada concatenación con `+=` requiere _asignar un bloque de memoria más grande_ para la nueva cadena y copiar el contenido de las cadenas existentes y la nueva parte en este nuevo espacio.
- **Recolector de basura:** El recolector de basura de Go automáticamente libera el espacio de memoria ocupado por las cadenas anteriores. Sin embargo, en operaciones de concatenación extremadamente grandes, esto puede llegar a afectar el rendimiento.

**Desglosando el código anterior:**

**Primera iteración:**

- `result` es una cadena vacía.
- `val` es "Hello".
- `result += val` da como resultado `"Hello"`.
- Se asigna memoria para "Hello".

**Segunda iteración:**

- `result` ahora es `"Hello"`.
- `val` es "World".
- `result += val` da como resultado `"HelloWorld"`.
- Se asigna nueva memoria para "HelloWorld", copiando "Hello" y "World" en el nuevo espacio de memoria.
- La memoria usada para "Hello" se liberará automáticamente con el `recolector de basura (GC)`.

## Formas más eficientes de concatenar cadenas

Para solucionar esto, podemos usar el `paquete strings` de la biblioteca estándar de Go, que proporciona formas más eficientes de concatenar cadenas.

### Uso de strings.Builder

`strings.Builder` es una estructura en Go que facilita la construcción de cadenas de manera eficiente.

1. `strings.Builder` es una estructura vacía que internamente contiene un `búfer` donde se almacenan las partes de las cadenas que se van añadiendo.
2. El método `sb.WriteString(value)` se utiliza para agregar el contenido de la cadena value al búfer interno de strings.Builder. Retorna dos valores:
   - El número de bytes escritos (longitud de la cadena añadida).
   - Un error, que en este caso siempre es nil (nulo), porque WriteString nunca falla cuando se usa con strings.Builder.
3. El método `sb.String()` devuelve la cadena completa acumulada en el búfer de strings.Builder como un string.

- **Búfer:** Es un área de memoria utilizada para almacenar **temporalmente** datos mientras se están moviendo de un lugar a otro.

**Ejemplo:**

```go
package main

import (
 "fmt"
 "strings"
)

func main() {
 stringsToJoin := []string{"Hello", "World"}

 var builder strings.Builder
 for _, s := range stringsToJoin {
  builder.WriteString(s)
 }

 joinedString := builder.String()
 fmt.Println(joinedString) // Output: "HelloWorld"
}
```

**En este ejemplo:**

1. Creamos una instancia de `strings.Builder`.
2. Usamos `sb.WriteString(s)` para agregar cada cadena del slice al strings.Builder.
3. Finalmente, obtenemos la cadena completa con `sb.String()` y la imprimimos.

**Segundo ejemplo con strings.Builder usando un separador:**

```go
package main

import (
 "fmt"
 "strings"
)

func main() {
 stringsToJoin := []string{"Hello", "World", "!"}
 separator := " " // El separador que quieres usar

 var builder strings.Builder
 for i, s := range stringsToJoin {
  if i > 0 {
   builder.WriteString(separator)
  }
  builder.WriteString(s)
 }

 joinedString := builder.String()
 fmt.Println(joinedString) // Output: "Hello World !"
}
```

### Uso de strings.Join()

`strings.Join()` es una función que toma un slice de cadenas y las une en una sola cadena, usando el segundo argumento como **separador**.

Al usar `strings.Join()`, Go optimiza la concatenación internamente, evitando las múltiples asignaciones de memoria y copias que ocurren con el operador `+=`.

**IMPORTANTE:** Internamente, el método `Join` del paquete strings implementa `strings.Builder` y sus funciones, lo que le permite realizar la concatenación de manera eficiente.

```go
package main

import (
 "fmt"
 "strings"
)

func main() {
 stringsToJoin := []string{"Hello", "World"}
 joinedString := strings.Join(stringsToJoin, "")

 fmt.Println(joinedString) // Output: HelloWorld
}
```

## Uso de fmt.Sprintf() para concatenación

Otra forma común de concatenar cadenas en Go es utilizando `fmt.Sprintf()`, que es útil cuando se necesita formatear las cadenas al mismo tiempo.

`fmt.Sprintf()` es ideal cuando necesitas incluir variables de distintos tipos en una cadena o cuando el formato de la cadena es importante.

```go
package main

import "fmt"

func main() {
    name := "Mayer"
    greeting := fmt.Sprintf("Hello, %s!", name)
    fmt.Println(greeting) // Output: Hello, Mayer!
}
```

Aunque `fmt.Sprintf()` es más versátil y puede ser más legible en casos simples o cuando se requiere formateo, no es muy eficiente.
