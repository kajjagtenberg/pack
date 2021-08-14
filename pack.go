package pack

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
)

var (
	SEPERATOR = []byte(":")
)

func Pack(values ...[]byte) []byte {
	for i := 0; i < len(values); i++ {
		b := make([]byte, base64.StdEncoding.EncodedLen(len(values[i])))
		base64.StdEncoding.Encode(b, values[i])
		values[i] = b
	}

	return bytes.Join(values, SEPERATOR)
}

func Unpack(value []byte) [][]byte {
	values := bytes.Split(value, SEPERATOR)

	for i := 0; i < len(values); i++ {
		b := make([]byte, base64.StdEncoding.DecodedLen(len(values[i])))
		base64.StdEncoding.Decode(b, values[i])
		values[i] = b
	}

	return values
}

func Uint64ToBytes(value uint64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, value)
	return result
}

func BytesToUint64(value []byte) uint64 {
	return binary.BigEndian.Uint64(value)
}
