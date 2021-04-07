package utils

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

func GetHashValue(key string) uint64 {
	// Calculate hash value
	hash := md5.Sum([]byte(key))
	fmt.Println(hash)
	value := binary.LittleEndian.Uint64(hash[0:8])

	return value
}