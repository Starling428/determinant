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


type sum struct {
	amount int
	sync.Mutex
}

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
/*func showMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print("{")
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%5d", matrix[i][j])
			fmt.Print(",")
		}
		fmt.Print("},")
		fmt.Println()
	}
}*/

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
		count := new(sum)
		var wg sync.WaitGroup
		wg.Add(len(matrix))
		for i := 0; i < len(matrix); i++ {
			go func(ind int) {
				defer wg.Done()
				defer count.Unlock()
				count.Lock()				
				count.amount += deter(minor(matrix, ind)) * int(math.Pow(-1, float64(ind))) * matrix[0][ind]	
			}(i)
		}
		wg.Wait()
		return count.amount
	} else {
		return deter2(matrix)
	}
}

/*func deter(matrix [][]int) int {
	if len(matrix) > 2 {
		var sum int
		var wg sync.WaitGroup
		var mu sync.Mutex
		wg.Add(len(matrix))
		for i := 0; i < len(matrix); i++ {
			go func(ind int) {
				defer wg.Done()
				mu.Lock()
				sum += deter(minor(matrix, i)) * int(math.Pow(-1, float64(i))) * matrix[0][i]
				mu.Unlock()
				fmt.Println(sum)
			}(i)
		}
		wg.Wait()
		return sum
	} else {
		return deter2(matrix)
	}
}*/

func main() {
	//randX := makeX(11, 10)
	//showMatrix(randX)
	test := [][]int {
		{   -9,    4,   -1,   -6,   -2,   -6,   -7,    7,   -8,    4,  },// -1,   2},
		{    3,   -5,    3,   -2,    3,   10,   -4,    8,    6,    0,  },//  4,   5},
		{    7,   -3,    4,   10,    7,   -6,   -4,    7,   -7,    8,  },//  2,  -2},
		{    4,   -7,  -10,    6,  -10,  -10,   -1,    0,    3,    9,  },//  3,   9},
		{   -4,    0,  -10,    2,    7,   -4,   -4,    7,    5,    4,  },// -6,   0},
		{    7,    5,    2,   -7,    0,    2,   -2,    1,    0,   -6,  },//-10,  -7},
		{    6,   -6,    5,    2,    3,    8,   -6,   -2,    1,    4,  },//  1,   2},
		{    2,    3,    2,  -10,    3,    1,    5,    8,    8,    9,  },// -2,   5},
		{    1,   -8,    9,    1,   -6,   -9,  -10,    7,   -5,    9,  },// -7,  -8},
		{    0,   -3,   -7,    7,    1,   -9,    4,   -4,   10,   -7,  },//  3,   0},
		// {   10,   10,   -6,  -10,   10,   -1,    1,    6,   -8,    4,    3,  },//  1},
		//{   11,    9,   -3,    5,    0,   -9,    6,    2,   -7,    3,   -5,  }// -9},
	}
	showMatrix(test)
	t0 := time.Now()
	fmt.Println("\nΔ =", deter(test))
	t1 := time.Now()
	fmt.Printf("Elapsed time: %v", t1.Sub(t0))
}
