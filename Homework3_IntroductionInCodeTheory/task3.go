package Homework3_IntroductionInCodeTheory

import (
	"fmt"
)

// GF(7) – работа в полето
const q = 7
const n = q - 1 // тъй като GF(7), имаме RS код с n = 6

// Полиноми – представени като slice от коефициенти [c0, c1, c2, ..., cn] (c0 + c1*x + c2*x^2 + ...)

// Извършва умножение на два полинома в GF(q)
func polyMul(a, b []int) []int {
	res := make([]int, len(a)+len(b)-1)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			res[i+j] = (res[i+j] + a[i]*b[j]) % q
		}
	}
	return res
}

// Генерира пораждащ полином за RS код с минимално разстояние d
func generateGeneratorPoly(d int) []int {
	// Генерираме (x - α^1)(x - α^2)...(x - α^(d-1)) в GF(7)
	g := []int{1} // g(x) = 1

	for i := 1; i <= d-1; i++ {
		alphaPow := powMod(3, i, q)             // α = 3 е примитивен елемент в GF(7)
		factor := []int{(-alphaPow + q) % q, 1} // (x - α^i) = -α^i + x
		g = polyMul(g, factor)
	}
	return g
}

// Извежда полином като четим string
func printPolly(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != 0 {
			if i == 0 {
				fmt.Printf("%d", p[i])
			} else if i == 1 {
				fmt.Printf("%dx + ", p[i])
			} else {
				fmt.Printf("%dx^%d + ", p[i], i)
			}
		}
	}
	fmt.Println()
}

// Степенуване по модул q
func powMod(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		b /= 2
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Основна програма
func main() {
	fmt.Println("=== BCH/RS код над GF(7) ===")
	fmt.Printf("Полето GF(%d), дължина на кода n = %d\n", q, n)

	// избираме t = 1 (до 1 грешка), тогава d = 2t + 1 = 3
	d := 3
	k := n - (d - 1) // параметър k (броят на информационните символи)

	// Генерираме пораждащия полином
	fmt.Println("\nГенериране на пораждащ полином:")
	g := generateGeneratorPoly(d)
	fmt.Print("g(x) = ")
	printPolly(g)

	// Параметри на кода
	fmt.Println("\nПараметри на кода:")
	fmt.Printf("n = %d, k = %d, d ≥ %d\n", n, k, d)

	// Дуален код
	fmt.Println("\nПараметри на дуалния код:")
	kDual := n - k
	fmt.Printf("n = %d, k = %d\n", n, kDual)
}
