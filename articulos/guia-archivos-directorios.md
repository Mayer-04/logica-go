# Buenas prácticas para nombrar archivos y directorios en Go

En el lenguaje de programación Go, seguir una convención de nombres clara y coherente es crucial para la mantenibilidad y legibilidad del código. Estas prácticas están inspiradas en el estilo utilizado en los repositorios oficiales de Go en `GitHub`.

## Nombres de archivos de prueba (test)

Los archivos de prueba **(test files)** en Go deben seguir la convención `snake_case` y terminar con la palabra `test`. Esto facilita su identificación. **Por ejemplo:**

- **Correcto:** `image_test.go`
- **Incorrecto:** `ImageTest.go`

Esta convención facilita distinguir rápidamente los archivos de prueba de otros archivos de código en un proyecto.

## Nombres de archivos estándar

Para archivos que no son de prueba, es común escribir los nombres completamente en `minúsculas`, sin separar palabras. Aunque el uso de `snake_case` también es _aceptable_, la tendencia predominante es evitar caracteres de separación u otras convenciones _(camelCase)_ si el nombre es lo suficientemente corto.

Go en algunos casos, especialmente cuando se hace referencia a detalles de la plataforma o el sistema operativo, se utiliza la convención `snake_case`. Esto facilita la identificación de archivos específicos de una plataforma, versión o característica del sistema.

**Por ejemplo:**

- Para nombres cortos, lo recomendable es: `config.go`
- Para archivos específicos de plataformas o sistemas: `sys_unix.go`, `env_windows.go`, `exec_freebsd.go`

Se debe optar siempre por lo más legible y mantener consistencia dentro del proyecto.

# Nombres de directorios

Los nombres de los directorios en Go siempre se escriben en `minúsculas`, incluso cuando contienen más de una palabra. No se utilizan convenciones como _camelCase_ o _snake_case_ para nombrar directorios. Si el nombre es demasiado largo, se recomienda usar una `abreviatura` o `simplificación`.

**Por ejemplo:**

- **Correcto:** `validators/`, `httpclient/`
- **Incorrecto:** `Validator/`, `httpClient/`, `http_client/`

Esta convención sigue la guía de estilo oficial del equipo de Go y se observa en sus repositorios oficiales.

# Archivo principal dentro de directorios

Cuando se crea un **directorio** para un `paquete`, es una buena práctica que el **archivo principal** dentro de ese directorio tenga el mismo nombre que el directorio. Esto facilita la identificación del archivo que contiene el código principal del paquete.

**Por ejemplo:**

- Si el directorio se llama `validator/`, se recomienda que el archivo principal debería ser `validator.go`.

Esto ayuda a mantener una estructura coherente y predecible, lo que es útil para otros desarrolladores que naveguen por el código.
