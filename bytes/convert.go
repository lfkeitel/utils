package bytes

import "encoding/binary"

// Itob converts an int64 into an 8 byte slice
func Itob(i int64) []byte {
	out := make([]byte, 8)
	binary.PutVarint(out, i)
	return out
}

// Btoi converts a byte slice to an int64 ignoring errors.
func Btoi(in []byte) int64 {
	out, _ := binary.Varint(in)
	return out
}
