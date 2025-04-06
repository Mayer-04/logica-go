# Conceptos importantes de Go

## Concurrencia

La concurrencia es la capacidad de un programa para manejar múltiples tareas alternando su ejecución sin que necesariamente ocurran en paralelo. Es decir, las tareas pueden **turnarse** en su ejecución en lugar de correr todas al mismo tiempo.

**Ejemplos cotidianos:**

Imagina que una sola persona está cocinando:

- Primero corta las verduras 🥕
- Luego espera a que el agua hierva 💦
- Mientras tanto, saltea algunos ingredientes 🍳

También puedes pensar en una persona que atiende varias llamadas telefónicas. Aunque solo puede hablar con una persona a la vez, cambia rápidamente de una llamada a otra, dando la impresión de que atiende a varias personas simultáneamente.

➡️ Lidiar con muchas cosas a la vez, pero de manera alternada.

## Paralelismo

El paralelismo consiste en ejecutar múltiples tareas simultáneamente (al mismo tiempo), aprovechando múltiples núcleos o procesadores para mejorar el rendimiento.

**Ejemplos cotidianos:**

Imagina que varias personas están en la cocina, cada una realizando una tarea al mismo tiempo en diferentes estaciones 👨‍🍳👩‍🍳.

Piensa en una central de llamadas donde varias personas atienden llamadas al mismo tiempo, cada una en una línea diferente 📲.

➡️ Hacer muchas cosas al mismo tiempo.

## Runtime de Go

Es el entorno de ejecución de Go, un software que administra la ejecución de los programas escritos en este lenguaje. Se encarga de aspectos clave como la **gestión de memoria**, **la ejecución de goroutines**, la recolección de basura (garbage collector) y la comunicación entre procesos mediante *channels* y *locks*.

## Scheduler

El scheduler es una parte del `runtime de Go` que administra la **ejecución de las goroutines** de manera eficiente. Se encarga de distribuirlas sobre los hilos (threads) del sistema operativo, decidiendo cuándo y cómo se ejecutan para optimizar el uso de los recursos disponibles.

➡️ **Otra forma de verlo:** El scheduler organiza y asigna las **goroutines** para que se ejecuten de manera eficiente sin necesidad de que el programador gestione manualmente los hilos del sistema.
