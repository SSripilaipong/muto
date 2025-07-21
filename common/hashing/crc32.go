package hashing

import (
	"encoding/binary"
	"hash/crc32"
	"math"

	"github.com/SSripilaipong/go-common/optional"
)

func CRC32(value any) optional.Of[uint32] {
	table := crc32.MakeTable(crc32.IEEE)

	switch v := value.(type) {
	case int:
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(v))
		return optional.Value(crc32.Checksum(b, table))
	case float64:
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, math.Float64bits(v))
		return optional.Value(crc32.Checksum(b, table))
	case bool:
		if v {
			return optional.Value(crc32.Checksum([]byte{1}, table))
		} else {
			return optional.Value(crc32.Checksum([]byte{0}, table))
		}
	case string:
		return optional.Value(crc32.Checksum([]byte(v), table))
	default:
		return optional.Empty[uint32]()
	}
}
