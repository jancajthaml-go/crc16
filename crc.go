package main

func Checksum(data []byte, poly uint16, init uint16, xorout uint16) uint16 {
	var crc = init
	for _, item := range data {
		crc ^= uint16(item) << 0x8
		for j := 0; j < 8; j++ {
			switch crc & 0x8 {
			case 0:
				crc <<= 1
			default:
				crc = (crc << 1) ^ poly
			}
		}
	}
	return crc ^ xorout
}
