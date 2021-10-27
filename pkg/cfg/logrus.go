package cfg

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
	//"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
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
	env := os.Getenv("ENV")
	if env == "" && ConfigParam != nil {
		env = ConfigParam.GetString("env")
	}
	logToFile := os.Getenv("LOG_TO_FILE")
	if logToFile == "" && ConfigParam != nil {
		logToFile = ConfigParam.GetString("logtoFile")
	}
	if logToFile == "" {
		logToFile = "DEFAULT_WRITE_FILE"
	}
	logMaxAgeParam := os.Getenv("LOG_KEEP_HOUR")
	if logMaxAgeParam == "" && ConfigParam != nil {
		logMaxAgeParam = ConfigParam.GetString("maxLogKeepHour")
	}
	var logMaxAge time.Duration
	if logMaxAgeParam == "" {
		logMaxAge = 2 * time.Hour
	} else {
		paramHour, err := strconv.Atoi(logMaxAgeParam)
		if err != nil {
			logMaxAge = 2 * time.Hour
		} else {
			logMaxAge = time.Duration(paramHour) * time.Hour
		}
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
	logFilePath := path.Join(baseLogPath, filename)
	writer, err := rotatelogs.New(
		logFilePath,
		rotatelogs.WithHandler(rotatelogs.HandlerFunc(func(e rotatelogs.Event) {
			if e.Type() != rotatelogs.FileRotatedEventType {
				return
			}
			println("rotate:", e.(*rotatelogs.FileRotatedEvent).PreviousFile(), e.(*rotatelogs.FileRotatedEvent).CurrentFile())
		})),
		rotatelogs.WithMaxAge(logMaxAge),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		println("config local file system logger error:", err.Error())
		fmt.Printf("config local file system logger error: %v", errors.WithStack(err))
		os.Exit(-1)
		return
	}
	logType := os.Getenv("LOG_TYPE")
	if logType == "" && ConfigParam != nil {
		logType = ConfigParam.GetString("log.type")
	}
	if logType == "json" {
		(*logger).SetFormatter(new(logrus.JSONFormatter))
	} else {
		(*logger).SetFormatter(new(MyFormatter))
	}
	if logToFile == "" {
		fmt.Printf("%6.9s;%6.9s; %9.9s;  [logtoFile]:false , write stdout\n", projectName, env, level)
		(*logger).SetOutput(os.Stdout)
	} else if printConsole {
		fmt.Printf("%6.9s;%6.9s; %9.9s;  [logtoFile]:true [printConsole]:true, write stdout and file: %s \n", projectName, env, level, logFilePath)
		writers := []io.Writer{
			writer,
			os.Stdout}
		fileAndStdoutWriter := io.MultiWriter(writers...)
		(*logger).SetOutput(fileAndStdoutWriter)
	} else {
		fmt.Printf("%6.9s;%6.9s; %9.9s;  [logtoFile]:ture [printConsole]:false, write file: %s \n", projectName, env, level, logFilePath)
		(*logger).SetOutput(writer)
	}

	go (*logger).Infof("%9.9s;%9.9s; %9.9s; - init log succeed", projectName, env, level)
}
