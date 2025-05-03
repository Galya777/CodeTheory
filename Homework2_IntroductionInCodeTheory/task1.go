package main

import (
	"fmt"
	"math/rand"
	"time"
)

func truthTableToZhegalkin(y []int) []int {
	n := len(y)
	a := make([]int, n)
	copy(a, y)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if (i & j) == j {
				a[i] ^= a[j]
			}
		}
	}
	return a
}

func zhegalkinToTruthTable(y []int) []int {
	n := len(y)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if (i & j) == j {
				sum ^= y[j]
			}
		}
		a[i] = sum
	}
	return a
}

func generateBooleanVector(m int) []int {
	n := 1 << m
	vec := make([]int, n)
	for i := 0; i < n; i++ {
		vec[i] = rand.Intn(2)
	}
	return vec
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализира генератора на случайни числа

	var m int
	fmt.Print("Enter number of variables (m): ")
	fmt.Scan(&m)

	// Тестов вектор за проверка: f(x1, x2) = x1 XOR x2 -> [0, 1, 1, 0]
	if m == 2 {
		fmt.Println("\nTesting with example: f(x1, x2) = x1 XOR x2")
		testVector := []int{0, 1, 1, 0}
		zheg := truthTableToZhegalkin(testVector)
		reconstructed := zhegalkinToTruthTable(zheg)
		fmt.Println("Original (truth table):", testVector)
		fmt.Println("Zhegalkin polynomial:   ", zheg)
		fmt.Println("Reconstructed table:    ", reconstructed)
		fmt.Println()
	}

	// Генерирай произволен вектор и провери
	y := generateBooleanVector(m)
	zheg := truthTableToZhegalkin(y)
	reconstructed := zhegalkinToTruthTable(zheg)

	fmt.Println("Randomly generated truth table: ", y)
	fmt.Println("Zhegalkin polynomial coefficients:", zheg)
	fmt.Println("Reconstructed truth table:       ", reconstructed)
}
