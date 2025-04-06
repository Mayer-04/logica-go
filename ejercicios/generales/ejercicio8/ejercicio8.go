package main

import (
	"fmt"
	"strconv"
)

/*
Escribe una función que convierta un slice de enteros en un slice de strings.
La función debería devolver el nuevo slice.
*/

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := convertToStringSlice(arr)
	fmt.Println(result)
}

func convertToStringSlice(arr []int) []string {
	strArr := make([]string, len(arr))
	for i, value := range arr {
		strArr[i] = strconv.Itoa(value)
	}
	return strArr
}
