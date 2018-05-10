package util

import (
	"crypto/md5"
	"encoding/hex"
)
// MD5 Hash
func MD5Hash(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
