package main

import (
	"fmt"
	"math/rand"
	"time"
)

// using Silvester's formula
// this function works only with powers of 2
func hadamardMatrix(n int) [][]int {

	if n == 1 {
		return [][]int{{1}}
	}

	h := hadamardMatrix(n / 2)
	res := make([][]int, n)

	for i := range res {
		res[i] = make([]int, n)
	}

	// [ H  H ]
	// [ H -H ]
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			res[i][j] = h[i][j]          //upper left part
			res[i][j+n/2] = h[i][j]      //upper right part
			res[i+n/2][j+n/2] = -h[i][j] //down right
			res[i+n/2][j] = h[i][j]      //down left

		}
	}

	return res
}

// Генериране на Пейли матрица за p = 3 (3x3)
func payleyBase() [][]int {
	return [][]int{
		{1, 1, 1},
		{1, -1, 1},
		{1, 1, -1},
	}
}

// Функция за разширяване на матрицата с помощта на Пейли
func payley(size int) [][]int {
	// Проверка дали размерът е кратен на 3
	if size%3 != 0 {
		fmt.Println("Размерът трябва да е кратен на 3!")
		return nil
	}

	// Генерираме основната 3x3 Пейли матрица
	base := payleyBase()

	// Изчисляваме колко пъти ще комбинираме 3x3 матрици
	blocks := size / 3

	// Създаваме новата матрица с размер size x size
	result := make([][]int, size)
	for i := range result {
		result[i] = make([]int, size)
	}

	// Запълваме резултата, като комбинираме 3x3 матрици
	for i := 0; i < blocks; i++ {
		for j := 0; j < blocks; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					result[i*3+k][j*3+l] = base[k][l]
				}
			}
		}
	}

	return result
}

// H20 Адамарова матрица за размер 20
var H20 = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
	{1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1},
	{1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1},
	{1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1},
	{1, -1, 1, -1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1},
	{1, 1, -1, -1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1},
	{1, -1, -1, 1, -1, 1, 1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, -1, 1, 1},
	{1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, 1, 1, 1, 1, 1, 1, -1, -1},
	{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
	{1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1},
	{1, -1, -1, 1, -1, 1, 1, -1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1, 1, -1, -1},
	{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
	{1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1},
	{1, -1, -1, 1, -1, 1, 1, -1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, 1, 1, 1, 1},
	{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
	{1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1},
	{1, -1, -1, 1, -1, 1, 1, -1, 1, -1, 1, -1, -1, 1, 1, -1, 1, 1, -1, 1},
}

// H24 Адамарова матрица за размер 24
var H24 = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
	{1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1},
	{1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1},
	{1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1, -1, -1, -1, -1},
	{1, -1, 1, -1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1},
	{1, 1, -1, -1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1},
	{1, -1, -1, 1, -1, 1, 1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, -1, 1, 1, 1, -1, -1, 1},
	{1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1},
	{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
	{1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1},
	{1, -1, -1, 1, -1, 1, 1, -1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1},
	{1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1, 1, -1},
	{1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1, 1, 1, -1, -1},
	{1, -1, -1, 1, -1, 1, 1, -1, 1, -1, 1, -1, -1, 1, 1, -1, 1, 1, -1, 1, 1, -1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, 1, 1, 1, 1, 1, 1, 1, 1},
}

// Хамингово разстояние между две кодови думи
// Функция за изчисляване на Хемингова дистанция между две бинарни вектора
func hammingDistance(v1, v2 []int) int {
	distance := 0
	for i := 0; i < len(v1); i++ {
		if v1[i] != v2[i] {
			distance++
		}
	}
	return distance
}

// Функция за намиране на минималната Хемингова дистанция между всички кодови думи
func minHammingDistance(codewords [][]int) int {
	minDistance := len(codewords[0]) // Начално разстояние е максималното
	for i := 0; i < len(codewords); i++ {
		for j := i + 1; j < len(codewords); j++ {
			distance := hammingDistance(codewords[i], codewords[j])
			if distance < minDistance {
				minDistance = distance
			}
		}
	}
	return minDistance
}

// Функция за генериране на случайна кодова дума с дължина n
func generateRandomCodeword(n int) []int {
	codeword := make([]int, n)
	for i := range codeword {
		codeword[i] = rand.Intn(2)*2 - 1 // генерира -1 или 1
	}
	return codeword
}

// Функция за създаване на нелинейни кодове с минимална Хаминг дистанция
func generateNonLinearCodes(n, minDistance, numCodes int) [][]int {
	rand.Seed(time.Now().UnixNano())
	var codes [][]int

	for len(codes) < numCodes {
		// Генериране на нова случайна кодова дума
		newCode := generateRandomCodeword(n)
		valid := true

		// Проверка за минималната Хаминг дистанция спрямо съществуващите кодови думи
		for _, code := range codes {
			if hammingDistance(newCode, code) < minDistance {
				valid = false
				break
			}
		}

		// Ако новата кодова дума е валидна, я добавяме към списъка
		if valid {
			codes = append(codes, newCode)
		}
	}

	return codes
}
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
func ConvertHadamardToBinary(matrix [][]int) [][]int {
	binaryMatrix := make([][]int, len(matrix))
	for i := range matrix {
		binaryMatrix[i] = make([]int, len(matrix[i]))
		for j, val := range matrix[i] {
			if val == 1 {
				binaryMatrix[i][j] = 0
			} else if val == -1 {
				binaryMatrix[i][j] = 1
			} else {
				panic(fmt.Sprintf("Невалиден елемент в матрицата на Адамар: %d", val))
			}
		}
	}
	return binaryMatrix
}

func printResults() {
	n := 16
	matrix16 := hadamardMatrix(n)
	matrix12 := payley(12)

	fmt.Println("Hadamard Matrix 16x16:")
	printMatrix(matrix16)
	fmt.Println("Нелинейни кодове:")
	ConvertHadamardToBinary(matrix16)
	nonLinearCodes16 := generateNonLinearCodes(n, minHammingDistance(matrix16), 8)
	printMatrix(nonLinearCodes16)
	fmt.Println("Hadamard Matrix 12x12:")
	printMatrix(matrix12)
	fmt.Println("Нелинейни кодове:")
	ConvertHadamardToBinary(matrix12)
	nonLinearCodes12 := generateNonLinearCodes(12, minHammingDistance(matrix12), 12)
	printMatrix(nonLinearCodes12)
	fmt.Println("\nHadamard Matrix 20x20:")
	printMatrix(H20)
	fmt.Println("Нелинейни кодове:")
	ConvertHadamardToBinary(H20)
	nonLinearCodes20 := generateNonLinearCodes(20, minHammingDistance(H20), 20)
	printMatrix(nonLinearCodes20)
	fmt.Println("\nHadamard Matrix 24x24:")
	printMatrix(H24)
	fmt.Println("Нелинейни кодове:")
	ConvertHadamardToBinary(H24)
	nonLinearCodes24 := generateNonLinearCodes(24, minHammingDistance(H24), 24)
	printMatrix(nonLinearCodes24)
}
