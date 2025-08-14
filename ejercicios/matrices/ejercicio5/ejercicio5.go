package main

import "fmt"

func main() {
	// matriz := [][]int{
	// 	{2, 3, 4},
	// 	{5, 6, 7},
	// 	{8, 9, 10},
	// }

	matriz := make([][]int, 3)
	for i := range matriz {
		matriz[i] = make([]int, 3)
		for j := range matriz[i] {
			if j == i {
				matriz[i][j] = 1
			} else {
				matriz[i][j] = 0
			}
		}
	}

	for _, fila := range matriz {
		fmt.Println(fila)
	}
}
