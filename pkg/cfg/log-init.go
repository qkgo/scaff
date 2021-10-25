package cfg

func InitLog(projectName string) {
	InitLogByProjectNameV3(&LogDebug, projectName, "debug", false)
	InitLogByProjectNameV3(&LogInfo, projectName, "info", true)
	InitLogByProjectNameV3(&LogHttp, projectName, "http", true)
	InitLogByProjectNameV3(&LogRpc, projectName, "key", true)
	InitLogByProjectNameV3(&SqlLog, projectName, "sql", true)
}