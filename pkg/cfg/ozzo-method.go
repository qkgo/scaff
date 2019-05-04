package cfg

import (
	"fmt"
	ozzoConfig "github.com/go-ozzo/ozzo-config"
	"os"
	"strconv"
)

type OzzConfig struct{}

func (*OzzConfig) GetConfig() *ozzoConfig.Config{
	return ConfigParam
}

func (*OzzConfig) GetCryptoOption() bool {
	CRYPTO := os.Getenv("CRYPTO")
	if CRYPTO != "" {
		fmt.Println("CRYPTO:", CRYPTO)
		if CRYPTO == "false" {
			return false
		}
		return true
	}
	return ConfigParam.GetBool("crypto.state")
}

func (*OzzConfig) GetForgetUrl() string {
	return ConfigParam.GetString("server.accountforget")
}

func (*OzzConfig) GetTokenHour() int64 {
	return ConfigParam.GetInt64("auth.expire")
}

func (*OzzConfig) IsProduction() bool {
	return ConfigParam.GetString("env") == "prod"
}

func (*OzzConfig) GetHookServerHost() string {
	return ConfigParam.GetString("server.hook")
}

func (*OzzConfig) GetSubHookServerHost(queryPath string) string {
	return ConfigParam.GetString("server.subhook." + queryPath)
}

func (*OzzConfig) GetValidationHost() string {
	return ConfigParam.GetString("server.validation")
}

func (*OzzConfig) GetCentralHost() string {
	return ConfigParam.GetString("host")
}

func (*OzzConfig) GetDialect() string {
	return ConfigParam.GetString("db.auth.dialect")
}

func (*OzzConfig) GetCanvasUrl() string {
	return ConfigParam.GetString("canvas.url")
}

func (*OzzConfig) GetDBAuthUrl() string {
	return ConfigParam.GetString("db.auth.url")
}

func (*OzzConfig) GetDBResourceUrl() string {
	return ConfigParam.GetString("db.rdbms.url")
}

func (*OzzConfig) GetServerBindPort(projectName string) string {
	if ConfigParam.GetInt("api.http."+projectName) != 0 {
		return "0.0.0.0:" + strconv.Itoa(ConfigParam.GetInt("api.http."+projectName))
	}
	return ""
}

func (*OzzConfig) GetRPCUrl() string {
	return "0.0.0.0:" + strconv.Itoa(ConfigParam.GetInt("rpc.port"))
}

func (*OzzConfig) GetSMTPHost() string {
	return ConfigParam.GetString("smtp.host")
}

func (*OzzConfig) GetSMTPPort() int {
	return ConfigParam.GetInt("smtp.port")
}

func (*OzzConfig) GetSMTPUsername() string {
	return ConfigParam.GetString("smtp.username")
}

func (*OzzConfig) GetSMTPPassword() string {
	return ConfigParam.GetString("smtp.password")
}

func (*OzzConfig) GetOssEntPoint() string {
	return ConfigParam.GetString("oss.endpoint")
}

func (*OzzConfig) GetOssAccessKeyId() string {
	return ConfigParam.GetString("oss.accessKeyID")
}

func (*OzzConfig) GetAccessKeySecret() string {
	return ConfigParam.GetString("oss.accessKeySecret")
}

func (*OzzConfig) GetSmtpHost() string {
	return ConfigParam.GetString("smtp.host")
}
