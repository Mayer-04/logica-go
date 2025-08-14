# Variables de entorno en Go

1. Se recomienda usar un archivo `.env` durante el desarrollo para simular el entorno de producción.
2. En producción, estas variables usualmente se inyectan directamente en el contenedor Docker.
3. Aunque puedes acceder directamente a variables de entorno desde cualquier parte del código usando `os.Getenv`, eso no es recomendable.
4. Las variables de entorno son globales, y el uso descontrolado de variables globales es una mala práctica porque:

- Hace difícil rastrear de dónde vienen los valores.
- Complica las pruebas y el mantenimiento del código.

## Buenas practicas

- Cargar todas las variables de entorno una sola vez, al inicio del programa.
- Guardarlas en una estructura (struct) llamada `Config`.
- Luego, pasar esta configuración como parámetro a las partes del código que la necesiten.
