package cfg

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

var ProjectName string

func SettingProjectName(projectName string) {
	ProjectName = projectName
}
func GetProjectName() string {
	return ProjectName
}


func InitLogByProjectNameV3(
	logger **logrus.Logger,
	projectName string,
	level string,
	printConsole bool) {
	if projectName == "" {
		println("projectName is not define")
		os.Exit(-1)
		return
	}
	baseLogPath := ConfigParam.GetString("log.path." + projectName)
	if baseLogPath == "" {
		println("baseLogPath is not define")
		currentPath, err := os.Getwd()
		if err != nil {
			println("os.Getwd() shutdown. err:", err.Error())
			os.Exit(-1)
			return
		}
		baseLogPath = path.Join(currentPath, "logs")
	}
	env := os.Getenv("env")
	if env == "" && ConfigParam != nil {
		env = ConfigParam.GetString("env")
	}
	*logger = logrus.New()
	(*logger).SetReportCaller(true)
	os.MkdirAll(baseLogPath, os.ModePerm)
	var filename string
	if env != "" {
		filename = projectName + "-" + env
	}
	if level != "" {
		filename = filename + "-" + level
	}
	filename = filename + ".%Y-%m-%d-%H.log"
	writer, err := rotatelogs.New(
		path.Join(baseLogPath, filename),
		rotatelogs.WithLinkName(path.Join(baseLogPath, projectName+".log")),
		rotatelogs.WithHandler(rotatelogs.HandlerFunc(func(e rotatelogs.Event) {
			if e.Type() != rotatelogs.FileRotatedEventType {
				return
			}
			println("rotate:", e.(*rotatelogs.FileRotatedEvent).PreviousFile(), e.(*rotatelogs.FileRotatedEvent).CurrentFile())
		})),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		println("config local file system logger error:", err.Error())
		fmt.Printf("config local file system logger error: %v", errors.WithStack(err))
		os.Exit(-1)
		return
	}
	if printConsole {
		writers := []io.Writer{
			writer,
			os.Stdout}
		fileAndStdoutWriter := io.MultiWriter(writers...)
		(*logger).SetOutput(fileAndStdoutWriter)
	} else {
		(*logger).SetOutput(writer)
	}
	(*logger).SetFormatter(new(MyFormatter))
	(*logger).Info("init log config success")
}


