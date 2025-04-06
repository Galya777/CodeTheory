package main

import "fmt"

// Функция за построяване на пораждаща матрица за удължен код на Голей
func generateGoleyExtendedCode() [][]int {
	// Матрица на Адамар (12x12)
	H := payley(12)
	ConvertHadamardToBinary(H)

	// Създаваме удължената матрица G на Голей (12x24)
	G := make([][]int, len(H)) // 12 реда
	for i := range G {
		G[i] = make([]int, len(H)*2) // 24 колони
	}

	// Копираме H в лявата част на G
	for i := 0; i < len(H); i++ {
		for j := 0; j < len(H); j++ {
			G[i][j] = H[i][j]
		}
	}

	// Добавяме единичната матрица I в дясната част на G
	for i := 0; i < len(H); i++ {
		// Индексите на единичната матрица са от 12 до 23
		G[i][i+len(H)] = 1
	}

	return G
}

// Функция за изчисляване на тегло на кодова дума
func weight(codeword []int) int {
	count := 0
	for _, bit := range codeword {
		if bit != 0 {
			count++
		}
	}
	return count
}
func generateGolayGeneratorMatrix() [][]int {
	// Стъпка 1: Създаване на идентична матрица I_12
	I_12 := make([][]int, 12)
	for i := 0; i < 12; i++ {
		I_12[i] = make([]int, 23)
		for j := 0; j < 12; j++ {
			if i == j {
				I_12[i][j] = 1
			} else {
				I_12[i][j] = 0
			}
		}
	}

	// Стъпка 2: Генериране на матрица на Адамар от ред 12
	HadamardMatrix := payley(12)

	// Стъпка 3: Комбинираме идентичната матрица с допълнителни колони от матрицата на Адамар
	// Добавяме колоните от матрицата на Адамар към идентичната матрица I_12
	for i := 0; i < 12; i++ {
		for j := 12; j < 23; j++ {
			I_12[i][j] = HadamardMatrix[i][j-12]
		}
	}

	return I_12
}

func proofGolley() {
	matrix := generateGoleyExtendedCode()
	var n, k, d int
	var codewords [][]int
	codewords = generateAllCodewords(matrix, len(matrix[0])/2)

	n = len(codewords[0])             // Дължина на кодовата дума
	k = len(codewords)                // Размерност на кода (брой информационни битове)
	d = minHammingDistance(codewords) // Минимално разстояние

	// Параметри на удължения код на Голей
	fmt.Printf("Параметри на удължения код на Голей: [n=%d, k=%d, d=%d]\n", n, k, d)

	// 3. Теглова функция
	// Изчисляваме теглото на всички кодови думи (редове на G)
	for i, codeword := range codewords {
		fmt.Printf("Тегло на кодова дума %d: %d\n", i+1, weight(codeword))
	}

	golayMatrix := generateGolayGeneratorMatrix()
	fmt.Println("Генераторна матрица на съвършения код на Голей:")
	for _, row := range golayMatrix {
		fmt.Println(row)
	}
	codewords = generateAllCodewords(golayMatrix, len(matrix[0])/2)
	for i, codeword := range codewords {
		fmt.Printf("Тегло на кодова дума %d: %d\n", i+1, weight(codeword))
	}

	// 4. Генериране на таблица на синдроми за декодиране
	// Създаване на проверочна матрица H за код на Голей
	H := GenerateParityCheckMatrix(matrix)

	// Генериране на таблица на синдроми
	syndromeTable := generateSyndromeDecodingTable(codewords, H)
	fmt.Println("Таблица на синдроми:")
	for _, syndrome := range syndromeTable {
		fmt.Println(syndrome)
	}

	// 5. Декодиране чрез синдроми (пример)
	// Примерен вектор, който искаме да декодираме
	noisyCodeword := []int{1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1}
	syndrome := calculateSyndrome(noisyCodeword, H)
	fmt.Println("Синдром на шумна кодова дума:", syndrome)
}
