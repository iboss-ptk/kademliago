package kademliago

import (
	"crypto/sha1"
)

func Hash(data []byte) [20]byte {
	return sha1.Sum(data)
}
