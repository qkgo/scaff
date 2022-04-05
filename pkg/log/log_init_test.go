package log

import (
	"github.com/qkgo/scaff/pkg/cfg"
	"testing"
)

func TestInitProjectLog(t *testing.T) {
	projectName := "projectName"
	InitLogByProjectName(&cfg.LogDebug, projectName, "./logs", "debug", false)
	InitLogByProjectName(&cfg.LogInfo, projectName, "./logs", "info", true)
	InitLogByProjectName(&cfg.LogHttp, projectName, "./logs", "http", true)
	InitLogByProjectName(&cfg.LogRpc, projectName, "./logs", "key", true)
	InitLogByProjectName(&cfg.SqlLog, projectName, "./logs", "sql", true)
	InitLogByProjectName(&cfg.Log, projectName, "./logs", "default", true)
	t.Logf("init log finished")
}
