# Cómo nombrar archivos y directorios en Go: buenas prácticas y convenciones

En el lenguaje de programación Go, adoptar una convención de nombres clara y coherente es fundamental para garantizar la mantenibilidad y la legibilidad del código. Estas prácticas se basan en el estilo utilizado en los repositorios oficiales de Go en `GitHub`.

## Convenciones para archivos de prueba en Go (test)

Los archivos de prueba **(test files)** en Go deben seguir la convención `snake_case` y terminar con la palabra `test`. Esto facilita su identificación.

**Por ejemplo:**

- **Correcto:** `image_test.go`
- **Incorrecto:** `ImageTest.go`

Esta convención facilita distinguir rápidamente los archivos de prueba de otros archivos de código en un proyecto.

## Nombres de archivos estándar

Para archivos que no son de prueba, es común escribir los nombres completamente en `minúsculas`, sin separación entre palabras. Aunque el uso de `snake_case` es _aceptable_, la práctica más habitual es evitar caracteres de separación u otras convenciones como **camelCase**, especialmente cuando el nombre es lo suficientemente corto.

En algunos casos, particularmente al hacer referencia a detalles específicos de la plataforma o del sistema operativo, Go utiliza la convención `snake_case`. Esto facilita la identificación de archivos relacionados con una plataforma, versión o característica del sistema.

**Por ejemplo:**

- Para nombres cortos, lo recomendable es: `config.go`
- Para archivos específicos de plataformas o sistemas: `sys_unix.go`, `env_windows.go`, `exec_freebsd.go`

Se debe optar siempre por lo más legible y mantener consistencia dentro del proyecto.

## Cómo nombrar directorios en Go: minúsculas y sin separadores

Los nombres de los directorios en Go siempre se escriben en `minúsculas`, incluso cuando contienen más de una palabra. No se emplean convenciones como **camelCase** ni **snake_case** para nombrar directorios. Si el nombre resulta demasiado largo, se recomienda utilizar una `abreviatura` o una versión simplificada.

**Por ejemplo:**

- **Correcto:** `validators/`, `httpclient/`
- **Incorrecto:** `Validator/`, `httpClient/`, `http_client/`

Esta convención sigue la guía de estilo oficial del equipo de Go y se observa en sus repositorios oficiales.

## Nombra el archivo principal igual que el directorio del paquete

Al crear un directorio para un paquete, es una buena práctica que el `archivo principal` dentro de ese directorio tenga el mismo nombre que el directorio. Esto facilita la identificación del archivo que contiene el código principal del paquete.

**Ejemplo:**

- Si el directorio se llama `validator/`, el archivo principal debería llamarse `validator.go`.

Seguir esta convención ayuda a mantener una estructura coherente y predecible, lo cual resulta especialmente útil para otros desarrolladores al explorar el código.
