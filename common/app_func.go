package lodash

import (
	"crypto/md5"
	"encoding/hex"
)

// str => md5
func Md5(str string) string {
	result := md5.Sum([]byte(str))
	return string(hex.EncodeToString(result[:]))
}
