package log

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"reflect"
)

var Log *logrus.Logger
var LogVerbose *logrus.Logger
var LogConnection *logrus.Logger
var LogStackTrace *logrus.Logger
var LogHttpLog *logrus.Logger
var LogStd *logrus.Logger

func GetRef() **logrus.Logger {
	return &Log
}

func Get() *logrus.Logger {
	return Log
}

func Printf(format string, args ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Print panic: %v", r)
			err = errors.Errorf("Print panic: %v", r)
		}
	}()
	if Log == nil {
		log.Println(args[0].(string), args[1:])
		return
	}
	Log.Infof(format, args...)
	return
}

func Print(format string, args ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Print panic: %v", r)
			err = errors.Errorf("Print panic: %v", r)
		}
	}()
	if Log == nil {
		log.Println(args[0].(string), args[1:])
		return
	}
	Log.Infof(format, args...)
	return
}

func Println(args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if Log == nil {
		log.Println(args[0].(string), args[1:])
		return
	}
	Log.Println(args...)
}

var lastMsg = ""
var sameCount = 0
var maxSameMsg = 5

func I(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil || len(args) == 0 {
		Log.Info(format)
		return
	}
	Log.Infof(format, args...)
}

func D(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if LogVerbose == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil || len(args) == 0 {
		LogVerbose.Debug(format)
		return
	}
	LogVerbose.Debugf(format, args...)
}

func W(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil || len(args) == 0 {
		Log.Warn(format)
		return
	}
	Log.Warnf(format, args...)
}

func E(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil {
		Log.Error(format)
		return
	}
	Log.Errorf(format, args...)
}

func EW(format string, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	log.Println(format, errors.WithStack(err))
}

func Fatalf(format string, args ...interface{}) {
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil {
			log.Fatalf(format)
			return
		}
		log.Fatalf(format, args...)
		return
	}
	if args == nil {
		Log.Fatal(format)
		return
	}
	Log.Fatalf(format, args...)
}

func Fatal(args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Fatal: %v \n", r)
		}
	}()
	if sameMsgProcess(args[0].(string), args[1:]) {
		return
	}
	if Log == nil {
		log.Printf(args[0].(string), args[1:])
		return
	}
	Log.Fatal(args...)
}

func Infof(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Info: %v \n", r)
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil || len(args) == 0 {
		Log.Info(format)
		return
	}
	Log.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil || len(args) == 0 {
		Log.Debug(format)
		return
	}
	Log.Debugf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil || len(args) == 0 {
		Log.Warn(format)
		return
	}
	Log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil {
		Log.Error(format)
		return
	}
	Log.Errorf(format, args...)
}

func V(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if LogVerbose == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil {
		LogVerbose.Debug(format)
		return
	}
	LogVerbose.Debugf(format, args...)
}

func C(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if LogConnection == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil {
		LogConnection.Info(format)
		return
	}
	LogConnection.Infof(format, args...)
}

func T(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if LogStackTrace == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		for i, arg := range args {
			if arg == nil || &arg == nil || reflect.DeepEqual(arg, nil) {
				args[i] = ""
			}
		}
		log.Printf(format, args...)
		return
	}
	if args == nil {
		LogStackTrace.Info(format)
		return
	}
	LogStackTrace.Infof(format, args...)
}

func sameMsgProcess(format string, args []interface{}) bool {
	if lastMsg == format && (args == nil || len(args) == 0) {
		sameCount++
		if sameCount == maxSameMsg {
			log.Println("duplications output")
		}
		if sameCount >= maxSameMsg {
			return true
		}
	} else {
		sameCount = 0
		lastMsg = format
	}
	if format == "" && (args == nil || len(args) == 0) {
		sameCount += maxSameMsg / 2
		if sameCount >= maxSameMsg {
			return true
		}
	}
	return false
}

func Default(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if Log == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil || len(args) == 0 {
		Log.Info(format)
		return
	}
	Log.Infof(format, args...)
}

func HC(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				println(x)
			case error:
				println(x.Error())
			default:
				log.Println(x)
			}
		}
	}()
	if sameMsgProcess(format, args) {
		return
	}
	if LogHttpLog == nil {
		if args == nil || len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args...)
		return
	}
	if args == nil || len(args) == 0 {
		LogHttpLog.Info(format)
		return
	}
	LogHttpLog.Infof(format, args...)
}
