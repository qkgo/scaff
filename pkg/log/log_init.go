package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/qkgo/scaff/pkg/util/system"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var ProjectName string

func SettingProjectName(projectName string) {
	ProjectName = projectName
}

func GetProjectName() string {
	return ProjectName
}

type LogInitConfiguration struct {
	ProjectName  string
	ProjectPath  string
	PrintLevel   string
	PrintConsole bool
	LogKeepHour  int
}

func InitLogCfg(
	logger **logrus.Logger,
	initCfg *LogInitConfiguration) {
	InitLogByProjectName(
		logger,
		initCfg.ProjectName,
		initCfg.ProjectPath,
		initCfg.PrintLevel,
		initCfg.PrintConsole,
	)
}

func InitLogByProjectName(
	logger **logrus.Logger,
	projectName string,
	projectPath string,
	level string,
	printConsole bool) {
	if projectName == "" {
		log.Println("projectName is not define")
		system.Exit(-1)
		return
	}
	baseLogPath := getOutputPath(projectPath)
	if baseLogPath == "" {
		return
	}
	env := os.Getenv("ENV")
	//if env == "" && cfg.ConfigParam != nil {
	//	env = cfg.ConfigParam.GetString("env")
	//}
	logToFile := os.Getenv("LOG_TO_FILE")
	//if logToFile == "" && cfg.ConfigParam != nil {
	//	logToFile = cfg.ConfigParam.GetString("logtoFile")
	//}
	if logToFile == "" {
		logToFile = "DEFAULT_WRITE_FILE"
	}
	logMaxAgeParam := os.Getenv("LOG_KEEP_HOUR")
	//if logMaxAgeParam == "" && cfg.ConfigParam != nil {
	//	logMaxAgeParam = cfg.ConfigParam.GetString("maxLogKeepHour")
	//}
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
	filename := ""
	if env != "" {
		filename = projectName + "-" + env
		if level != "" {
			filename = filename + "-" + level
		}
	} else if level != "" {
		filename = projectName + "-" + level
	} else {
		filename = projectName
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
		fmt.Printf("config local file system logger error: %v", errors.WithStack(err))
		system.Exit(-1)
		return
	}

	settingLogFormatTypeByOSEnv(logger)

	settingLogLevelByOSEnv(logger)

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

func settingLogFormatTypeByOSEnv(logger **logrus.Logger) {
	logType := os.Getenv("LOG-TYPE")
	if logType == "" {
		logType = os.Getenv("LOG_TYPE")
	}
	//if logType == "" && cfg.ConfigParam != nil {
	//	logType = cfg.ConfigParam.GetString("log.type")
	//}
	switch strings.ToLower(logType) {
	case "logback-json":
		(*logger).SetFormatter(new(JavaJsonFormatter))
	case "json":
		(*logger).SetFormatter(new(logrus.JSONFormatter))
	default:
		(*logger).SetFormatter(new(PlainFormatter))
	}
}

func settingLogLevelByOSEnv(logger **logrus.Logger) {
	logLevelFromEnv := os.Getenv("LOG-LEVEL")
	if logLevelFromEnv == "" {
		logLevelFromEnv = os.Getenv("LOG_LEVEL")
	}
	if logLevelFromEnv != "" {
		settingLogLevel, parsingLogLevelError := logrus.ParseLevel(logLevelFromEnv)
		if parsingLogLevelError != nil {
			log.Printf("logrus log level not match at: %s  has: %v", logLevelFromEnv, parsingLogLevelError)
		} else {
			(*logger).SetLevel(settingLogLevel)
		}
	}
}

func getOutputPath(baseLogPath string) string {
	if baseLogPath == "" {
		baseLogPath = os.Getenv("LOG-PATH")
		if baseLogPath == "" {
			baseLogPath = os.Getenv("LOG_PATH")
		}
		if baseLogPath != "" {
			return baseLogPath
		}
		log.Println("baseLogPath is not define")
		currentPath, err := os.Getwd()
		if err != nil {
			log.Printf("os.Getwd() shutdown. err: %v", err.Error())
			system.Exit(-1)
			return ""
		}
		return path.Join(currentPath, "logs")
	} else {
		return baseLogPath
	}
}
