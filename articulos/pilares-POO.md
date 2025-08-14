# Pilares de la programación orientada a objetos en Go

Go a diferencia de otros lenguajes orientados a objetos como Java o C++, no tiene un enfoque tradicional de **clases o herencia clásica**. Sin embargo, sigue ofreciendo gran parte de los pilares fundamentales de la `POO` mediante diferentes mecanismos. A continuación se describen estos pilares y cómo se implementan en Go.

> [!NOTE]
> POO: Programación Orientada a Objetos en español.

## Abstracción

En Go, la abstracción se logra mediante la creación de **tipos de datos personalizados** usando `structs` y el uso de `interfaces`.

La abstracción permite modelar objetos del mundo real mediante la definición de estructuras que contienen propiedades que representan los datos esenciales de un concepto, ocultando detalles innecesarios.

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

No pones en esta estructura todos las propiedades o caracteristica de un carro, solo lo importante para ti y tu modelado.

## Encapsulamiento

El **encapsulamiento** nos permite controlar qué partes del código se pueden ver y usar desde fuera. Así protegemos los datos internos.

En Go, esto se hace con algo muy simple:

- Si un nombre empieza con **mayúscula**, es **público** (se puede usar desde fuera del paquete).
- Si empieza con **minúscula**, es **privado** (solo se puede usar dentro del mismo paquete).

```go
type Car struct {
    Brand string // Público
    model string // Privado
}
```

Aunque el campo `model` es **privado**, podemos usar métodos públicos para acceder a él de forma controlada:

```go
func (c *Car) Model() string {
    return c.model
}

func (c *Car) SetModel(m string) {
    c.model = m
}
```

Así mantenemos el control de cómo se usa el campo `model`.

## Composición (en lugar de Herencia)

Go no tiene **herencia** como en otros lenguajes orientados a objetos, pero ofrece algo similar llamado `composición`.

En vez de decir que una cosa *"es un tipo de"*, decimos que una cosa *"tiene un"*.

En lugar de que una clase herede de otra, una **estructura** en Go puede incluir otra **estructura** como un **campo**, conocido como `incrustación de estructuras`.

```go
type Engine struct {
    HorsePower int
}

type Car struct {
    Brand  string
    Model  string
    Year   int
    Engine // Composición: Incrustación de la estructura "Engine"
}
```

Gracias a la composición, podemos acceder a los datos del motor directamente desde el coche:

```go
c := Car{
    Brand: "Toyota",
    Model: "Corolla",
    Year:  2024,
    Engine: Engine{
        HorsePower: 150,
    },
}

fmt.Println(c.HorsePower) // 150
```

La estructura Car ahora incorpora todos los **campos** y **métodos** de `Engine`, lo que permite una reutilización de código comparable a la herencia.

## Polimorfismo

El polimorfismo suena complicado, pero es fácil de entender. Significa que diferentes tipos pueden comportarse de forma similar si cumplen con una misma interfaz.

En Go, usamos `interfaces` para esto. Una interfaz define qué métodos debe tener un tipo para poder usarla, pero no cómo están hechos.

- Cualquier tipo que implemente esos métodos se dice que **"satisface"** la interfaz.
- Una interfaz es como un contrato que dice qué métodos deben existir, pero no dice cómo implementarlos.
- A diferencia de otros lenguajes de programación, no es necesario declarar *explícitamente* que un tipo **implementa** una interfaz; Go lo determina automáticamente.

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

Go permite añadir métodos a cualquier **tipo personalizado**, no solo a estructuras. Esto extiende la capacidad de modelar comportamientos:

```go
type MyInt int

func (mi MyInt) Double() MyInt {
    return mi * 2
}
```

### Interfaces implícitas

En muchos lenguajes, tienes que decir explícitamente que un tipo implementa una interfaz. En Go, no. Si un tipo tiene los métodos necesarios, ya la implementa automáticamente.

Eso hace que el código sea más simple y flexible.

### Composición de interfaces

Go permite componer interfaces más grandes a partir de interfaces más pequeñas, lo que fomenta la creación de interfaces pequeñas y específicas.

Go recomienda crear interfaces pequeñas que hagan una sola cosa.

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
