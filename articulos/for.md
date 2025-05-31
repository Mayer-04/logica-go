# Bucle `for` en Go

En Go, el bucle `for` es la única estructura de control de flujo utilizada para ejecutar código repetidamente mientras se cumpla una condición. A diferencia de otros lenguajes, Go no utiliza _while_ o _do-while_; sin embargo, el bucle for puede replicar la funcionalidad de estos.

> [!NOTE]
> En Go, no se utilizan paréntesis alrededor de la condición.

## Sintaxis

El bucle for en Go se compone de tres partes opcionales:

1. **Inicialización:** Se ejecuta solo una vez al inicio del bucle.
2. **Condición:** Evalúa si el bucle continúa ejecutándose. Si es falsa, el bucle termina.
3. **Actualización:** Se ejecuta después de cada iteración para actualizar variables de control.

```go
for [inicialización]; [condición]; [actualización] {
    // código a ejecutar
}
```

**Ejemplo:**

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

Este ejemplo imprimirá los números del 1 al 5.

## Variantes del bucle `for`

### Bucle estilo `while`

Para simular un `while`, simplemente omitimos la parte de inicialización y actualización en la estructura for.

```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
```

Aquí, el bucle continuará ejecutándose mientras i sea menor o igual a 5.

### Bucle estilo `do-while`

En Go, podemos simular un `do-while` usando un bucle **for infinito** con una declaración **break** para salir cuando se cumpla una condición.

```go
i := 1
for {
    fmt.Println(i)
    i++
    if i > 5 {
        break
    }
}
```

Este ejemplo imprimirá los números del 1 al 5 y saldrá del bucle cuando i sea mayor que 5.

### for con `range`

El bucle for con `range` permite iterar sobre colecciones de datos como arreglos, slices, strings y mapas. Esta variante proporciona dos variables en cada iteración: el índice y el valor correspondiente.

> [!NOTE]
> Puedes usar el operador de subrayado (blank) `_` para ignorar el índice o el valor si no los necesitas.

**Sintaxis:**

```go
for índice, valor := range [arreglo | slice | string | mapa] {
    // código
}
```

**Ejemplo:**

Iteración sobre un arreglo:

```go
arr := [5]int{1, 2, 3, 4, 5}
for i, value := range arr {
    fmt.Printf("Índice: %d, Valor: %d\n", i, value)
}
```

### Uso de `range` para ignorar valores

Si deseas solo el valor o el índice, puedes usar `_` para omitir la variable que no necesitas:

```go
nombres := [3]string{"Daniela", "Pedro", "Andrés"}
for _, nombre := range nombres {
    fmt.Println(nombre)
}
```

### `range` con un entero (novedad desde Go 1.22)

Desde la versión 1.22 de Go, es posible usar `range` directamente con un entero, lo que permite iterar desde `0` hasta el número entero -1.

```go
for i := range 5 {
    fmt.Println(i)
}

// Salida:
// 0
// 1
// 2
// 3
// 4
```

Esta funcionalidad es útil para simplificar ciertos bucles cuando solo necesitas iterar un número específico de veces sin crear una variable de conteo explícita.
