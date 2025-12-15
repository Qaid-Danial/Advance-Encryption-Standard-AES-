package keygen

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func GenerateKey(a string) string {

	hasher := md5.New()
	io.WriteString(hasher, a)

	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString
	
}
