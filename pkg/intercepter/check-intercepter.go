package intercepter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/util"
	"net/http"
)

var CheckServerHost string

func SetCheckServerHost(newCheckServerHost string) {
	CheckServerHost = newCheckServerHost
}

func CheckRequestProcessContext(requestBodyString []byte, context *gin.Context, currentPermission string) error {
	if cfg.CheckServerStateMap[currentPermission] == false {
		return nil
	}
	var resp *resty.Response
	var err error
	if context.Request.Method == "GET" {
		bodyString := util.JsonParse(map[string]string{
			"path": context.Request.URL.Path,
		})
		resp, err = resty.R().
			SetHeader("Real-Path", context.Request.URL.Path).
			SetHeader("Content-Type", "application/json").
			SetHeader("token", context.GetHeader("token")).
			SetBody(bodyString).
			Post(CheckServerHost + currentPermission + "/check")
	} else if context.Request.Method == "DELETE" {
		bodyString := util.JsonParse(map[string]string{
			"path": context.Request.URL.Path,
		})
		resp, err = resty.R().
			SetHeader("Real-Path", context.Request.URL.Path).
			SetHeader("Content-Type", "application/json").
			SetHeader("token", context.GetHeader("token")).
			SetBody(bodyString).
			Post(CheckServerHost + currentPermission + "/check")
	} else {
		resp, err = resty.R().
			SetHeader("Real-Path", context.Request.URL.Path).
			SetHeader("Content-Type", "application/json").
			SetHeader("token", context.GetHeader("token")).
			SetBody(requestBodyString).
			Post(CheckServerHost + currentPermission + "/check")
	}
	if err != nil {
		cfg.Log.Errorf("%#v", err)
		context.Header("message", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"step":       "query request body failed",
			"message":    "query request body failed",
			"validation": err.Error(),
			"code":       -17})
		return errors.New("check request failed")
	}
	if resp.StatusCode() > 300 && resp.StatusCode() != 502 {
		context.JSON(resp.StatusCode(), gin.H{
			"step":       "check request failed",
			"message":    "check request failed",
			"validation": resp.String(),
			"code":       -41})
		return errors.New("check request failed")
	}
	return nil
}
