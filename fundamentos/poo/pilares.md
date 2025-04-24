# Pilares de la Programación Orientada a Objetos

## Conceptos claves

### Objeto

Un objeto es una instancia de una clase, es decir, algo concreto que se ha creado a partir de esa plantilla. Si la clase es el plano, el objeto es la construcción final.

**Ejemplos:**

- Un objeto de la **clase Animal** podría ser un perro llamado "Zeus".
- Un objeto de la **clase Car** podría ser un coche específico, como un Nissan Tiida del año 2014.

### Clase

Una clase es como un plano o molde que define cómo va a ser un objeto. Es una plantilla que describe qué características (propiedades) y qué acciones (métodos) puede tener algo en tu programa.

Los métodos pueden interactuar con las propiedades de la clase.

**Ejemplos:**

- La **clase Animal** define qué características tendrán todos los animales, como especie, género, nombre y peso.
- La **clase Car** especifica qué características tendrán todos los coches, como marca, modelo, color y año.

### Atributos (o Propiedades)

Son las características o datos que describen a un objeto.

**Ejemplos:**

- En la **clase Animal**, los atributos pueden ser especie, género, nombre y peso.
- En la **clase Car**, los atributos pueden ser marca, modelo, color y año.

### Métodos

Un método es una función que pertenece a una clase. Define una acción que los objetos de esa clase pueden realizar. Suelen ser verbos.

**Ejemplos:**

- Algunas acciones o métodos que podria tener la **clase Animal** podrian ser `Dormir` o `Comer`.
- Algunas acciones o métodos que podria tener la **clase Car** podrian ser `Manejar` o `Acelerar`.

### Instanciación

Es el proceso de crear un objeto a partir de una clase. Al `instanciar` una clase, estás creando un nuevo objeto basado en esa plantilla.
Cada vez que se instancia una clase, se crea un nuevo objeto con su propio conjunto de atributos, pero todos basados en la misma definición de clase.

**Ejemplos:**

- Podemos crear un objeto llamado _perro_, que es una instancia de la **clase Animal**. Este objeto tendrá características y acciones definidas en la clase Animal, como especie, género, nombre, Dormir o Comer.
- Podemos crear un objeto llamado _toyota_, que es una instancia de la **clase Car**. Este objeto tendrá características y acciones definidas en la clase Car, como marca, modelo, año, Manejar o Acelerar.

### Constructor

Un constructor es un `método especial` que se llama automáticamente cuando se crea un objeto. Su propósito es inicializar el objeto, configurando sus atributos con valores iniciales.

El constructor es invocado automáticamente al crear un objeto.

**Ejemplos:**

- En la **clase Animal**, el constructor puede inicializar atributos como especie y nombre.
- En la **clase Car**, el constructor puede establecer atributos como marca y año.

## Abstracción

La abstracción consiste en identificar y seleccionar las características más importantes de un objeto del mundo real, y representarlas en una clase. Esto nos permite centrarnos en lo que es relevante y simplificar la complejidad del mundo real.

## Encapsulamiento

El encapsulamiento es como poner un escudo protector alrededor de los datos de un objeto. Esto evita que la información se modifique sin permiso, asegurando que solo las funciones dentro de la clase puedan acceder o cambiar esos datos.

## Herencia

La herencia permite que una clase nueva _(hija)_ herede las propiedades y métodos de otra clase existente _(padre)_. Esto ayuda a reutilizar código y a crear relaciones entre diferentes clases.

## Composición

La composición es una forma de construir un objeto utilizando otros objetos como componentes. En lugar de heredar características, un objeto las tiene porque incluye otros objetos dentro de sí mismo. Es como construir algo grande usando piezas más pequeñas.

## Polimorfismo

El polimorfismo permite que objetos de clases diferentes respondan de manera similar a los mismos métodos, pero cada uno puede hacerlo de manera única. En otras palabras, aunque diferentes objetos puedan realizar la misma acción, lo hacen a su manera.

- **Varias formas de hacer lo mismo:** Diferentes clases pueden tener métodos con el mismo nombre, pero implementados de manera distinta.
- **Uso de interfaces:** Una interfaz es como un contrato que dice qué métodos deben existir, pero no dice cómo implementarlos. Cada clase decide cómo hacer su trabajo mientras cumple con ese contrato.
