package main

import "fmt"

// La estructura 'Queue' acepta un tipo genérico que puede ser 'string' o 'int'.
type Queue[T string | int] struct {
	items []T
	size  int
}

func (q *Queue[T]) enqueue(item T) {
	q.items = append(q.items, item)
	q.size++
}

func (q *Queue[T]) dequeue() T {
	firstItem := q.items[0]
	q.items = q.items[1:]
	q.size--
	return firstItem
}

func (q *Queue[T]) peek() T {
	firstItem := q.items[0]
	return firstItem
}

func (q *Queue[T]) isEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) sizeQueue() int {
	return q.size
}

func (q *Queue[T]) clear() {
	q.items = make([]T, 0)
	q.size = 0
	// clear(q.items)
}

func (q *Queue[T]) DisplayQueue() {
	fmt.Printf("%#v\n", q.items)
}

func main() {
	var queue Queue[int]

	// Agregar elementos a la cola
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)

	// Mostrar la cola
	queue.DisplayQueue()

	// Eliminar elementos de la cola
	queue.dequeue()
	queue.dequeue()

	// Mostrar la cola actualizada
	queue.DisplayQueue()

	// Obtener el primer elemento de la cola
	fmt.Println("Primer elemento:", queue.peek())

	// Obtener el tamaño de la cola
	fmt.Println("Tamaño de la cola:", queue.sizeQueue())

	// Limpiar la cola
	queue.clear()
	fmt.Println("cola vacía:", queue.isEmpty())

	// Mostrar la cola vacía
	queue.DisplayQueue()

	// Agregar nuevos elementos a la cola
	queue.enqueue(4)
	queue.enqueue(5)

	// Mostrar la cola actualizada
	queue.DisplayQueue()

	// Nuevo tamaño de la cola
	fmt.Println("Tamaño de la cola:", queue.sizeQueue())
}
