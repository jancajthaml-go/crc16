## zero-alloc 16Bit Cyclic redundancy check

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/crc16)](https://goreportcard.com/report/jancajthaml-go/crc16)

CRC which encode messages by adding a fixed-length check value, for the purpose of error detection in communication networks, it can provide quick and reasonable assurance of the integrity of messages delivered.

However, it is not suitable for protection against intentional alteration of data.

Implementation provides both tableless and tabular checksum functions with variable 16bit polynomial.

### Supported standards ###

- CRC-16/CCITT-FALSE
- CRC-16/AUG-CCITT
- CRC-16/BUYPASS
- CRC-16/CDMA2000
- CRC-16/DDS-110
- CRC-16/DECT-R
- CRC-16/DECT-X
- CRC-16/EN-13757
- CRC-16/GENIBUS
- CRC-16/T10-DIF
- CRC-16/TELEDISK
- CRC-16/XMODEM

### Usage ###

```
import "github.com/jancajthaml-go/crc16"

data := []byte("abcdefgh")
poly := 0x1021
init := 0xFFFF
xorout := 0x0000

// for tableless
crc16.Checksum(data, poly, init, xorout) // 0x9AC1

// for precalculated tabular
instance = crc16.New(poly, init, xorout)
instance.Checksum(data) // 0x9AC1
```
