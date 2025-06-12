package main

import (
	"fmt"
	"strings"
)

// Преобразува стринг като "100110000" в []int
func parseBitString(s string) []int {
	bits := make([]int, len(s))
	for i, ch := range s {
		if ch == '1' {
			bits[i] = 1
		}
	}
	return bits
}

// Премахва водещи нули
func trimLeadingZeros(p []int) []int {
	i := 0
	for i < len(p) && p[i] == 0 {
		i++
	}
	return p[i:]
}

// Изваждане на полиноми по модул 2 (XOR)
func xorPolynomials(a, b []int) []int {
	a = trimLeadingZeros(a)
	b = trimLeadingZeros(b)

	// Допълваме по дължина
	if len(b) > len(a) {
		a, b = b, a
	}
	diff := len(a) - len(b)
	res := make([]int, len(a))
	copy(res, a)
	for i := 0; i < len(b); i++ {
		res[i+diff] ^= b[i]
	}
	return trimLeadingZeros(res)
}

// Деление по модул 2 – връща остатъка
func modPoly(dividend, divisor []int) []int {
	dividend = trimLeadingZeros(dividend)
	divisor = trimLeadingZeros(divisor)

	for len(dividend) >= len(divisor) {
		shift := len(dividend) - len(divisor)
		shiftedDivisor := make([]int, len(dividend))
		for i := 0; i < len(divisor); i++ {
			shiftedDivisor[i+shift] = divisor[i]
		}
		dividend = xorPolynomials(dividend, shiftedDivisor)
	}
	return trimLeadingZeros(dividend)
}

// Намира GCD на два полинома по модул 2
func gcdGF2(a, b []int) []int {
	for len(b) > 0 {
		r := modPoly(a, b)
		a, b = b, r
	}
	return trimLeadingZeros(a)
}

// Генерира x^n - 1 = [1, 0, ..., 0, 1]
func generateXNMinus1(n int) []int {
	p := make([]int, n+1)
	p[0] = 1
	p[n] = 1
	return p
}

// Преобразува []int в полиномиален формат
func polyToString(p []int) string {
	p = trimLeadingZeros(p)
	degree := len(p) - 1
	var terms []string
	for i, coef := range p {
		if coef == 1 {
			exp := degree - i
			switch exp {
			case 0:
				terms = append(terms, "1")
			case 1:
				terms = append(terms, "x")
			default:
				terms = append(terms, fmt.Sprintf("x^%d", exp))
			}
		}
	}
	return strings.Join(terms, " + ")
}

func main() {

	aStr := "100110000"
	aBits := parseBitString(aStr)

	n := len(aBits)
	xn1 := generateXNMinus1(n)

	g := gcdGF2(aBits, xn1)
	k := n - (len(g) - 1)

	fmt.Println("Кодов вектор a(x):    ", aBits)
	fmt.Println("x^n - 1:              ", xn1)
	fmt.Println("Пораждащ полином g(x):", g)
	fmt.Println("g(x) =                ", polyToString(g))
	fmt.Printf("Параметри на кода:    [n=%d, k=%d]\n", n, k)
}
