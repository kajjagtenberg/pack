package pack

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
)

var (
	SEPERATOR = []byte(":")
)

// Encodes the values to base32 to be able to safely seperated them by the SEPERATOR, which is a colon.
// It is not optimal use of the available bytes, but often key safety is more important than the size.
func Pack(values ...[]byte) []byte {
	for i := 0; i < len(values); i++ {
		b := make([]byte, base64.StdEncoding.EncodedLen(len(values[i])))
		base64.StdEncoding.Encode(b, values[i])
		values[i] = b
	}

	return bytes.Join(values, SEPERATOR)
}

// Returns the decoded values from a packed byte array. Simply the inverse of the Pack() function.
func Unpack(value []byte) [][]byte {
	values := bytes.Split(value, SEPERATOR)

	for i := 0; i < len(values); i++ {
		b := make([]byte, base64.StdEncoding.DecodedLen(len(values[i])))
		base64.StdEncoding.Decode(b, values[i])
		values[i] = b
	}

	return values
}

// Quick way to convert an uint64 to bytes for packing.
func Uint64ToBytes(value uint64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, value)
	return result
}

// Quick way to convert bytes to an uint64 after unpacking.
func BytesToUint64(value []byte) uint64 {
	return binary.BigEndian.Uint64(value)
}
