package util

import (
	"fmt"
	"github.com/qkgo/scaff/pkg/cfg"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		cfg.Log.Error(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func RecursionListFile(fileName string, pattern string) []string {
	var reg *regexp.Regexp
	var resultFileList []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return nil
	}
	fi, err := file.Stat()
	if err != nil {
		log.Println(err)
		return nil
	}
	if !fi.IsDir() {
		log.Println(fileName, " is not a dir")
		return nil
	}
	if pattern != "" {
		reg, err = regexp.Compile(pattern)
		if err != nil {
			log.Println(err)
			return nil
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 遍历目录
	filepath.Walk(fileName,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}
			if f.IsDir() {
				resultFileList = append(resultFileList, RecursionListFile(fileName, pattern)...)
				return nil
			}
			if pattern != "" {
				matched := reg.MatchString(f.Name())
				if matched {
					resultFileList = append(resultFileList, path)
				}
			}
			resultFileList = append(resultFileList, path)
			wg.Done()
			return nil
		})
	wg.Wait()
	return resultFileList
}
