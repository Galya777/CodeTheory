package main

import "fmt"

func crc8(data []byte) byte {
	var crc byte = 0x00
	poly := byte(0x07) // x^8 + x^2 + x + 1

	for _, b := range data {
		crc ^= b
		for i := 0; i < 8; i++ {
			if (crc & 0x80) != 0 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}

func main() {
	data := []byte{0x12, 0x34, 0x56}
	crc := crc8(data)
	fmt.Printf("Данни: % x, CRC-8: %02x\n", data, crc)

	// Проверка
	packet := append(data, crc)
	fmt.Printf("Пакет с CRC: % x\n", packet)
}
