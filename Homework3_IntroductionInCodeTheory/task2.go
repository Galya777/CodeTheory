package Homework3_IntroductionInCodeTheory

import "fmt"

// x^5 + x^3 + 1 = 100101 = 37
// p := 0b100101  // 37 на десетична
// Генерира всички полиноми от степен 5 над GF(2)
func generatePolynomials() []int {
	var polys []int
	for i := 0b100000; i <= 0b111111; i++ {
		polys = append(polys, i)
	}
	return polys
}
func divide(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return quotient, remainder
}

func degree(n int) int {
	deg := -1
	for n > 0 {
		n >>= 1
		deg++
	}
	return deg
}

// Проверява дали даден полином е неразложим
func isIrreducible(p int) bool {
	for i := 0b11; i < p; i++ { // почваме от степен 1 (0b11 = x + 1)
		if degree(i) >= degree(p) {
			break
		}
		_, r := divide(p, i)
		if r == 0 {
			return false // дели се на по-малък полином → не е неразложим
		}
	}
	return true
}

const fieldSize = 32
const primPoly = 0b100101 // x^5 + x^2 + 1

var expTable [62]int
var logTable [fieldSize]int

// Инициализира таблиците exp и log
func initTables() {
	expTable[0] = 1
	for i := 1; i < 62; i++ {
		expTable[i] = expTable[i-1] << 1
		if expTable[i-1]&0b10000 != 0 { // ако 5-тият бит е 1
			expTable[i] ^= primPoly
		}
	}
	for i := 0; i < 31; i++ {
		logTable[expTable[i]] = i
	}
}

// Събиране/изваждане в GF(32)
func addGF(a, b int) int {
	return a ^ b
}

// Умножение в GF(32)
func mulGF(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	expIndex := (logTable[a] + logTable[b]) % 31
	return expTable[expIndex]
}

// Деление в GF(32)
func divGF(a, b int) int {
	if b == 0 {
		panic("деление на 0 в GF(32)")
	}
	if a == 0 {
		return 0
	}
	expIndex := (logTable[a] - logTable[b] + 31) % 31
	return expTable[expIndex]
}

func printTables() {
	fmt.Println("expTable:")
	for i := 0; i < 31; i++ {
		fmt.Printf("exp[%2d] = %02b\n", i, expTable[i])
	}
	fmt.Println("\nlogTable:")
	for i := 1; i < fieldSize; i++ {
		fmt.Printf("log[%02b] = %2d\n", i, logTable[i])
	}
}

// Polynomial division over GF(2): returns remainder
func polyMod(a, b int) int {
	degB := degree(b)
	for degree(a) >= degB {
		shift := degree(a) - degB
		a ^= b << shift
	}
	return a
}
func findIrreduciblePolynomials(maxDegree int) []int {
	result := []int{}
	for i := 2; i < (1 << (maxDegree + 1)); i++ {
		if i&1 == 0 {
			continue // skip if not monic (constant term 1)
		}
		if isIrreducible(i) {
			result = append(result, i)
		}
	}
	return result
}

func formatPoly(p int) string {
	if p == 0 {
		return "0"
	}
	str := ""
	for i := 0; p > 0; i++ {
		if p&1 == 1 {
			if str != "" {
				str = " + " + str
			}
			if i == 0 {
				str = "1" + str
			} else if i == 1 {
				str = "x" + str
			} else {
				str = fmt.Sprintf("x^%d", i) + str
			}
		}
		p >>= 1
	}
	return str
}

func factorXnPlus1(n int) []int {
	xnPlus1 := (1 << (n + 1)) | 1 // x^n + 1 in binary
	factors := []int{}
	irreducibles := findIrreduciblePolynomials(n)
	for _, p := range irreducibles {
		if polyMod(xnPlus1, p) == 0 {
			factors = append(factors, p)
		}
	}
	return factors
}
func lcm(a, b int) int {
	product := a * b
	for b != 0 {
		a, b = b, a%b
	}
	return product / a
}

func getGenerator(t int) int {
	roots := []int{}
	for i := 1; i <= 2*t; i++ {
		roots = append(roots, i)
	}
	// Най-простата версия: просто взимаме минимални полиноми и ги умножаваме
	gen := 1
	for _, r := range roots {
		minPoly := 1 << r // прост модел за минимален полином: x^r
		gen = lcm(gen, minPoly)
	}
	return gen
}

func printPoly(p int) {
	fmt.Println(formatPoly(p))
}
func main() {
	//a
	polys := generatePolynomials()
	for _, p := range polys {
		if isIrreducible(p) {
			fmt.Printf("Irreducible: %06b\n", p)
		}
	}

	//б
	initTables()
	printTables()

	a := 0b00101 // 5
	b := 0b00011 // 3

	fmt.Printf("\nПримери със стойности a = %d и b = %d\n", a, b)
	fmt.Printf("Събиране: %d + %d = %d\n", a, b, addGF(a, b))
	fmt.Printf("Умножение: %d * %d = %d\n", a, b, mulGF(a, b))
	fmt.Printf("Деление: %d / %d = %d\n", a, b, divGF(a, b))

	//в
	n := 31
	fmt.Printf("Factoring x^%d + 1 over GF(2):\n", n)
	factors := factorXnPlus1(n)
	for _, f := range factors {
		fmt.Println(formatPoly(f))
	}

	//г
	for t := 1; t <= 3; t++ {
		fmt.Printf("BCH код, коригиращ до %d грешки:\n", t)
		g := getGenerator(t)
		fmt.Print("g(x) = ")
		printPoly(g)
		fmt.Println()
	}
}
