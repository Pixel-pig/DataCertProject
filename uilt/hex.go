package uilt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5HexToString(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}
