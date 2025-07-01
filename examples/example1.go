package main

import "fmt"

// crc4 изчислява CRC-4 на даден байтов с генераторен полином 0x3 (x^4 + x + 1)
func crc4(data uint8) uint8 {
	poly := uint8(0x3) // генераторен полином
	data <<= 4         // оставяме място за CRC 4 бита

	for i := 7; i >= 4; i-- {
		if (data>>i)&1 == 1 {
			data ^= poly << (i - 4)
		}
	}
	return data & 0xF // остатък (CRC)
}

func main() {
	var data uint8 = 0b1101 // примерни данни (4 бита)
	crc := crc4(data)
	fmt.Printf("Данни: %04b, CRC-4: %04b\n", data, crc)
}
