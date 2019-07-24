## 16Bit Cyclic redundancy check

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/crc16)](https://goreportcard.com/report/jancajthaml-go/crc16)

CRC which encode messages by adding a fixed-length check value, for the purpose of error detection in communication networks, it can provide quick and reasonable assurance of the integrity of messages delivered.

However, it is not suitable for protection against intentional alteration of data.

Implementation provides both tableless and tabular checksum functions with variable 16bit polynomial.

### Performance ###

```
BenchmarkCrcSmall               63.42 MB/s   0 B/op  0 allocs/op
BenchmarkCrcLarge               28.34 MB/s   0 B/op  0 allocs/op
BenchmarkPrecalculatedCrcSmall  536.36 MB/s  0 B/op  0 allocs/op
BenchmarkPrecalculatedCrcLarge  384.65 MB/s  0 B/op  0 allocs/op
```

### Usage ###

```
import "github.com/jancajthaml-go/crc16"

data := []byte("abcdefgh")
poly := 0x1021
init := 0xFFFF
xorout := 0x0000

// for tableless
crc16.Checksum(data, poly, init, xorout) // 0x9AC1

// for tabular
instance = crc16.New(poly, init, xorout)
instance.Checksum(data) // 0x9AC1
```
