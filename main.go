/*

считает пока без горутин, для тестов внёс матрицы 2, 3, 4 и 7 порядков.
для проверки можно использовать этот ресурс
https://matrix.reshish.ru/detCalculation.php
он позволяет вносить данные скопом в таком виде:

3  -3  -5  8  -3  3  8
-3  2  4  -6  -3  -5  3
2  -5  -7  5  -4  -1  8
-4  3  5  -6  -3  -2  -6
3  -3  -1  8  -2  7  2
-3  6  4  -6  -3  -5  -3
7  -5  -7  5  9  -5  8

*/

package main

import (
	"fmt"
	"math"
)

func dupl(matrix [][]int) [][]int {
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func showMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%9d", matrix[i][j])
		}
		fmt.Println()
	}
}

func deter2x2(matrix [][]int) int {
	return matrix[0][0]*matrix[1][1] - matrix[1][0]*matrix[0][1]
}

func minor(matrixOrig [][]int, colNum int) [][]int {
	matrix := dupl(matrixOrig)
	if len(matrix) <= 2 {
		return matrix
	}
	matrix = matrix[1:]
	for i := 0; i < len(matrix); i++ {
		matrix[i] = append(matrix[i][:colNum], matrix[i][colNum+1:]...)
	}
	return matrix
}

func deter(matrix [][]int) int {
	if len(matrix) > 2 {
		var sum int
		for i := 0; i < len(matrix); i++ {
			// showMatrix(minor(matrix, i))
			// fmt.Println()
			// fmt.Println(deter(minor(matrix, i)) * int(math.Pow(-1, float64(i))) * matrix[0][i])
			// fmt.Println()
			sum += deter(minor(matrix, i)) * int(math.Pow(-1, float64(i))) * matrix[0][i]
		}
		return sum
	} else {
		return deter2x2(matrix)
	}
}

func main() {
	// x := [][]int{
	// 	{11, -3},
	// 	{-15, -2},
	// }
	// y := [][]int{
	// 	{1, -2, 3},
	// 	{4, 0, 6},
	// 	{-7, 8, 9},
	// }
	// z := [][]int{
	// 	{3, -3, -5, 8},
	// 	{-3, 2, 4, -6},
	// 	{2, -5, -7, 5},
	// 	{-4, 3, 5, -6},
	// }
	s := [][]int{
		{3, -3, -5, 8, -3, 3, 8},
		{-3, 2, 4, -6, -3, -5, 3},
		{2, -5, -7, 5, -4, -1, 8},
		{-4, 3, 5, -6, -3, -2, -6},
		{3, -3, -1, 8, -2, 7, 2},
		{-3, 6, 4, -6, -3, -5, -3},
		{7, -5, -7, 5, 9, -5, 8},
	}
	/*
		fmt.Println(deter2x2(x))
		showMatrix(y)
		fmt.Println()
		showMatrix(minor(y, 0))
		fmt.Println(deter(x))
		fmt.Println()
		showMatrix(y)
		fmt.Println(y)

		fmt.Println("y")
		showMatrix(y)
		dupy := dupl(y)
		fmt.Println()

		fmt.Println("dupy")
		showMatrix(dupy)
		fmt.Println()

		fmt.Println("dupy minor")
		showMatrix(minor(dupy, 2))
		fmt.Println()

		fmt.Println("dupy")
		showMatrix(dupy)

		fmt.Println()
		fmt.Println("y")
		showMatrix(y)
	*/

	showMatrix(s)
	fmt.Println("\ndet(z) =", deter(s))
}
