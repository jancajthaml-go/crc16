## 16Bit Cyclic redundancy check

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/crc16)](https://goreportcard.com/report/jancajthaml-go/crc16)

CRC which encode messages by adding a fixed-length check value, for the purpose of error detection in communication networks, it can provide quick and reasonable assurance of the integrity of messages delivered.

However, it is not suitable for protection against intentional alteration of data.

Implementation is tableless with variable 16bit polynomial.


### Performance ###

```
BenchmarkCrcSmall    100000000    54.2 ns/op      92.30 MB/s    0 B/op    0 allocs/op
BenchmarkCrcLarge    10000        506271 ns/op    98.76 MB/s    0 B/op    0 allocs/op
```

### Usage ###

```
import "github.com/jancajthaml-go/crc16"

crc16.Checksum(input, 0x1021, 0xFFFF, 0x0000) // 0x9AC1
```
