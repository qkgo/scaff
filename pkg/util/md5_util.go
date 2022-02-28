package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5Sum(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func MD5(inputBytes []byte) string {
	h := md5.New()
	h.Write(inputBytes)
	return hex.EncodeToString(h.Sum(nil))
}
