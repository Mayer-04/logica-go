# Pointers

**Pointers** in Go are an essential feature that allows working directly with a programs memory. A pointer stores the memory address of another variable. This can be useful for accessing and modifying the value stored at that specific memory address.

## Pointer declaration

In Go, a pointer is declared using the dereferencing operator `(*)` before the data type to indicate that the variable is a pointer to that type of data. Then, to assign it the memory address of another variable, the `&` operator is used:

```go
var x int = 7
var ptr *int = &x

fmt.Println(ptr) // Output: 0xc000096068
```

In this example, `ptr` is a pointer to an integer `(*int)` that stores the memory address of the variable `x`.

## Accessing value through a pointer

To access the value stored at the memory address pointed to by a pointer, you use the `*` operator (`dereferencing operator`).

```go
fmt.Println(*ptr) // Output: 7
```

This will print the value stored at the memory address pointed to by ptr.

## Modifying value through a pointer

Pointers allow you to indirectly modify the value of a variable by accessing the memory address of that variable.

```go
*ptr = 10
```

This will change the value of the variable `x` to 20, as `ptr` points to the memory address of `x`.

## Checking for null pointers

In Go, pointers are automatically initialized with the null value `(nil)`. It’s important to check if a pointer is null before attempting to access the value it points to. Attempting to dereference a null pointer will result in a runtime panic, which may cause the program to stop abruptly:

```go
var ptr *int
fmt.Println(ptr) // Output: <nil>
```

## Example

Here’s a comprehensive example that declares, initializes, accesses, and modifies values through pointers in Go:

```go
package main

import "fmt"

func main() {
 x := 7
 var ptr *int = &x

 fmt.Println("Value of x:", x)              // Output: 7
 fmt.Println("Memory address of x:", ptr)   // Output: 0xc00000a0f8
 fmt.Println("Value pointed by ptr:", *ptr) // Output: 7

 *ptr = 10
 fmt.Println("New value of x:", x) // Output: 10
}
```
