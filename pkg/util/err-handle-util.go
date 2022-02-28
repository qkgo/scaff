package util

import (
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"github.com/qkgo/scaff/pkg/cfg"
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
