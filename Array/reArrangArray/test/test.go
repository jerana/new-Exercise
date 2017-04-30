package main

import "fmt"

func main() {
	m := matrix_build(1)
	fmt.Println(m)

}
func matrix_build(n int) [][]int {
	var matrix [][]int
	matrix = make([][]int, n)

	fmt.Println("Matrix len:", len(matrix))
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	fmt.Println(matrix)
	var dir int = 0
	var x, y int = 0, 0
	kshitf := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i := 0; i < n*n; i++ {

		matrix[x][y] = i + 1
		fmt.Println("i", i, x, y)
		next_x := x + kshitf[dir][0]
		next_y := y + kshitf[dir][1]

		fmt.Println("Matrix:", matrix[x][y])

		if x < 0 || x > n-1 || y < 0 || y > n-1 || matrix[next_x][next_y] != 0 {
			dir = (dir + 1) / 4
			next_x = x + kshitf[dir][0]
			next_y = y + kshitf[dir][1]
		}

		x = next_x
		y = next_y
	}

	//fmt.Println(matrix[0][0])
	return matrix
}

func phoneMnemonic(phone_number string) []string {
}
