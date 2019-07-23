package crc16

func Checksum(data []byte, poly uint16, init uint16, xorout uint16) uint16 {
	var crc uint16 = init
	var bit uint16

	for _, item := range data {
		for j := uint16(0x0080); j != 0; j >>= 1 {
			if (uint16(item) & j) != 0 {
				bit = (crc & 0x8000) ^ 0x8000
			} else {
				bit = crc & 0x8000
			}
			switch bit {
			case 0:
				crc <<= 1
			default:
				crc = (crc << 1) ^ poly
			}
		}
	}
	return crc ^ xorout
}
