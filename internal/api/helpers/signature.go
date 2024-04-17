package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

var Secret string

func GenerateSignature(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data + Secret))
	return hex.EncodeToString(hasher.Sum(nil))
}
