package main

import (
	"bufio"
	"fmt"
	_ "io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadMatrixFromConsole() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	var matrix [][]int

	fmt.Println("Въведи редовете на матрицата. Празен ред = край:")

	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}

		fields := strings.Fields(line)
		var row []int
		for _, f := range fields {
			num, err := strconv.Atoi(f)
			if err != nil || (num != 0 && num != 1) {
				fmt.Println("Невалидна стойност – само 0 или 1!")
				continue
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	return matrix
}
func ReadMatrixFromFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var row []int
		for _, f := range fields {
			num, err := strconv.Atoi(f)
			if err != nil || (num != 0 && num != 1) {
				return nil, fmt.Errorf("невалидна стойност: %s", f)
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

// Функция за генериране на проверочна матрица H от пораждаща матрица G
func GenerateParityCheckMatrix(G [][]int) [][]int {
	k := len(G)    // брой редове на G (размерността на кода)
	n := len(G[0]) // брой колони в G (дължината на кодовата дума)

	// Инициализиране на матрица H с размери (n-k) x n
	H := make([][]int, n-k)
	for i := range H {
		H[i] = make([]int, n)
	}

	// Запълваме лявата част на H с част от G (всички колони от G, започвайки от ред k)
	for i := 0; i < n-k; i++ {
		// Проверка дали индексите са в правилните граници
		if k+i < len(G) {
			for j := 0; j < k; j++ {
				H[i][j] = G[k+i][j]
			}
		}
	}

	// Добавяме единичната матрица в дясната част на H
	for i := 0; i < n-k; i++ {
		H[i][k+i] = 1
	}

	return H
}

// Функция за генериране на пораждаща матрица G от проверочна матрица H
func GenerateGeneratorMatrix(H [][]int) [][]int {
	r := len(H[0]) // Дължина на кодовата дума
	n := len(H)    // Брой проверяващи уравнения

	k := n - r // Размерност на кода

	// Инициализиране на G с размер k x n
	G := make([][]int, k)
	for i := range G {
		G[i] = make([]int, n)
	}

	// Попълване на I_k (единичната матрица)
	for i := 0; i < k; i++ {
		G[i][i] = 1
	}

	// Попълване на P частта, ако H е в стандартна форма [P^T | I_r]
	for i := 0; i < k; i++ {
		for j := 0; j < r-k; j++ {
			G[i][k+j] = H[j][i] // Взимаме P^T и го "транспонираме"
		}
	}

	return G
}

// извеждане на кода за пораждаща матрица
func multiplyVectorByMatrixMod2(vector []int, matrix [][]int) []int {
	k := len(vector)
	if k == 0 || k != len(matrix) {
		panic("Sizes must be equal length!")
	}
	n := len(matrix[0])
	result := make([]int, n)

	// Обходи j от 0 до n-1 (всяка колона)
	for j := 0; j < n; j++ {
		sum := 0
		for i := 0; i < k; i++ {
			sum += matrix[i][j] * vector[i]
		}
		result[j] = sum % 2
	}

	return result
}

// Радиус на покритие
func coverageRadius(d int) int {
	return (d - 1) / 2
}

// Броят на поправените грешки
func errorCorrectionCapability(d int) int {
	return (d - 1) / 2
}

// Броят на откритите грешки
func errorDetectionCapability(d int) int {
	return d - 1
}

func generateAllCodewords(G [][]int, total int) [][]int {
	k := len(G) // брой редове на G (дължина на входния вектор)
	codewords := make([][]int, 0, total)
	for i := 0; i < total; i++ {
		u := make([]int, k)
		for j := 0; j < k; j++ {
			if (i>>j)&1 == 1 {
				u[j] = 1
			}
		}
		codeword := multiplyVectorByMatrixMod2(u, G)
		codewords = append(codewords, codeword)
	}
	return codewords
}

// Функция за транспониране на матрица
func TransposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	transposed := make([][]int, cols)

	for i := 0; i < cols; i++ {
		transposed[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			transposed[i][j] = matrix[j][i]
		}
	}

	return transposed
}

// Намиране на дуалната проверочна матрица
func FindDualParityCheckMatrix(H [][]int) [][]int {
	return TransposeMatrix(H)
}

// Намиране на дуалната пораждаща матрица
func FindDualGeneratorMatrix(H [][]int) [][]int {
	dualH := FindDualParityCheckMatrix(H)
	return GenerateGeneratorMatrix(dualH)
}

// Функция за кодиране на информационен вектор чрез пораждаща матрица G
func encodeInformationVector(informationVector []int, G [][]int) []int {
	// Множим информационния вектор с пораждащата матрица G по модул 2
	return multiplyVectorByMatrixMod2(informationVector, G)
}

// Алгоритъм за кодиране (точка 4)
func encodingAlgorithm(G [][]int) []int {
	// Въвеждаме информационния вектор
	var input string
	fmt.Println("Въведи информационния вектор:")
	fmt.Scan(&input)

	// Преобразуваме входа в числов вектор
	informationVector := make([]int, len(input))
	for i, char := range input {
		if char == '1' {
			informationVector[i] = 1
		} else if char == '0' {
			informationVector[i] = 0
		} else {
			panic("Невалиден вход, само 0 и 1 са разрешени.")
		}
	}

	// Проверка на дължината на информационния вектор
	if len(informationVector) != len(G) {
		panic("Грешка: Дължината на информационния вектор не съответства на дължината на кодовата дума.")
	}

	// Извършваме кодиране
	encodedVector := encodeInformationVector(informationVector, G)

	return encodedVector
}

// Функция за имитация на канал с шум
func simulateChannelNoise(encodedVector []int, numErrors int) []int {
	// Създаваме копие на кодовата дума, за да не я променяме директно
	noisyVector := make([]int, len(encodedVector))
	copy(noisyVector, encodedVector)

	// Инициализиране на генератора за случайни числа
	rand.Seed(time.Now().UnixNano())

	// Добавяме грешки на случайни позиции
	for i := 0; i < numErrors; i++ {
		// Избираме случайна позиция в вектора
		position := rand.Intn(len(noisyVector))
		// Променяме стойността на позицията (от 0 на 1 или от 1 на 0)
		noisyVector[position] = 1 - noisyVector[position]
	}

	// Връщаме новата кодова дума с грешки
	return noisyVector
}

// Функция за добавяне на шум (случайни грешки) към кодова дума
func addNoiseToCodeword(codeword []int, numErrors int) []int {
	n := len(codeword)
	// Създаваме копие на кодовата дума
	noisyCodeword := make([]int, n)
	copy(noisyCodeword, codeword)

	// Добавяме случайни грешки (инвертиране на случайни битове)
	for i := 0; i < numErrors; i++ {
		pos := rand.Intn(n)
		noisyCodeword[pos] = 1 - noisyCodeword[pos] // Инвертиране на бит
	}

	return noisyCodeword
}

// Функция за генериране на шумни кодови думи
func generateNoisyCodewords(codewords [][]int, numErrors int) [][]int {
	noisyCodewords := make([][]int, len(codewords))

	// За всяка кодова дума добавяме шум (грешки)
	for i, codeword := range codewords {
		noisyCodewords[i] = addNoiseToCodeword(codeword, numErrors)
	}

	return noisyCodewords
}

// Функция за изчисляване на синдром
func calculateSyndrome(codeword []int, H [][]int) []int {
	syndrome := make([]int, len(H))
	for i := 0; i < len(H); i++ {
		sum := 0
		for j := 0; j < len(codeword); j++ {
			sum += codeword[j] * H[i][j]
		}
		syndrome[i] = sum % 2 // Извършваме операция по модул 2
	}
	return syndrome
}

// Функция за генериране на таблица на Слепян
func generateSlepianTable(codewords [][]int, H [][]int) map[string][]int {
	table := make(map[string][]int)

	// За всяка кодова дума изчисляваме синдрома и го добавяме в таблицата
	for _, codeword := range codewords {
		syndrome := calculateSyndrome(codeword, H)
		syndromeStr := fmt.Sprint(syndrome)

		// Добавяме кодовата дума в таблицата по съответния синдром
		table[syndromeStr] = codeword
	}

	return table
}

// Функция за запис на таблицата на Слепян във файл
func writeSlepianTableToFile(table map[string][]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записваме всяка двойка синдром-кодова дума в таблицата
	for syndrome, codeword := range table {
		file.WriteString(fmt.Sprintf("Синдром: %s -> Кодова дума: ", syndrome))
		for _, value := range codeword {
			file.WriteString(fmt.Sprintf("%d ", value))
		}
		file.WriteString("\n")
	}

	return nil
}

// Функция за генериране на таблица за декодиране чрез синдроми
func generateSyndromeDecodingTable(codewords [][]int, H [][]int) map[string][]int {
	decodingTable := make(map[string][]int)

	// За всяка кодова дума изчисляваме синдрома и добавяме в таблицата
	for _, codeword := range codewords {
		syndrome := calculateSyndrome(codeword, H)
		syndromeStr := fmt.Sprint(syndrome) // Преобразуваме синдрома в низ (за ключ)

		// Добавяме кодовата дума в таблицата по съответния синдром
		decodingTable[syndromeStr] = codeword
	}

	return decodingTable
}

// Функция за декодиране чрез синдроми
func decodeBySyndrome(noisyCodeword []int, H [][]int, decodingTable map[string][]int) []int {
	// Изчисляваме синдрома на шумната кодова дума
	syndrome := calculateSyndrome(noisyCodeword, H)
	syndromeStr := fmt.Sprint(syndrome) // Преобразуваме синдрома в низ (за ключ)

	// Търсим в таблицата за декодиране съответстващата кодова дума
	decodedCodeword, exists := decodingTable[syndromeStr]
	if !exists {
		fmt.Println("Грешка: Не е намерен съответстващ код за този синдром.")
		return nil
	}

	return decodedCodeword
}

// Функция за запис на таблицата за декодиране в файл
func writeSyndromeDecodingTableToFile(table map[string][]int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записваме всяка двойка синдром-кодова дума в таблицата
	for syndrome, codeword := range table {
		file.WriteString(fmt.Sprintf("Синдром: %s -> Кодова дума: ", syndrome))
		for _, value := range codeword {
			file.WriteString(fmt.Sprintf("%d ", value))
		}
		file.WriteString("\n")
	}

	return nil
}
func writeTofile(matrix [][]int, file *os.File) error {
	for _, row := range matrix {
		for _, value := range row {
			file.WriteString(fmt.Sprintf("%d ", value))
		}
		file.WriteString("\n")
	}
	return nil
}
func readType(choice int) string {

	// Записваме типа на матрицата
	var matrixType string
	if choice == 1 {
		fmt.Scan(&matrixType)
	} else if choice == 2 {
		var filename string
		fmt.Print("Име на файла: ")
		fmt.Scan(&filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(strings.TrimSpace(string(data)), "\n")
		matrixType = strings.TrimSpace(lines[0])
	} else {
		panic("Невалиден избор.")
	}
	return matrixType
}

func readInputMatrix(choice int) [][]int {

	var matrix [][]int
	var err error

	// Четене на матрица от конзолата или файл
	if choice == 1 {
		matrix = ReadMatrixFromConsole()
	} else if choice == 2 {
		var filename string
		fmt.Print("Име на файла: ")
		fmt.Scan(&filename)
		matrix, err = ReadMatrixFromFile(filename)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Невалиден избор.")
	}

	return matrix
}
func errosAndRadius(matrix [][]int, file *os.File) {
	var n, k, d int
	var codewords [][]int
	codewords = generateAllCodewords(matrix, len(matrix[0])/2)

	n = len(codewords[0])             // Дължина на кодовата дума
	k = len(codewords)                // Размерност на кода (брой информационни битове)
	d = minHammingDistance(codewords) // Минимално разстояние

	R := coverageRadius(d)                      // Радиус на покритие
	correctable := errorCorrectionCapability(d) // Поправими грешки
	detectable := errorDetectionCapability(d)   // Откриваеми грешки

	// Записване във файл
	file.WriteString(fmt.Sprintf("\nДължина на кодовата дума (n): %d\n", n))
	file.WriteString(fmt.Sprintf("Размерност на кода (k): %d\n", k))
	file.WriteString(fmt.Sprintf("Минимално разстояние (d): %d\n", d))
	file.WriteString(fmt.Sprintf("Радиус на покритие (R): %d\n", R))
	file.WriteString(fmt.Sprintf("Брой на поправените грешки: %d\n", correctable))
	file.WriteString(fmt.Sprintf("Брой на откритите грешки: %d\n", detectable))

}
func linearCodes() {
	// Избор на вход (конзола или файл)
	var choice int
	fmt.Println("Избери вход (1 = конзола, 2 = файл):")
	fmt.Scan(&choice)
	matrixType := readType(choice)

	// Записваме въведената матрица и резултатите в един файл
	file, err := os.Create("linear_code_results.txt")
	if err != nil {
		fmt.Println("Грешка при създаване на файла:", err)
		return
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("\nТип матрица: %s\n", matrixType))

	//Точка 1
	var G [][]int
	var H [][]int
	if matrixType == "G" || matrixType == "g" {
		// Генерираме проверочната матрица H от G
		G = readInputMatrix(choice)
		H = GenerateParityCheckMatrix(G)
		printMatrix(H)
	} else if matrixType == "H" || matrixType == "h" {
		// Генерираме пораждаща матрица G от H
		H = readInputMatrix(choice)
		G = GenerateGeneratorMatrix(H)
		printMatrix(G)
	} else {
		panic("Невалиден тип матрица.")
	}

	file.WriteString("\nПораждащата матрица G:\n")
	writeTofile(G, file)
	file.WriteString("\nПроверочната матрица H:\n")
	writeTofile(H, file)

	//Точка 2
	// Генерираме дуалната проверочна и пораждаща матрица
	dualH := FindDualParityCheckMatrix(H)
	file.WriteString("\nДуалната проверочна матрица H:\n")
	writeTofile(dualH, file)

	dualG := FindDualGeneratorMatrix(G)
	file.WriteString("\nДуалната пораждаща матрица G:\n")
	writeTofile(dualG, file)

	//Точка 3
	errosAndRadius(G, file)
	errosAndRadius(dualG, file)
	errosAndRadius(dualG, file)
	errosAndRadius(dualH, file)

	//Точка 4 Алгоритъм за кодиране
	encodedVector := encodingAlgorithm(G)
	// Записваме кодовата дума в файла
	file.WriteString("\nКодова дума:\n")
	for _, value := range encodedVector {
		file.WriteString(fmt.Sprintf("%d ", value))
	}
	file.WriteString("\n")

	// Точка 5
	// Въведете брой грешки за имитация на шум
	var numErrors int
	fmt.Println("Въведи брой грешки за имитация на шум:")
	fmt.Scan(&numErrors)

	// Извикваме функцията за имитация на шум
	noisyEncodedVector := simulateChannelNoise(encodedVector, numErrors)

	// Записваме кодовата дума със шум във файла
	file.WriteString("\nКодова дума със шум:\n")
	for _, value := range noisyEncodedVector {
		file.WriteString(fmt.Sprintf("%d ", value))
	}
	file.WriteString("\n")

	var codewords [][]int
	codewords = generateAllCodewords(G, len(G[0])/2)

	//Точка 6
	// Генерираме таблица на Слепян
	// Генерираме шумни кодови думи
	noisyCodewords := generateNoisyCodewords(codewords, numErrors)

	// Генерираме таблица на Слепян за шумни кодови думи
	table := generateSlepianTable(noisyCodewords, G)

	// Записваме таблицата на Слепян в текстов файл
	err = writeSlepianTableToFile(table, "noisy_slepian_table.txt")
	if err != nil {
		fmt.Println("Грешка при запис на таблицата на Слепян:", err)
		return
	}
	fmt.Println("Таблицата на Слепян за шумни кодови думи беше записана успешно във файла 'noisy_slepian_table.txt'.")

	//Тoчка 7
	//Генерираме таблица за декодиране чрез синдроми
	decodingTable := generateSyndromeDecodingTable(codewords, G)

	// Записваме таблицата за декодиране в текстов файл
	err = writeSyndromeDecodingTableToFile(decodingTable, "syndrome_decoding_table.txt")
	if err != nil {
		fmt.Println("Грешка при запис на таблицата за декодиране:", err)
		return
	}

	fmt.Println("Таблицата за декодиране беше записана успешно във файла 'syndrome_decoding_table.txt'.")

	// Декодиране на шумни кодови думи чрез синдроми
	noisyCodeword := noisyEncodedVector
	decodedCodeword := decodeBySyndrome(noisyCodeword, G, decodingTable)

	if decodedCodeword != nil {
		fmt.Println("Декодирана кодова дума:", decodedCodeword)
	}

}
