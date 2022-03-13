package util

import (
	"fmt"
	"github.com/qkgo/scaff/pkg/cfg"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		cfg.Log.Error(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func RecursionListFile(fileName string, resultFileList *[]string, pattern *string) ([]string, error) {
	var reg *regexp.Regexp
	defer func() []string {
		if err := recover(); err != nil {
			log.Printf("RecursionListFile error: %v", err)
		}
		return *resultFileList
	}()
	if resultFileList == nil {
		resultFileList = new([]string)
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return *resultFileList, nil
	}
	fi, err := file.Stat()
	if err != nil {
		log.Println(err)
		return *resultFileList, nil
	}
	if !fi.IsDir() {
		log.Println(fileName, " is not a dir")
		*resultFileList = append(*resultFileList, fileName)
		return *resultFileList, nil
	}
	if pattern != nil && *pattern != "" {
		reg, err = regexp.Compile(*pattern)
		if err != nil {
			log.Println(err)
			return *resultFileList, nil
		}
	}
	// recursively read path
	filepath.Walk(fileName,
		func(currentFilePath string, f os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}
			if fileName == currentFilePath {
				return nil
			}
			//if f.IsDir() {
			//RecursionListFile(currentFilePath, resultFileList, nil)
			//return nil
			//}
			if pattern != nil && *pattern != "" {
				matched := reg.MatchString(f.Name())
				if matched {
					*resultFileList = append(*resultFileList, currentFilePath)
				}
			}
			*resultFileList = append(*resultFileList, currentFilePath)
			return nil
		})
	return *resultFileList, nil
}
