package log

import (
	"github.com/sirupsen/logrus"
	"xorm.io/core"
)

type SqlLogger struct {
	Log       *logrus.Logger
	isShowSql bool
}

func (s SqlLogger) Debug(v ...interface{}) {
	s.Log.Debug(v)
}

func (s SqlLogger) Debugf(format string, v ...interface{}) {
	s.Log.Debugf(format, v)
}

func (s SqlLogger) Error(v ...interface{}) {
	s.Log.Error(v)
}

func (s SqlLogger) Errorf(format string, v ...interface{}) {
	s.Log.Errorf(format, v)
}

func (s SqlLogger) Info(v ...interface{}) {
	s.Log.Info(v)
}

func (s SqlLogger) Infof(format string, v ...interface{}) {
	s.Log.Infof(format, v)
}

func (s SqlLogger) Warn(v ...interface{}) {
	s.Log.Warn(v)
}

func (s SqlLogger) Warnf(format string, v ...interface{}) {
	s.Log.Warnf(format, v)
}

func (s SqlLogger) Level() core.LogLevel {
	// todo parseLevel to iota
	return core.LOG_INFO
}

func (s SqlLogger) SetLevel(l core.LogLevel) {
	//TODO implement me
}

func (s SqlLogger) ShowSQL(show ...bool) {
	s.isShowSql = show[0]
}

func (s SqlLogger) IsShowSQL() bool {
	return s.isShowSql
}
