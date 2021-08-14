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

key := pack.Pack(
    namespace,
    user,
) // dXNlcnM=:am9obmRvZTEyMw==

// Set a key in your favorite key-value store
db.Set(key, value)
```