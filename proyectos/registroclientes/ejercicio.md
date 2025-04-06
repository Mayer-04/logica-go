# Sistema de Registro de Clientes

Se solicita implementar un sistema de registro de clientes para una empresa, utilizando una lista para almacenar la información de cada cliente.

Cada cliente debe contar con los siguientes datos:

- ID (identificador único)
- Nombre
- Dirección
- Teléfono

## Requisitos

1. **Cantidad máxima de clientes:** Define una `constante` que indique el número máximo de clientes que se podrán almacenar en el sistema.
2. **Menú de opciones:** Implementa un menú interactivo que permita al usuario realizar las siguientes acciones:

- **Registrar clientes:** Permite al usuario ingresar un nuevo cliente, solicitando todos los datos requeridos (id, nombre, dirección y número de teléfono).

- **Listar clientes:** Muestra en pantalla la información de todos los clientes registrados.
- **Buscar cliente:** Permite buscar clientes por **ID o nombre**, mostrando sus datos si existen en la lista.
- **Eliminar cliente:** Elimina a un cliente de la lista según un criterio específico **(por ejemplo, su id o nombre)**.
- **Salir:** Termina la ejecución del programa.

## Validaciones

Se deben aplicar las siguientes validaciones al ingresar datos:

- Valida que no se puedan agregar más clientes de los permitidos por el límite máximo establecido en la *constante*.
- **ID:** Debe ser un número entero único.
- **Nombre:** No puede estar vacío ni contener números o caracteres especiales.
- **Dirección:** No puede estar vacía.
- **Teléfono:** Debe contener solo dígitos y cumplir con un formato válido (por ejemplo, una longitud mínima de 7 dígitos y máxima de 15).
