package cfg

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
)

type JavaJsonFormatter struct {
	Prefix string
	Suffix string
}

func (jjf *JavaJsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return JavaJsonFormatWithWrapper(entry)
}

func JavaJsonFormatWithWrapper(entry *logrus.Entry) ([]byte, error) {
	_, fileName, lineNumber, succeed := runtime.Caller(9)
	if !succeed {
		return PrintJavaJsonWithEntry(entry)
	}
	timestamp := time.Now().Local().Format("0102-150405.000")
	file := filepath.Base(fileName)
	if len(file) > 0 {
		var fLen int
		if fLen = utf8.RuneCountInString(file); fLen > 3 {
			fLen = fLen - 3
		}
		logMap := map[string]interface{}{
			"@timestamp": timestamp,
			"level":      strings.ToUpper(entry.Level.String()),
			"class":      file[:fLen],
			"line":       lineNumber,
			"thread":     getGID(),
			"message":    entry.Message,
		}
		return []byte(fmt.Sprintf("%s\n", jsonParse(logMap))), nil
	} else {
		logMap := map[string]interface{}{
			"@timestamp": timestamp,
			"level":      strings.ToUpper(entry.Level.String()),
			"line":       lineNumber,
			"thread":     getGID(),
			"message":    entry.Message,
		}
		return []byte(fmt.Sprintf("%s\n", jsonParse(logMap))), nil
	}
}

func PrintJavaJsonWithEntry(entry *logrus.Entry) ([]byte, error) {
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
		logMap := map[string]interface{}{
			"@timestamp": timestamp,
			"level":      strings.ToUpper(entry.Level.String()),
			"class":      file[:fLen],
			"line":       lineNumber,
			"thread":     getGID(),
			"message":    entry.Message,
		}
		return []byte(fmt.Sprintf("%s\n", jsonParse(logMap))), nil
	} else {
		logMap := map[string]interface{}{
			"@timestamp": timestamp,
			"level":      strings.ToUpper(entry.Level.String()),
			"line":       lineNumber,
			"thread":     getGID(),
			"message":    entry.Message,
		}
		return []byte(fmt.Sprintf("%s\n", jsonParse(logMap))), nil
	}
}

func jsonParse(input interface{}) []byte {
	var jsonByte []byte
	defer func() []byte {
		if err := recover(); err != nil {
			log.Printf("json parse error: %+v", err)
			return []byte(fmt.Sprintf("%#v", input))
		}
		return jsonByte
	}()
	jsonByte, err := json.Marshal(input)
	if err != nil {
		log.Printf("json parse error: %+v", err)
		if jsonByte == nil {
			return []byte(fmt.Sprintf("%#v", input))
		}
	}
	return jsonByte
}
