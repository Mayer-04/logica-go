# Conceptos importantes de Go

## Concurrencia

La concurrencia es la capacidad de un programa para manejar varias tareas "a la vez", alternando entre ellas, aunque no necesariamente se ejecuten *al mismo tiempo*. En otras palabras, una sola CPU puede cambiar rápidamente entre tareas, dando la sensación de simultaneidad.

**Ejemplo cotidiano:**

Imagina que estás cocinando solo:

- Cortas las verduras 🥕
- Pones agua a hervir 💧
- Mientras hierve, salteas otros ingredientes 🍳

Aunque haces una cosa a la vez, vas cambiando entre tareas sin esperar a que una termine del todo para comenzar la siguiente.

➡️ Concurrencia = Hacer varias cosas alternadamente, no al mismo tiempo.

## Paralelismo

El paralelismo ocurre cuando múltiples tareas se ejecutan exactamente al mismo tiempo, aprovechando varios núcleos o procesadores de nuestra computadora. Esto mejora el rendimiento en tareas intensivas o repetitivas.

**Ejemplo cotidiano:**

Ahora imagina que tienes tres personas cocinando, cada una en su estación:

- Una corta verduras 🥕
- Otra hierve agua 💧
- Y otra saltea ingredientes 🍳

Todas las tareas ocurren **simultáneamente**.

➡️ Paralelismo = Hacer muchas cosas al mismo tiempo.

## Concurrencia vs Paralelismo

```md
| Concepto     | ¿Al mismo tiempo? | ¿Comparte CPU? | ¿Requiere múltiples núcleos? |
| ------------ | ----------------- | -------------- | ---------------------------- |
| Concurrencia | No                | Sí             | No                           |
| Paralelismo  | Sí                | No             | Sí                           |
```

## Runtime de Go

El runtime de Go es el motor interno que se encarga de ejecutar los programas escritos en este lenguaje. Administra cosas como:

- La **memoria** que usa tu programa.
- La ejecución de **goroutines**.
- La **recolección de basura** (elimina datos que ya no se usan).
- La **comunicación entre goroutines**, usando canales *(channels)*.
- El uso de bloqueos *(locks)* para proteger datos compartidos.

Puedes pensar en el runtime como el *"sistema operativo interno"* de Go, que hace que todo funcione sin que tengas que preocuparte por los detalles de bajo nivel.

## Scheduler

El scheduler (o planificador) es una parte del **runtime de Go** que se encarga de decidir **cuándo** y **dónde** se ejecuta cada **goroutine**. Es como un organizador que reparte el tiempo y los recursos del sistema entre todas las tareas.

➡️ **Otra forma de verlo:** El scheduler organiza y asigna las **goroutines** para que se ejecuten de manera eficiente sin necesidad de que el programador gestione manualmente los hilos del sistema.

Gracias al scheduler, Go puede manejar **miles de goroutines** de forma eficiente, incluso en computadoras con pocos núcleos.

## Goroutines

Cuando usas **goroutines** en Go:

- Go crea muchos **hilos ligeros**, llamados **goroutines**, que el **runtime** y el **scheduler** manejan de forma eficiente.
- Aunque tengas solo 1 o 2 núcleos, Go puede **intercalar** muchas goroutines (concurrencia).
- Si tienes varios núcleos, Go puede aprovecharlos para ejecutar goroutines realmente al mismo tiempo (paralelismo).
