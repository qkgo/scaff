package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type PlainFormatter struct {
	Prefix string
	Suffix string
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func (mf *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return FormatWithWrapper(entry)
}

func FormatWithWrapper(entry *logrus.Entry) ([]byte, error) {
	_, fileName, lineNumber, succeed := runtime.Caller(9)
	if !succeed {
		return PrintWithEntry(entry)
	}
	timestamp := time.Now().Local().Format("0102-150405.000")
	file := filepath.Base(fileName)
	if len(file) > 0 {
		var fLen int
		if fLen = utf8.RuneCountInString(file); fLen > 3 {
			fLen = fLen - 3
		}
		msg := fmt.Sprintf("#%s [%-3.3s][%+15.15v:%4.4d][GID:%-6.6d]%-5.5s %s\n",
			timestamp, strings.ToUpper(entry.Level.String()), file[:fLen], lineNumber, getGID(), "", entry.Message)
		return []byte(msg), nil
	} else {
		msg := fmt.Sprintf("#%s [%-3.3s][GID:%-6.6d]%-5.5s %s\n",
			timestamp, strings.ToUpper(entry.Level.String()), getGID(), "", entry.Message)
		return []byte(msg), nil
	}
}

func PrintWithEntry(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("0102-150405.000")
	var file string
	var lineNumber int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		lineNumber = entry.Caller.Line
	}
	if len(file) > 0 {
		var fLen int
		if fLen = utf8.RuneCountInString(file); fLen > 3 {
			fLen = fLen - 3
		}
		msg := fmt.Sprintf("#%s [%-3.3s][%+15.15v:%4.4d][GID:%-6.6d]%-5.5s %s\n",
			timestamp, strings.ToUpper(entry.Level.String()), file[:fLen], lineNumber, getGID(), "", entry.Message)
		return []byte(msg), nil
	} else {
		msg := fmt.Sprintf("#%s [%-3.3s][GID:%-6.6d]%-5.5s %s\n",
			timestamp, strings.ToUpper(entry.Level.String()), getGID(), "", entry.Message)
		return []byte(msg), nil
	}
}
