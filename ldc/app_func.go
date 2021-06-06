package ldc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

// str => md5
func Md5(str string) string {
	result := md5.Sum([]byte(str))
	return string(hex.EncodeToString(result[:]))
}

// to str from obj
func ToStr(obj interface{}) string {
	switch obj.(type) {
	case int:
		return strconv.Itoa(obj.(int))
	case int8:
		return strconv.Itoa(int(obj.(int8)))
	case int16:
		return strconv.Itoa(int(obj.(int16)))
	case int32:
		return strconv.Itoa(int(obj.(int32)))
	case int64:
		return strconv.Itoa(int(obj.(int64)))
	case uint:
		return strconv.Itoa(int(obj.(uint)))
	case uint8:
		return strconv.Itoa(int(obj.(uint8)))
	case uint16:
		return strconv.Itoa(int(obj.(uint16)))
	case uint32:
		return strconv.Itoa(int(obj.(uint32)))
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", obj)
	}
}

// parse int silence
func ParseInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// parse int with err
func ParseIntErr(str string) (int, error) {
	return strconv.Atoi(str)
}
