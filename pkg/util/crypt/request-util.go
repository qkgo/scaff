package crypt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/util/seconds"
	"io/ioutil"
	"strconv"
	"time"
)

var tokenKey []byte
var mpKey []byte

type CryptFunc func(input []byte) ([]byte, error)

var decryptFunc CryptFunc
var encryptFunc CryptFunc

func SetCryptFunc(
	crypt ...CryptFunc) {
	if len(crypt) > 1 {
		decryptFunc = crypt[0]
		encryptFunc = crypt[1]
	}
}

func DecryptionByte(encryptionData string) []byte {
	originString, err := base64.StdEncoding.DecodeString(encryptionData)
	if err != nil {
		cfg.LogInfo.Info(err)
		return nil
	}
	if decryptFunc == nil {
		return originString
	}
	stringRequest, err := decryptFunc(originString)
	if err != nil {
		cfg.LogInfo.Info(err)
		return nil
	}
	return stringRequest
}

func DecryptionByteArray(encryptionData []byte) []byte {
	originString, err := base64.StdEncoding.DecodeString(string(encryptionData))
	if err != nil {
		cfg.LogInfo.Info(err)
		return nil
	}
	if decryptFunc == nil {
		return originString
	}
	stringRequest, err := decryptFunc(originString)
	if err != nil {
		cfg.LogInfo.Info(err)
		return nil
	}
	return stringRequest
}

func GetDecryptedString(context *gin.Context) []byte {
	cryptoData, err := context.GetRawData()
	if err != nil {
		cfg.LogInfo.Info(err)
		context.Abort()
		context.JSON(401, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"code":    -1,
		})
		return nil
	}
	if cryptoData == nil || len(cryptoData) < 1 {
		cfg.LogInfo.Info("request data is null")
		context.Abort()
		context.JSON(401, map[string]interface{}{
			"success": false,
			"error":   "request data is null",
			"code":    -1,
		})
		return nil
	}
	if !cfg.OzConfig.GetCryptoOption() {
		return cryptoData
	}
	if decryptFunc == nil {
		return cryptoData
	}
	originString, err := base64.StdEncoding.DecodeString(string(cryptoData))
	if err != nil {
		cfg.LogInfo.Info(err)
		context.Abort()
		context.JSON(401, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"code":    -2,
		})
		return nil
	}
	stringRequest, err := decryptFunc(originString)
	if err != nil {
		cfg.LogInfo.Info(err)
		context.Abort()
		context.JSON(401, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"code":    -3,
		})
		return nil
	}
	return stringRequest
}

func EncryptionString(jsonRes string) (string, error) {
	if !cfg.OzConfig.GetCryptoOption() {
		return jsonRes, nil
	}
	if encryptFunc == nil {
		return jsonRes, nil
	}
	stringRequest, err := encryptFunc([]byte(jsonRes))
	if err != nil {
		return "", err
	}
	resultString := base64.StdEncoding.EncodeToString(stringRequest)
	return resultString, nil
}

func EncryptionInterface(resultJson map[string]interface{}) (string, error) {
	jsonRes, _ := json.Marshal(resultJson)
	if !cfg.OzConfig.GetCryptoOption() {
		return string(jsonRes), nil
	}
	if encryptFunc == nil {
		return string(jsonRes), nil
	}
	stringRequest, err := encryptFunc(jsonRes)
	if err != nil {
		return "", err
	}
	resultString := base64.StdEncoding.EncodeToString(stringRequest)
	return resultString, nil
}

func EncryptionInterfaceLog(resultJson map[string]interface{}) string {
	jsonRes, _ := json.Marshal(resultJson)
	if !cfg.OzConfig.GetCryptoOption() {
		return string(jsonRes)
	}
	if encryptFunc == nil {
		return string(jsonRes)
	}
	stringRequest, err := encryptFunc(jsonRes)
	if err != nil {
		cfg.LogInfo.Info(err)
		return ""
	}
	resultString := base64.StdEncoding.EncodeToString(stringRequest)
	return resultString
}

func EncryptionSend(code int, resultJson map[string]interface{}, context *gin.Context) {
	jsonRes, _ := json.Marshal(resultJson)
	if !cfg.OzConfig.GetCryptoOption() {
		context.Header("Content-Type", "application/octet-stream")
		context.Data(code, "application/octet-stream", jsonRes)
		return
	}
	if encryptFunc == nil {
		context.Header("Content-Type", "application/octet-stream")
		context.Data(code, "application/octet-stream", jsonRes)
		return
	}
	responseByte := []byte(jsonRes)
	stringRequest, err := encryptFunc(responseByte)
	if err != nil {
		cfg.LogInfo.Info(err)
		context.Header("Content-Type", "application/octet-stream")
		context.Data(400, "application/octet-stream", []byte(err.Error()))
		return
	}
	resultString := base64.StdEncoding.EncodeToString(stringRequest)
	context.Data(code, "application/octet-stream", []byte(resultString))
}

func EncryptionStreamSend(code int, result []byte, context *gin.Context) {
	if encryptFunc == nil {
		context.Header("Content-Type", "application/octet-stream")
		context.Data(code, "application/octet-stream", result)
		return
	}
	stringRequest, err := encryptFunc(result)
	if err != nil {
		cfg.LogInfo.Info(err)
		context.Header("Content-Type", "application/octet-stream")
		context.Data(400, "application/octet-stream", []byte("Error"))
		return
	}
	resultString := base64.StdEncoding.EncodeToString(stringRequest)
	context.Header("Content-Type", "application/octet-stream")
	context.Data(code, "application/octet-stream", []byte(resultString))
	return
}

func EncryptionByte(resultStr []byte) (string, error) {
	stringRequest, err := encryptFunc(resultStr)
	if err != nil {
		return "", err
	}
	resultString := base64.StdEncoding.EncodeToString(stringRequest)
	return resultString, nil
}

func EncryptionByteNoError(resultStr []byte) string {
	stringRequest, err := encryptFunc(resultStr)
	if err != nil {
		cfg.LogInfo.Info(err)
		return ""
	}
	resultString := base64.StdEncoding.EncodeToString(stringRequest)
	return resultString
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	newBody, _ := EncryptionByte(b)
	return w.ResponseWriter.Write([]byte(newBody))
}

func Crypto() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Header("Request-UUID", strconv.Itoa(int(startTime.UnixNano())))
		println("id:", startTime.UnixNano(), ",req:", c.Request.RequestURI, ",method:", c.Request.Method, ",t1:", startTime.Format(time.RFC3339))
		cfg.LogInfo.Info("id:", startTime.UnixNano(), ",req:", c.Request.RequestURI, ",method:", c.Request.Method, ",t1:", startTime.Format(time.RFC3339))
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			requestBody := GetDecryptedString(c)
			if requestBody == nil {
				c.Abort()
				return
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
		}
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		c.Writer = blw
		c.Header("Content-Type", "application/octet-stream")
		endTime := time.Now()
		distance := fmt.Sprintf("%.3f", float64(endTime.Sub(startTime).Nanoseconds())/seconds.NanoSecondRate)
		email, err := c.Get("email")
		if err {
			email = "-"
		}
		println("id:", startTime.UnixNano(), ",req:", c.Request.RequestURI, ",method:", c.Request.Method, ",by:", email, ",ip:", c.Request.RemoteAddr, "t2:", distance, "s")
		cfg.LogInfo.Info("id:", startTime.UnixNano(), ",req:", c.Request.RequestURI, ",method:", c.Request.Method, ",by:", email, ",ip:", c.Request.RemoteAddr, "t2:", distance, "s")
		return
	}
}

func QuietCrypto() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			requestBody := GetDecryptedString(c)
			if requestBody == nil {
				c.Abort()
				return
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
		}
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		c.Writer = blw
		c.Header("Content-Type", "application/octet-stream")
		return
	}
}

func CryptPrivate() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Header("Request-UUID", strconv.Itoa(int(startTime.UnixNano())))
		println("req:", c.Request.RequestURI, ",method:", c.Request.Method, ",t1:", startTime.Format(time.RFC3339))
		cfg.LogInfo.Info("req:", c.Request.RequestURI, ",method:", c.Request.Method, ",t1:", startTime.Format(time.RFC3339))
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			requestBody := GetDecryptedString(c)
			if requestBody == nil {
				c.Abort()
				return
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
		}

		tokenString := c.GetHeader("token")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return tokenKey, nil
		})
		if err != nil {
			cfg.Log.Error("rua210:", err.Error())
			c.Abort()
			c.JSON(400, map[string]interface{}{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		tokenMap := SignData{}
		subject := token.Claims.(jwt.MapClaims)["sub"].(string)
		err = json.Unmarshal([]byte(subject), &tokenMap)
		if err != nil {
			cfg.Log.Error("rua224:", err.Error())
			c.Abort()
			c.JSON(400, map[string]interface{}{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.Set("role", tokenMap)
		println("id:", startTime.UnixNano(), ",req:", c.Request.RequestURI, ",method:", c.Request.Method, ",by:", tokenMap.Email, "ip:", c.Request.RemoteAddr)
		cfg.LogInfo.Info("id:", startTime.UnixNano(), ",req:", c.Request.RequestURI, ",method:", c.Request.Method, ",by:", tokenMap.Email, "ip:", c.Request.RemoteAddr)
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		c.Writer = blw
		c.Header("Content-Type", "application/octet-stream")
		endTime := time.Now()
		distance := fmt.Sprintf("%.3f", float64(endTime.Sub(startTime).Nanoseconds())/seconds.NanoSecondRate)
		println("id:", startTime.UnixNano(), ",req:", c.Request.RequestURI, ",method:", c.Request.Method, ",by:", tokenMap.Email, ",ip:", c.Request.RemoteAddr, ",t2:", distance, "s")
		cfg.LogInfo.Info("id:", startTime.UnixNano(), ",req:", c.Request.RequestURI, ",method:", c.Request.Method, ",by:", tokenMap.Email, "ip:", c.Request.RemoteAddr, ",t2:", distance, "s")
		return
	}
}

func TokenRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return tokenKey, nil
		})
		if err != nil {
			cfg.Log.Error("rua210:", err.Error())
			c.Abort()
			c.JSON(400, map[string]interface{}{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		tokenMap := SignData{}
		subject := token.Claims.(jwt.MapClaims)["sub"].(string)
		err = json.Unmarshal([]byte(subject), &tokenMap)
		if err != nil {
			cfg.Log.Error("rua224:", err.Error())
			c.Abort()
			c.JSON(400, map[string]interface{}{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.Set("role", tokenMap)
		c.Set("email", tokenMap.Email)
		println(c.Request.URL.String(), ":", tokenMap.Email, ":", c.Request.RemoteAddr)
		c.Next()
		return
	}
}

func AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return tokenKey, nil
		})
		if err != nil {
			cfg.Log.Error(err.Error())
			c.Abort()
			c.JSON(400, map[string]interface{}{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		tokenMap := SignData{}
		subject := token.Claims.(jwt.MapClaims)["sub"].(string)
		err = json.Unmarshal([]byte(subject), &tokenMap)
		if err != nil {
			cfg.Log.Error(err.Error())
			c.Abort()
			c.JSON(400, map[string]interface{}{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.Set("role", tokenMap)
	}
}

type SignData struct {
	Id           int      `json:"id" `
	Email        string   `json:"email" `
	CustomerId   int64    `json:"custom" `
	CustomerCode string   `json:"customerCode" `
	Granted      string   `json:"granted" `
	Level        int64    `json:"userLevel" `
	PhoneNumber  string   `json:"phoneNumber"`
	JiraAccount  string   `json:"jiraAccount"`
	JiraPassword string   `json:"jiraPassword"`
	RoleCode     []string `json:"role" `
	GroupCode    []string `json:"groupCode"`
}

func CheckTokenString(tokenString string) (*SignData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		println(err.Error())
		return nil, err
	}
	tokenMap := SignData{}
	subject := token.Claims.(jwt.MapClaims)["sub"].(string)
	err = json.Unmarshal([]byte(subject), &tokenMap)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return &tokenMap, nil
}

func CheckToken(context *gin.Context) (*SignData, error) {
	tokenString, err := context.Cookie("token")
	if err != nil || tokenString == "" {
		tokenString = context.GetHeader("token")
	}
	if tokenString == "" {
		return nil, errors.New("token is null")
	}
	println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		println(err.Error())
		return nil, err
	}
	tokenMap := SignData{}
	subject := token.Claims.(jwt.MapClaims)["sub"].(string)
	err = json.Unmarshal([]byte(subject), &tokenMap)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return &tokenMap, nil
}

func CheckTokenPlain(context *gin.Context) (*SignData, error) {
	tokenString := context.GetHeader("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil {
		cfg.Log.Error(err.Error())
		return nil, err
	}
	tokenMap := SignData{}
	subject := token.Claims.(jwt.MapClaims)["sub"].(string)
	err = json.Unmarshal([]byte(subject), &tokenMap)
	if err != nil {
		cfg.Log.Error(err.Error())
		return nil, err
	}
	return &tokenMap, nil
}

func GetRawString(context *gin.Context) []byte {
	data, err := context.GetRawData()
	if err != nil {
		cfg.LogInfo.Info(err)
	}
	return data
}

func GetDeCryptoString(context *gin.Context) []byte {
	cryptoData, err := context.GetRawData()
	if err != nil {
		cfg.LogInfo.Info(err)
		context.Abort()
		context.JSON(401, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"code":    -1,
		})
		return nil
	}
	if cryptoData == nil || len(cryptoData) < 1 {
		cfg.LogInfo.Info("request data is null")
		context.Abort()
		context.JSON(401, map[string]interface{}{
			"success": false,
			"error":   "request data is null",
			"code":    -1,
		})
		return nil
	}
	if !cfg.OzConfig.GetCryptoOption() {
		return cryptoData
	}
	originString, err := base64.StdEncoding.DecodeString(string(cryptoData))
	if err != nil {
		cfg.LogInfo.Info(err)
		context.Abort()
		context.JSON(401, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"code":    -2,
		})
		return nil
	}
	if decryptFunc == nil {
		stringRequest := originString
		return stringRequest
	}
	//cfg.LogInfo.Info(originString)
	stringRequest, decryptErr := decryptFunc(originString)
	if decryptErr != nil {
		cfg.LogInfo.Info(decryptErr)
		context.Abort()
		context.JSON(401, map[string]interface{}{
			"success": false,
			"error":   decryptErr.Error(),
			"code":    -3,
		})
		return nil
	}
	return stringRequest
}
