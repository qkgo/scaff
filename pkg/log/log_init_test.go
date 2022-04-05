package log

import (
	"github.com/qkgo/scaff/pkg/cfg"
	"testing"
)

func TestInitProjectLog(t *testing.T) {
	projectName := "projectName"
	InitLogByProjectName(&cfg.LogDebug, projectName, "debug", false)
	InitLogByProjectName(&cfg.LogInfo, projectName, "info", true)
	InitLogByProjectName(&cfg.LogHttp, projectName, "http", true)
	InitLogByProjectName(&cfg.LogRpc, projectName, "key", true)
	InitLogByProjectName(&cfg.SqlLog, projectName, "sql", true)
	InitLogByProjectName(&cfg.Log, projectName, "default", true)
	t.Logf("init log finished")
}
