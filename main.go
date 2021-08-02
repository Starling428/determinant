/*

для проверки можно использовать этот ресурс
https://matrix.reshish.ru/detCalculation.php
он позволяет вносить данные скопом в таком виде:

 3  -3  -5
-3   2   4
 2  -5  -7

*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func makeX(size int, value int) [][]int {
	rand.Seed(time.Now().UnixNano())
	newX := make([][]int, size)
	for i := 0; i < size; i++ {
		newX[i] = make([]int, size)
		for j := 0; j < size; j++ {
			newX[i][j] = -value + rand.Intn(2*value+1)
		}
	}
	return newX
}

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
			fmt.Printf("%5d", matrix[i][j])
		}
		fmt.Println()
	}
}

func deter2(matrix [][]int) int {
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
		var wg sync.WaitGroup
		var mu sync.Mutex
		wg.Add(len(matrix))
		for i := 0; i < len(matrix); i++ {
			go func(ind int) {
				defer wg.Done()
				mu.Lock()
				sum += deter(minor(matrix, ind)) * int(math.Pow(-1, float64(ind))) * matrix[0][ind]
				mu.Unlock()
				// fmt.Println(sum)
			}(i)
		}
		wg.Wait()
		return sum
	} else {
		return deter2(matrix)
	}
}

func main() {
	randX := makeX(12, 10)
	showMatrix(randX)
	fmt.Println("\nΔ =", deter(randX))
}
