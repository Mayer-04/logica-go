# Pilares de la Programación Orientada a Objetos en Go

El lenguaje Go, a diferencia de otros lenguajes orientados a objetos como Java o C++, no tiene una implementación directa de la POO **(como clases o herencia clásica)**. Sin embargo, sigue ofreciendo los pilares fundamentales de la POO mediante diferentes mecanismos. A continuación se describen estos pilares y cómo se implementan en Go.

> [!IMPORTANT]
> POO: Programación Orientada a Objetos en español.

## Abstracción

En Go, la abstracción se logra mediante la creación de **tipos de datos personalizados** usando `structs` y el uso de `interfaces`. La abstracción permite modelar objetos del mundo real mediante la definición de estructuras que contienen propiedades que representan los datos esenciales de un concepto, ocultando detalles innecesarios.

**Tipos personalizados:** Son aquellos que defines tú mismo, más allá de los tipos primitivos como `int`, `string` o `bool`. Estos tipos personalizados se crean utilizando `structs` o mediante `alias de tipos` existentes.

**Ejemplo:**

Si deseas representar un **coche**, podrías crear una estructura `Car` con atributos como `Brand`, `Model` y `Year`, que definen las características clave de un **coche**:

```go
type Car struct {
    Brand string
    Model string
    Year  int
}
```

## Encapsulamiento

El encapsulamiento en Go se consigue controlando la `visibilidad` de los identificadores **(variables, funciones, métodos)** mediante convenciones de **mayúsculas** y **minúsculas**:

- Los nombres que empiezan con mayúscula son públicos (exportados fuera del paquete).
- Los nombres que empiezan con minúscula son privados (solo accesibles dentro del mismo paquete).

```go
type Car struct {
	Brand string // Público (exportado)
	model string // Privado (no exportado)
}
```

```go
func (c *Car) Model() string {
    return c.model
}

func (c *Car) SetModel(model string) {
    c.model = model
}
```

En este ejemplo, el campo `model` es privado, pero podemos acceder a él a través de los métodos públicos `Model()` y `SetModel()`.

## Herencia

Go no tiene herencia como en otros lenguajes orientados a objetos, pero puedes lograr algo similar mediante la `composición`.

## Composición

La composición es la forma principal en que Go maneja la reutilización de código y la construcción de objetos complejos. En lugar de que una clase herede de otra, una estructura en Go puede incluir otra estructura como un campo, conocido como `incrustación de estructuras`.

```go
type Engine struct {
    HorsePower int
}

type Car struct {
    Brand  string
    Model  string
    Year   int
    Engine // Composición: incrustación de la estructura Engine
}
```

La estructura `Car` ahora incluye todos los **campos** y **métodos** de `Engine`, logrando una forma de reutilización de código similar a la herencia.

## Polimorfismo

En Go, el polimorfismo se logra utilizando `interfaces`. Una interfaz define un conjunto de métodos, pero no proporciona la implementación de esos métodos.

- Cualquier tipo que implemente esos métodos se dice que **"satisface"** la interfaz.
- Una interfaz es como un contrato que dice qué métodos deben existir, pero no dice cómo implementarlos.
- A diferencia de otros lenguajes de programación, no es necesario declarar _explícitamente_ que un tipo implementa una interfaz; Go lo determina automáticamente.

```go
type Vehicle interface {
    Start() string
}

// NOTA: La estructura "Car" como la estructura "Bike" implementan el método "Start" de la interfaz "Vehicle".
type Car struct {
    Brand string
}

// Implementa de manera implícita la interfaz Vehicle.
func (c Car) Start() string {
    return "Car is starting"
}

type Bike struct {
    Brand string
}

// Implementa de manera implícita la interfaz Vehicle.
func (b Bike) Start() string {
    return "Bike is starting"
}

// StartVehicle puede trabajar con cualquier tipo que cumpla la interfaz "Vehicle".
func StartVehicle(v Vehicle) {
    fmt.Println(v.Start())
}

// Llamada al polimorfismo.
car := Car{Brand: "Toyota"}
bike := Bike{Brand: "Yamaha"}
StartVehicle(car)
StartVehicle(bike)
```

## Características adicionales de Go

### Métodos en tipos personalizados

Go permite añadir métodos a cualquier tipo personalizado, no solo a estructuras. Esto extiende la capacidad de modelar comportamientos:

```go
type MyInt int

func (mi MyInt) Double() MyInt {
    return mi * 2
}
```

### Interfaces implícitas

En Go, un tipo no necesita declarar explícitamente que implementa una interfaz. Si el tipo tiene los métodos requeridos por la interfaz, se considera que la implementa. Esto proporciona una gran flexibilidad y desacoplamiento en el diseño de software.

### Composición de interfaces

Go permite componer interfaces más grandes a partir de interfaces más pequeñas, lo que fomenta la creación de interfaces pequeñas y específicas:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

Esta capacidad de composición de interfaces refuerza el principio de diseño de "composición sobre herencia" en Go.
