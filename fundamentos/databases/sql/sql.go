package main

/*
* Paquete database/sql:
- Proporciona una interfaz genérica para interactuar con bases de datos SQL en Go.
- No implementa directamente la comunicación con la base de datos; en su lugar, depende de los "drivers específicos".
- Permite realizar operaciones como abrir conexiones, ejecutar consultas y manejar transacciones de manera estándar.

* Paquete driver:
- Define las interfaces que deben ser implementadas por los drivers específicos.
- No se usa directamente en el código de las aplicaciones; en su lugar,
los drivers específicos lo implementan para integrarse con `database/sql`.

* Drivers específicos:
- Aunque `database/sql` proporciona una API común, se necesita un driver específico para cada base de datos.
- Estos drivers son paquetes externos desarrollados por la comunidad.
* Ejemplos de drivers populares:
 - PostgreSQL: `github.com/lib/pq` o `github.com/jackc/pgx/v5`
 - MySQL: `github.com/go-sql-driver/mysql`
 - SQLite: `github.com/mattn/go-sqlite3`
 - Microsoft SQL Server: `github.com/microsoft/go-mssqldb`
*/

func main() {

}
