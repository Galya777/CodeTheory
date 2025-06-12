package main

import (
	"fmt"
)

// Генерира всички булеви вектори с дължина m
func generateBooleanVectors(m int) [][]int {
	n := 1 << m
	vectors := make([][]int, n)
	for i := 0; i < n; i++ {
		vectors[i] = make([]int, m)
		for j := 0; j < m; j++ {
			vectors[i][j] = (i >> (m - j - 1)) & 1
		}
	}
	return vectors
}

// Връща всички комбинации от индекси за мономите със степен ≤ r
func generateMonomialIndices(m, r int) [][]int {
	var helper func(start, depth int, current []int)
	var result [][]int

	helper = func(start, depth int, current []int) {
		if len(current) == depth {
			temp := make([]int, depth)
			copy(temp, current)
			result = append(result, temp)
			return
		}
		for i := start; i < m; i++ {
			helper(i+1, depth, append(current, i))
		}
	}

	for deg := 0; deg <= r; deg++ {
		helper(0, deg, []int{})
	}

	return result
}

// Изчислява стойността на моном върху булев вектор
func evaluateMonomial(vector []int, indices []int) int {
	val := 1
	for _, idx := range indices {
		val *= vector[idx]
	}
	return val
}

// Генерира пораждащата матрица G
func generateRMGeneratorMatrix(r, m int) [][]int {
	n := 1 << m
	monomials := generateMonomialIndices(m, r)
	vectors := generateBooleanVectors(m)

	G := make([][]int, len(monomials))
	for i, mono := range monomials {
		G[i] = make([]int, n)
		for j, v := range vectors {
			G[i][j] = evaluateMonomial(v, mono)
		}
	}
	return G
}

// Несистематично кодиране: code = u * G
func encode(u []int, G [][]int) []int {
	n := len(G[0])
	k := len(G)
	if len(u) != k {
		panic("Length of u must match number of rows in G")
	}

	codeword := make([]int, n)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			sum ^= u[j] * G[j][i]
		}
		codeword[i] = sum
	}
	return codeword
}

func main() {
	r, m := 2, 5 // Пример: RM(1,3)
	G := generateRMGeneratorMatrix(r, m)

	fmt.Println("Пораждаща матрица G:")
	for _, row := range G {
		fmt.Println(row)
	}

	u := []int{1, 0, 1, 1} // примерен входен вектор
	code := encode(u, G)
	fmt.Println("Кодирана дума:")
	fmt.Println(code)
}
