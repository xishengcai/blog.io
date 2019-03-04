package utils

import (
	"crypto/md5"
	sha12 "crypto/sha1"
	"encoding/hex"
)

func Sha1(data string) string {
	sha1 := sha12.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum(nil))
}

func Md5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	return hex.EncodeToString(md5.Sum(nil))
}