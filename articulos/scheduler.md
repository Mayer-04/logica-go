# Scheduler del runtime de Go (planificador en tiempo de ejecución de Go)

El runtime de Go es el `entorno de ejecución` que gestiona aspectos fundamentales como la **asignación de memoria**, **la recolección de basura** y **la concurrencia**.

En Go, las `goroutines` son unidades ligeras de ejecución que permiten la concurrencia de manera eficiente. En este artículo, te explico cómo el `scheduler` del runtime de Go gestiona estas goroutines.

## Gestión de goroutines

El `scheduler` en Go es el encargado de gestionar la ejecución de las goroutines de manera concurrente. En lugar de asignar cada goroutine a un hilo físico del sistema operativo, el scheduler agrupa muchas goroutines para que utilicen unos pocos hilos de manera eficiente.

> [!IMPORTANT]
> El scheduler utiliza un modelo conocido como `GMP`.

### Modelo GMP: Goroutine, Machine, Processor

Cada goroutine (G) se ejecuta en un hilo del sistema operativo (M) que se asigna a un procesador lógico o CPU lógica (P).

Go utiliza un modelo conocido como `GMP` para gestionar la ejecución de las goroutines:

- **G (goroutine):** Es una tarea ligera que ejecuta código en Go.
- **M (machine):** Representa un hilo del sistema operativo.
- **P (processor):** Es una estructura interna que gestiona la ejecución de goroutines. Cada `P` se asocia con un `M` y permite que las goroutines se ejecuten en los hilos disponibles. `P` se encarga de asignar y distribuir las tareas a los **hilos (M)** para que el trabajo se realice de manera eficiente.

_El número de M (hilos del sistema operativo) suele ser menor que el número de G (goroutines), lo que permite que muchas goroutines se ejecuten de manera eficiente utilizando pocos hilos._

### Ejemplos de analogía para el modelo GMP

1. **100 tareas y 10 trabajadores:** Imagina que tienes 100 tareas (goroutines) pero solo 10 trabajadores (hilos). Los trabajadores se turnan para completar las tareas, mientras un supervisor (el P) se asegura de que las tareas sean distribuidas equitativamente entre ellos.

2. Imagina que tienes una cocina con `100 recetas (las goroutines)` que preparar, pero solo `10 cocineros (los hilos)`. En lugar de asignar un cocinero a cada receta, esos 10 cocineros se turnan para trabajar en las 100 recetas. Cada cocinero tiene acceso a `estaciones de trabajo (los P)`, que son espacios donde pueden preparar las recetas. Las estaciones de trabajo organizan y distribuyen las recetas a los cocineros, asegurando que todo se haga de manera eficiente. Si un cocinero se bloquea esperando un ingrediente (una operación bloqueante), otro cocinero puede usar la misma estación de trabajo para continuar con otra receta.

## Colas locales y cola global en el scheduler

Las goroutines se colocan en colas de espera antes de ejecutarse. Existen colas locales asociadas a cada `P` y una cola global. Las goroutines en las colas locales son las primeras en ser ejecutadas, y si un `P` no tiene más goroutines en su cola, puede tomar una de la cola global o _"robar"_ una goroutine de otra cola local. Este mecanismo asegura que el trabajo se distribuya equitativamente entre los procesadores `P`.

### Colas locales

Cada procesador `P` en Go tiene su propia cola local, donde almacena las goroutines que están listas para ejecutarse. Estas colas locales son accesibles exclusivamente por el `P` correspondiente, lo que significa que el acceso y la gestión de estas goroutines son rápidos y no requieren sincronización entre múltiples hilos.

### Cola global

La cola global es una estructura compartida por todos los `P (procesadores)` en el runtime de Go. Cuando un P tiene más goroutines de las que puede manejar en su cola local, puede mover algunas a la cola global para que otros `P` puedan utilizarlas. Si un `P` se queda sin goroutines en su cola local y no puede robar trabajo de otros `P`, buscará en la cola global. La cola global actúa como un respaldo cuando los `P` necesitan trabajo extra o tienen exceso de goroutines.

## Asignación a hilos

Cada procesador `P` tiene su propia cola de goroutines y toma un hilo del sistema operativo `M` para ejecutar esas goroutines. Si una goroutine realiza una operación bloqueante, el hilo `M` que estaba ejecutando esa goroutine puede ser liberado por el scheduler y reasignado a otra goroutine que esté lista para ejecutarse, optimizando así el uso de los recursos.

Esto significa que no hay una relación fija entre un hilo `M` y una goroutine `G`. El scheduler gestiona dinámicamente qué hilos ejecutan qué goroutines, permitiendo que muchas goroutines compartan pocos hilos y que los hilos estén siempre ocupados en tareas útiles.

## Operaciones bloqueantes

Cuando una goroutine está bloqueada esperando una _operación de red o disco_, el `scheduler` puede reasignar el hilo a otra goroutine que esté lista para ejecutarse.

> [!NOTE]
> El scheduler simplemente usa ese hilo del sistema operativo para ejecutar otra goroutine.

Una vez que la operación bloqueante se completa, el scheduler decide cuándo reanudar la ejecución de la goroutine bloqueada.

## Context switching (cambio de contexto)

Como las goroutines son ligeras, el `cambio de contexto` entre ellas (el proceso de detener una goroutine y comenzar a ejecutar otra) es mucho más rápido y eficiente en comparación con el `cambio de contexto` entre hilos del sistema operativo. El `scheduler` de Go optimiza estos cambios para reducir el tiempo que las goroutines pasan esperando para ser ejecución.
