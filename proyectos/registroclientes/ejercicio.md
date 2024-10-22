# Registro de clientes

Se solicita implementar un sistema de registro de clientes para una empresa, utilizando una lista para almacenar la información de cada cliente. Cada cliente deberá tener los siguientes campos: `id`, `nombre`, `dirección` y `teléfono`.

## Requisitos

1. **Cantidad máxima de clientes:** Define una `constante` que indique el número máximo de clientes que se podrán almacenar en el sistema.
2. **Menú de opciones:** Implementa un menú interactivo que permita al usuario realizar las siguientes acciones:

- **Agregar cliente:** Permite al usuario registrar un nuevo cliente ingresando su id, nombre, dirección y número de teléfono.
- **Mostrar clientes:** Muestra la lista completa de los clientes registrados, junto con su información.
- **Eliminar cliente:** Elimina a un cliente de la lista según un criterio específico **(por ejemplo, su id o nombre)**.
- **Salir:** Termina la ejecución del programa.

## Validaciones

Debes realizar las siguientes validaciones al ingresar datos:

- Valida que no se puedan agregar más clientes de los permitidos por el límite máximo establecido.
- **Nombre:** Asegúrate de que no esté vacío y no contenga números ni caracteres especiales.
- **Dirección:** Verifica que no esté vacío.
- **Teléfono:** Valida que el número ingresado siga un formato adecuado y solo contenga dígitos.
