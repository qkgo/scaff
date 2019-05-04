package cfg

import (
	"log"
	"os"
	"time"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func InitLog(projectName string) {
	InitLogByProjectNameV3(&Log, projectName, "debug", false)
	InitLogByProjectNameV3(&LogInfo, projectName, "info", true)
	InitLogByProjectNameV3(&LogHttp, projectName, "http", true)
	InitLogByProjectNameV3(&LogKey, projectName, "key", true)
	InitLogByProjectNameV3(&SqlLog, projectName, "sql", true)
}

func InitLogCcs(projectName string) {
	Log = logrus.New()
	var baseLogPath string
	if os.Getenv("projectName") != "" {
		baseLogPath = os.Getenv("projectName")
	} else if ConfigParam.GetString("log.path."+projectName) != "" {
		baseLogPath = ConfigParam.GetString("log.path." + projectName)
	}
	var env string
	if os.Getenv("env") != "" {
		baseLogPath = os.Getenv("env")
	} else if ConfigParam.GetString("env") != "" {
		baseLogPath = ConfigParam.GetString("env")
	}
	Log.SetFormatter(&logrus.TextFormatter{})
	Log.SetReportCaller(true)

	writer, err := rotatelogs.New(
		baseLogPath+"-default-"+env+".%Y-%m-%d-%H.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	(*Log).SetFormatter(new(MyFormatter))
	if err != nil {
		log.Println("[Init]: config local file system logger error=", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})
	Log.AddHook(lfHook)

}
