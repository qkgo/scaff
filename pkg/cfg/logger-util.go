package cfg

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type MyFormatter struct {
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

func (mf *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("0102-150405.000")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	var fLen int
	if fLen = utf8.RuneCountInString(file); fLen > 3 {
		fLen = fLen - 3
	}
	msg := fmt.Sprintf("%s [%3.3s][%+15.15v:%4.4d][GOID:%3.3d] %s\n", timestamp, strings.ToUpper(entry.Level.String()), file[:fLen], len, getGID(), entry.Message)
	return []byte(msg), nil
}
