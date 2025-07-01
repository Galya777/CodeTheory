package main

import "fmt"

func checkCRC(dataWithCRC uint8) bool {
	return crc4(dataWithCRC) == 0
}

func main() {
	var data uint8 = 0b1101
	crc := crc4(data)
	packet := (data << 4) | crc

	fmt.Printf("Пакет: %08b\n", packet)
	fmt.Println("Проверка на пакета:", checkCRC(packet))

	// Случай с грешка
	packetWithError := packet ^ 0b00010000 // променяме един бит
	fmt.Printf("Пакет с грешка: %08b\n", packetWithError)
	fmt.Println("Проверка на грешния пакет:", checkCRC(packetWithError))
}
