package crc16

// CRC holds crc16 parameters and precalculated table
type CRC struct {
	table  []uint16
	poly   uint16
	xorout uint16
	init   uint16
}

// New returns crc16 instance with precalculated table
func NewCRC(poly uint16, init uint16, xorout uint16) CRC {
	return CRC{
		poly:   poly,
		xorout: xorout,
		init:   init,
		table:  createTable(poly, init, xorout),
	}
}

func createTable(poly uint16, init uint16, xorout uint16) []uint16 {
	result := make([]uint16, 256)
	var bit uint16
	for divident := 0; divident < 256; divident++ {
		var current uint16 = 0x0000
		for j := uint16(0x0080); j != 0; j >>= 1 {
			if (uint16(divident) & j) != 0 {
				bit = (current & 0x8000) ^ 0x8000
			} else {
				bit = current & 0x8000
			}
			switch bit {
			case 0:
				current <<= 1
			default:
				current = (current << 1) ^ poly
			}
		}
		result[divident] = current & 0xFFFF
	}
	return result
}

// Checksum returns crc16 checksum of given CRC instance
func (crc *CRC) Checksum(data []byte) uint16 {
	var pos uint8
	var result = crc.init
	for _, item := range data {
		result ^= uint16(item) << 8
		pos = uint8((result >> 8) & 0xFF)
		result <<= 8
		result ^= crc.table[pos]
	}
	return result ^ crc.xorout
}

// Checksum returns crc16 checksum for given parameters
func Checksum(data []byte, poly uint16, init uint16, xorout uint16) uint16 {
	var (
		crc uint16 = init
		bit uint16
	)

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
