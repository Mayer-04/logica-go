package main

/*
* sync.RWMutex
Es una estructura del paquete `sync` que implementa un mutex de lectura/escritura.
- Es fundamental para concurrencia segura en donde varias goroutines acceden simultaneamente
a un recurso compartido.
- sync.RWMutex significa Read-Write Mutex.
- RWMutex distingue entre lectores y escritores.
- Permite múltiples lectores simultáneamente.
- Permite solo un escritor a la vez, bloqueando tanto lectores como a otros escritores mientras escribe.


* Métodos
- RLock(): Adquiere el candado de lectura. Muchos pueden leer a la vez.
- RUnlock(): Libera el candado de lectura.
- Lock(): Adquire el candado de escritura. Solo uno puede escribir, y nadie puede leer ni escribir mientras tanto.
- UnLock(): Libera el candado de escritura.
*/

func main() {

}
