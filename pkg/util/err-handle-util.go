package util

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qkgo/scaff/pkg/cfg"
	"io"
)

func ErrorCatch(err error, errorMessage string, c *gin.Context) bool {
	if err == nil {
		return false
	}
	cfg.Log.Info(errorMessage)
	cfg.Log.Info(err)
	c.JSON(400, map[string]interface{}{
		"error":       errorMessage,
		"success":     false,
		"code":        md5.Sum([]byte(errorMessage)),
		"errorDetail": err.Error(),
		"data":        nil,
	})
	return true
}

func ErrorCatchEncryption(err error, errorMessage string, c *gin.Context) bool {
	if err == nil {
		return false
	}
	cfg.Log.Info(errorMessage)
	cfg.Log.Info(err)

	EncryptionSend(200, map[string]interface{}{
		"error":       errorMessage,
		"success":     false,
		"code":        md5.Sum([]byte(errorMessage)),
		"errorDetail": err.Error(),
		"data":        nil,
	}, c)
	return true
}

func MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
