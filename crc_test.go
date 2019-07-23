package main

import (
	"strings"
	"testing"
)

var largeText = []byte(strings.Repeat("a", 50000))
var smallText = []byte(strings.Repeat("a", 5))

func naive(data []byte, poly uint16, init uint16, xorout uint16) uint16 {
	var crc = init
	for _, item := range data {
		crc ^= uint16(item) << 0x8
		for j := 0; j < 8; j++ {
			if crc&0x8 == 0 {
				crc <<= 1
			} else {
				crc <<= 1
				crc ^= poly
			}
		}
	}
	return crc ^ xorout
}

func AssetEqual(t *testing.T, expected uint16, actual uint16) {
	if expected != actual {
		t.Errorf("Expected 0x%04X got 0x%04X", expected, actual)
	}
}

func TestCrc16EmptyVector(t *testing.T) {
	AssetEqual(t, 0xFFFF, Checksum(nil, 0x9AC1, 0xFFFF, 0x0000))
}

func TestNormalized(t *testing.T) {

	input := []byte("abcdefgh")

	t.Log("CRC-16/CCITT-FALSE")
	{
		AssetEqual(t, 0x9AC1, Checksum(input, 0x1021, 0xFFFF, 0x0000))
	}

	//t.Log("CRC-16/ARC")
	//{
	// CRC-16/ARC	0xE8C1	0xBB3D	0x8005	0x0000	true	true	0x0000
	//}

	t.Log("CRC-16/AUG-CCITT")
	{
		AssetEqual(t, 0x4AC6, Checksum(input, 0x1021, 0x1D0F, 0x0000))
	}

	t.Log("CRC-16/BUYPASS")
	{
		AssetEqual(t, 0x7D68, Checksum(input, 0x8005, 0x0000, 0x0000))
	}

	t.Log("CRC-16/CDMA2000")
	{
		AssetEqual(t, 0x2007, Checksum(input, 0xC867, 0xFFFF, 0x0000))
	}

	t.Log("CRC-16/DDS-110")
	{
		AssetEqual(t, 0x7388, Checksum(input, 0x8005, 0x800D, 0x0000))
	}

	t.Log("CRC-16/DECT-R")
	{
		AssetEqual(t, 0x73B6, Checksum(input, 0x0589, 0x0000, 0x0001))
	}

	t.Log("CRC-16/DECT-X")
	{
		AssetEqual(t, 0x73B7, Checksum(input, 0x0589, 0x0000, 0x0000))
	}

	//t.Log("CRC-16/DNP")
	//{
	// CRC-16/DNP	0xB350	0xEA82	0x3D65	0x0000	true	true	0xFFFF
	//}

	t.Log("CRC-16/EN-13757")
	{
		AssetEqual(t, 0x4A7D, Checksum(input, 0x3D65, 0x0000, 0xFFFF))
	}

	t.Log("CRC-16/GENIBUS")
	{
		AssetEqual(t, 0x653E, Checksum(input, 0x1021, 0xFFFF, 0xFFFF))
	}

	//t.Log("CRC-16/MAXIM")
	//{
	// CRC-16/MAXIM	0x173E	0x44C2	0x8005	0x0000	true	true	0xFFFF
	//}

	//t.Log("CRC-16/MCRF4XX")
	//{
	// CRC-16/MCRF4XX	0x7D08	0x6F91	0x1021	0xFFFF	true	true	0x0000
	//}

	//t.Log("CRC-16/RIELLO")
	//{
	// CRC-16/RIELLO	0xEB3B	0x63D0	0x1021	0xB2AA	true	true	0x0000
	//}

	t.Log("CRC-16/T10-DIF")
	{
		AssetEqual(t, 0xE9A3, Checksum(input, 0x8BB7, 0x0000, 0x0000))
	}

	t.Log("CRC-16/TELEDISK")
	{
		AssetEqual(t, 0x9F9D, Checksum(input, 0xA097, 0x0000, 0x0000))
	}

	//t.Log("CRC-16/TMS37157")
	//{
	// CRC-16/TMS37157	0xF7B8	0x26B1	0x1021	0x89EC	true	true	0x0000
	//}

	//t.Log("CRC-16/USB")
	//{
	// CRC-16/USB	0x5781	0xB4C8	0x8005	0xFFFF	true	true	0xFFFF
	//}

	//t.Log("CRC-A")
	//{
	// CRC-A	0x2371	0xBF05	0x1021	0xC6C6	true	true	0x0000
	//}

	//t.Log("CRC-16/KERMIT")
	//{
	// CRC-16/KERMIT	0x728F	0x2189	0x1021	0x0000	true	true	0x0000
	//}

	//t.Log("CRC-16/MODBUS")
	//{
	// CRC-16/MODBUS	0xA87E	0x4B37	0x8005	0xFFFF	true	true	0x0000
	//}

	//t.Log("CRC-16/X-25")
	//{
	// CRC-16/X-25	0x82F7	0x906E	0x1021	0xFFFF	true	true	0xFFFF
	//}

	t.Log("CRC-16/XMODEM")
	{
		AssetEqual(t, 0xABFF, Checksum(input, 0x1021, 0x0000, 0x0000))
	}

}

func BenchmarkNaiveSmall(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(smallText)))
	for n := 0; n < b.N; n++ {
		naive(smallText, 0x9AC1, 0xFFFF, 0x0000)
	}
}

func BenchmarkNaiveLarge(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(largeText)))
	for n := 0; n < b.N; n++ {
		naive(largeText, 0x9AC1, 0xFFFF, 0x0000)
	}
}

func BenchmarkCrcSmall(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(smallText)))
	for n := 0; n < b.N; n++ {
		Checksum(smallText, 0x9AC1, 0xFFFF, 0x0000)
	}
}

func BenchmarkCrcLarge(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(largeText)))
	for n := 0; n < b.N; n++ {
		Checksum(largeText, 0x9AC1, 0xFFFF, 0x0000)
	}
}
