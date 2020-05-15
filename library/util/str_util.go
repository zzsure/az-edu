package util

import (
	"crypto/md5"
	"fmt"
)

func GetMd5Str(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
