// for this task I chose to implement variant 2
package main

import (
	"fmt"
)

// Функция, която обръща даден вектор
func reverseVector(v []int) []int {
	n := len(v)
	// Създаваме нов вектор със същата дължина
	reversed := make([]int, n)
	for i := 0; i < n; i++ {
		reversed[i] = v[n-i-1]
	}
	return reversed
}

// Връща бинарното представяне на i с дължина m
func indexToBinary(i, m int) []int {
	b := make([]int, m)
	for j := 0; j < m; j++ {
		b[m-j-1] = (i >> j) & 1
	}
	return b
}

// XOR между два вектора
func xorVectors(a, b []int) []int {
	res := make([]int, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}

// Разделя вектор на две равни части
func splitY(y []int) ([]int, []int) {
	n := len(y)
	return y[:n/2], y[n/2:]
}

// Мажоритарна функция
func majorityBit(bits []int) int {
	sum := 0
	for _, b := range bits {
		sum += b
	}
	if sum*2 >= len(bits) {
		return 1
	}
	return 0
}

// Изчислява стойностите на квадратичния член (x1·x2)
func evaluateQuadraticTerm(a12 int) []int {
	res := make([]int, 4)
	for i := 0; i < 4; i++ {
		x := indexToBinary(i, 2)
		res[i] = a12 * x[0] * x[1]
	}
	return res
}

// Декодира RM(1,2): линейни + константа
func decodeRM1(y []int) []int {
	m := 2
	coeffs := make([]int, m+1)

	// Константа
	coeffs[2] = majorityBit(y)

	// Линейни
	for i := 0; i < m; i++ {
		count := 0
		for j := 0; j < len(y); j++ {
			x := indexToBinary(j, m)
			if x[i] == 1 {
				count += (1 - 2*y[j]) // 1 → -1, 0 → +1
			}
		}
		if count < 0 {
			coeffs[i] = 1
		}
	}
	return coeffs // [a1, a2, a0]
}

// Комбинира всички коефициенти: [a12, a1, a2, a0]
func combine(a12 int, linConst []int) []int {
	return append([]int{a12}, linConst...)
}

// Декодира RM(2,2)
func decodeRM2(y []int) []int {
	if len(y) != 4 {
		panic("decodeRM2: очакван вход с 4 елемента")
	}

	// Разделяне на входа на две части
	y0, y1 := splitY(y)

	// Използваме XOR за да получим разликата
	e := xorVectors(y0, y1)

	// Използваме мажоритарния бит за намиране на квадратичния член (a12)
	a12 := majorityBit(e)

	// Оценяваме квадратичния член
	quadTerm := evaluateQuadraticTerm(a12)

	// Премахваме квадратичния ефект от y
	corrected := xorVectors(y, quadTerm)

	return reverseVector(corrected)
}

// Кодира RM(2,2)
func encodeRM2(coeffs []int) []int {
	y := make([]int, 4)
	for i := 0; i < 4; i++ {
		x := indexToBinary(i, 2)
		a12, a1, a2, a0 := coeffs[0], coeffs[1], coeffs[2], coeffs[3]
		y[i] = a12*x[0]*x[1] ^ a1*x[0] ^ a2*x[1] ^ a0
	}
	return y
}

func main() {
	// Пример: f(x1,x2) = x1*x2 + x2
	coeffs := []int{1, 0, 1, 0}
	y := encodeRM2(coeffs)

	fmt.Println("Encoded y:  ", y)

	decoded := decodeRM2(y)
	fmt.Println("Decoded:    ", decoded)
	fmt.Println("Expected:   ", coeffs)
}
