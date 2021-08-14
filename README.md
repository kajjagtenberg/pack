# Pack

Library to easiliy and predictably pack and namespace byte keys for key-value stores.

## Install

Use the following command to install:

```
go get github.com/kajjagtenberg/pack
```

## Usage

```Go
namespace := []byte("users")
user := []byte("johndoe123")
version := pack.Uint64ToBytes(uint64(100))

key := pack.Pack(
    namespace,
    user,
    version
)

// Set a key in your favorite key-value store
db.Set(key, value)
```

## API

```Go
func Pack(values ...[]byte) []byte

func Unpack(value []byte) [][]byte

func Uint64ToBytes(value uint64) []byte

func BytesToUint64(value []byte) uint64
```