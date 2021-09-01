package code_based_hash

import (
	"crypto/md5"
	"fmt"
)

func MD5(bytes []byte) string {
	instance := md5.New()
	instance.Write(bytes)
	result := instance.Sum(nil)
	return fmt.Sprintf("%x", result)
}
