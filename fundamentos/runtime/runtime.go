package main

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"sync"

	_ "github.com/lib/pq"
)

/*
* Introducción a runtime.AddCleanup (Go 1.24) 🧹
Go 1.24 introdujo `runtime.AddCleanup`, una mejora significativa sobre `runtime.SetFinalizer`.

- Los `finalizadores` son funciones que se ejecutan automáticamente cuando un objeto ya no es alcanzable (no se usa más).
- Se usan para liberar recursos como archivos, conexiones a bases de datos, memoria, etc.

🧪 Antes: `runtime.SetFinalizer`
- Solo permitía un finalizador por objeto.
- Podía provocar fugas de memoria en objetos cíclicos.
- Podía retrasar el recolector de basura (GC).
- No permitía limpiar correctamente objetos internos (interior pointers).

✅ Ahora: `runtime.AddCleanup`
- Permite múltiples tareas de limpieza por objeto.
- Compatible con referencias cíclicas.
- No bloquea el GC.
- Más seguro y flexible.

Sintaxis:
	runtime.AddCleanup(objeto, func() {
		- Código que se ejecuta cuando 'objeto' ya no se usa
	})

🔁 ¿Cuándo se ejecuta?
Cuando el objeto ya no se usa y el recolector de basura lo elimina.
*/

type Resource struct {
	id int
}

func basicExample() {
	fmt.Println("\n=== Ejemplo básico: Resource ===")

	// Creamos un recurso que queremos limpiar automáticamente
	res := &Resource{id: 42}

	// Asociamos una función de limpieza al recurso
	// Pasamos el snapshot explícitamente (res.id)
	runtime.AddCleanup(res, func(id int) {
		fmt.Println("🧼 Limpiando recurso:", id)
	}, res.id)

	// Simulamos que dejamos de usar el recurso
	res = nil

	// Forzamos la recolección de basura para ver el efecto (solo para demostración)
	runtime.GC()
	runtime.Gosched()
}

/*
📁 Ejemplo real: Limpieza automática de archivos

Muchas veces abrimos archivos pero olvidamos cerrarlos.
Con `runtime.AddCleanup`, podemos registrar una función que lo cierre automáticamente.
*/

func openFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("❌ Error al abrir el archivo:", err)
		return nil
	}

	// Creamos un snapshot con el nombre del archivo y el descriptor
	runtime.AddCleanup(f, func(fname string) {
		fmt.Println("📂 Cerrando archivo:", fname)
		f.Close() // f sigue accesible, pero ya no es fuerte referencia
	}, f.Name())

	return f
}

func fileExample() {
	fmt.Println("\n=== Ejemplo con archivos ===")
	file := openFile("example.txt")
	if file != nil {
		fmt.Println("✅ Archivo abierto:", file.Name())
	}
	file = nil
	runtime.GC()
}

/*
🗄️ Ejemplo real: Cierre de conexiones de base de datos

Las conexiones abiertas innecesariamente pueden saturar un servidor.
Aquí las cerramos automáticamente cuando el objeto DB deja de usarse.
*/

func openDatabase() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=test sslmode=disable")
	if err != nil {
		fmt.Println("❌ Error al abrir la base de datos:", err)
		return nil
	}

	runtime.AddCleanup(db, func(msg string) {
		fmt.Println(msg)
		db.Close()
	}, "🔌 Cerrando conexión a la base de datos")

	return db
}

func databaseExample() {
	fmt.Println("\n=== Ejemplo con base de datos ===")
	db := openDatabase()
	if db != nil {
		fmt.Println("✅ Conexión establecida")
	}
	db = nil
	runtime.GC()
}

/*
🧠 Ejemplo avanzado: Limpieza de caché

En aplicaciones grandes, los cachés mal manejados pueden causar consumo excesivo de memoria.
Con `runtime.AddCleanup`, podemos limpiar entradas automáticamente cuando ya no se usan.
*/

type Cache struct {
	mu    sync.Mutex
	items map[string]*string
}

func NewCache() *Cache {
	return &Cache{items: make(map[string]*string)}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = &value

	runtime.AddCleanup(&value, func(k string) {
		fmt.Println("🧹 Eliminando del caché:", k)
		c.mu.Lock()
		defer c.mu.Unlock()
		delete(c.items, k)
	}, key)
}

func cacheExample() {
	fmt.Println("\n=== Ejemplo con caché ===")
	cache := NewCache()
	cache.Set("user:123", "Alice")
	cache.items["user:123"] = nil
	runtime.GC()
}

/*
📊 Comparación rápida: SetFinalizer vs AddCleanup

| Característica                        | SetFinalizer | AddCleanup |
|--------------------------------------|--------------|------------|
| Múltiples funciones por objeto       | ❌ No        | ✅ Sí      |
| Soporte para ciclos de referencia    | ❌ No        | ✅ Sí      |
| Compatible con punteros internos     | ❌ No        | ✅ Sí      |
| Retarda el GC                        | ✅ Sí        | ❌ No      |
| Recomendado para nuevo código        | ❌ No        | ✅ Sí      |
*/

func main() {
	basicExample()
	fileExample()
	databaseExample()
	cacheExample()
}
