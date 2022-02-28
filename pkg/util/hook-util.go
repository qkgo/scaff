package util

import (
	"errors"
	"fmt"
	"github.com/qkgo/scaff/pkg/cfg"
	"gopkg.in/resty.v1"
	"time"
)

var validationServerHost string

var hookServerHost string

var QueryPathMapper = map[string]string{}

const NanoSecondRate float64 = 1000000000

func InitValidationServerHost() string {
	if validationServerHost != "" {

	} else {
		validationServerHost = cfg.OzConfig.GetValidationHost()
	}
	if validationServerHost == "" {
		cfg.LogInfo.Error("hook server url not found")
	}
	return validationServerHost
}

func InitHookServerHost(queryPath string) (string, string) {
	if queryPath != "" && QueryPathMapper[queryPath] != "" {
		QueryPathMapper[queryPath] = cfg.OzConfig.GetSubHookServerHost(queryPath)
	}
	if hookServerHost != "" {

	} else {
		hookServerHost = cfg.OzConfig.GetHookServerHost()
	}
	if validationServerHost == "" {
		cfg.LogInfo.Error("hook server url not found")
	}
	return hookServerHost, hookServerHost
}

func callRemoteValidationServer(checkUrl string, requestBody string) (*resty.Response, error) {
	if validationServerHost == "" {
		cfg.Log.Info("validation server url not found")
		return nil, errors.New("validation server url not found")
	}
	resp, err := resty.R().
		SetHeader("content-type", "application/json").
		SetBody(requestBody).
		Post(validationServerHost + checkUrl + "/check")
	if err != nil {
		cfg.Log.Info(err)
		return resp, err
	}
	return resp, err
}

func PreRequestCheck(checkUrl string, requestBody []byte) error {
	InitValidationServerHost()
	resp, err := callRemoteValidationServer(checkUrl, string(requestBody))
	if err != nil {
		return err
	}
	resBody := resp.Body()
	cfg.Log.Info("PreRequestCheck", string(resBody))
	var validationResponseObject map[string]interface{}
	if err = json.Unmarshal(resBody, &validationResponseObject); err != nil {
		return err
	}
	if resp.StatusCode() != 404 && resp.StatusCode() > 300 {
		if validationResponseObject != nil && validationResponseObject["path"] != "" {
			return errors.New(validationResponseObject["path"].(string))
		}
		return errors.New("status code is" + resp.Status() + ", and response string is :" + string(resBody))
	}
	return nil
}

func callRemoteHookServer(hookUrl string, hookPath string, requestBody string, token string) (*resty.Response, error) {
	startTime := time.Now()
	println("webhookurl:", hookUrl, "requestbody:", requestBody, "t1:", startTime.Format(time.RFC3339))
	cfg.Log.Info("webhookurl:", hookUrl, "requestbody:", requestBody, "t1:", startTime.Format(time.RFC3339))
	if hookUrl != "" {
		resp, err := resty.R().
			SetHeader("content-type", "application/json").
			SetHeader("token", token).
			SetBody(requestBody).
			Post(hookUrl)
		endTime := time.Now()
		distance := fmt.Sprintf("%.3f", float64(endTime.Sub(startTime).Nanoseconds())/NanoSecondRate)
		if err != nil {
			println("resp.webhookurl:", hookUrl, ".error:", err.Error(), "t2:",
				distance, "s")
			cfg.LogInfo.Info("resp.webhookurl:", hookUrl, ".error:", err.Error(), "t2:",
				distance, "s")
			return resp, err
		}
		println("resp.webhookurl:", hookUrl, ".status:", resp.Status(), "body:", string(resp.Body()), "t2:",
			distance, "s")
		cfg.LogInfo.Info("resp.webhookurl:", hookUrl, ".status:", resp.Status(), "body:", string(resp.Body()), "t2:",
			distance, "s")
		return resp, err
	}
	if hookServerHost == "" {
		cfg.Log.Error("hook server url not found")
		return nil, errors.New("hook server url not found")
	}
	secondUrl := hookServerHost + hookPath + ":check"
	resp, err := resty.R().
		SetHeader("content-type", "application/json").
		SetBody(requestBody).
		Post(secondUrl)
	endTime := time.Now()
	distance := fmt.Sprintf("%.3f", float64(endTime.Sub(startTime).Nanoseconds())/NanoSecondRate)
	if err != nil {
		println("resp.webhookurl:", hookUrl, ".error:", err.Error(), "t2:",
			distance, "s")
		cfg.LogInfo.Info("resp.webhookurl:", hookUrl, ".error:", err.Error(), "t2:",
			distance, "s")
		return resp, err
	}
	println("resp.webhookurl:", hookUrl, ".status:", resp.Status(), "body:", string(resp.Body()), "t2:",
		distance, "s")
	cfg.LogInfo.Info("resp.webhookurl:", hookUrl, ".status:", resp.Status(), "body:", string(resp.Body()), "t2:",
		distance, "s")
	return resp, err
}

func PostRequestHook(hookPath string, requestBody []byte, token string, hookUrl string, requestPath string, requestMethod string, extraString string) error {
	resp, err := callRemoteHookServer(hookUrl, hookPath, string(requestBody), token)
	if err != nil {
		return err
	}
	resBody := resp.Body()
	cfg.LogInfo.Info("PostRequestHook:", string(resBody))
	var validationResponseObject map[string]interface{}
	if err = json.Unmarshal(resBody, &validationResponseObject); err != nil {
		return err
	}
	if resp.StatusCode() != 404 && resp.StatusCode() > 300 {
		if validationResponseObject != nil && validationResponseObject["path"] != "" {
			return errors.New(validationResponseObject["path"].(string))
		}
		return errors.New("status code is" + resp.Status() + ", and response string is :" + string(resBody))
	}
	return nil
}
