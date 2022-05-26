package filesystem

import (
	"crypto/md5"
	"fmt"
	"github.com/karrick/godirwalk"
	"github.com/pkg/errors"
	"github.com/qkgo/scaff/pkg/util/system"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func FileStat(fileName string) (*os.File, os.FileInfo, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}
	return file, stat, nil
}

func FileMd5(filename string) (string, int64, error) {
	f, err := os.Open(filename)
	if nil != err {
		log.Println(err)
		return "", 0, err
	}
	defer func(f *os.File) {
		errOpen := f.Close()
		if errOpen != nil {
			log.Printf("open file [%s]  error: %v  \n", filename, errOpen)
		}
	}(f)
	st, err := f.Stat()
	if err != nil {
		errorMsg := fmt.Sprintf("open file [%s] has error: %v  \n", filename, err)
		log.Println(errorMsg)
		err = errors.Wrap(err, errorMsg)
		return "", st.Size(), err
	}
	if st.IsDir() || st.Size() == 0 {
		fileMd5 := fmt.Sprintf("%x", md5.Sum([]byte(filename)))
		return fileMd5, st.Size(), nil
	}
	md5Handle := md5.New()
	_, err = io.Copy(md5Handle, f)
	if nil != err {
		log.Println(err)
		return "", st.Size(), err
	}
	md := md5Handle.Sum(nil)
	md5str := fmt.Sprintf("%x", md)
	return md5str, st.Size(), nil
}

func Mkdirp(path string) {
	go func() {
		if path == "" {
			executePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				log.Println("marking dir error:  path is null, getting pwd at: ", os.Args[0], " err:", err)
				return
			}
			log.Println("using execute path:", executePath)
			path = executePath
		}
		log.Println("starting mkdirp :", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println("creating dir err:", path, " - ", err)
		}
	}()
}

func MkdirpSync(path string) (mkdirError error) {
	if path == "" {
		executePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Println("marking dir error:  path is null, getting pwd at: ", os.Args[0], " err:", err)
			return
		}
		log.Println("using execute path:", executePath)
		path = executePath
	}
	log.Println("starting mkdirp :", path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Printf("creating dir err: %s - %v ", path, err)
	}
	filepathResult, err := os.Stat(path)
	if err != nil {
		log.Printf("stat by path: %s , err: %v", path, err)
	}
	if filepathResult == nil {
		mkdirError = errors.Errorf("unexpect filepath is empty: %s", path)
		return
	}
	if filepathResult.IsDir() {
		return nil
	} else {
		mkdirError = errors.Errorf("unexpect filepath type is file: %s", path)
	}
	mkdirError = err
	return
}

func MkdirpList(paths []string) {
	go func() {
		for _, path := range paths {
			if path == "" {
				log.Printf("exiting program because configuration has syntax null error")
				system.Exit(-5)
				return
			}
			if path == "" {
				executePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
				if err != nil {
					log.Println("marking dir error:  path is null, getting pwd at: ", os.Args[0], " err:", err)
					return
				}
				log.Println("using execute path:", executePath)
				path = executePath
			}
			log.Println("starting mkdirp :", path)
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				log.Println("creating dir err:", path, " - ", err)
			}
		}
	}()
}

func GetFilesByFilter(currentPath string, filenameFilter string) (error, []string) {
	var files []string
	err := godirwalk.Walk(currentPath, &godirwalk.Options{
		Unsorted: false,
		Callback: func(path string, info *godirwalk.Dirent) error {
			if strings.Contains(path, filenameFilter) {
				files = append(files, path)
			}
			return nil
		},
	})
	sort.Strings(files)
	files = reverse(files)
	return err, files
}

func GetFiles(currentPath string) (error, []string) {
	var files []string
	err := godirwalk.Walk(currentPath, &godirwalk.Options{
		Unsorted: false,
		Callback: func(path string, info *godirwalk.Dirent) error {
			if strings.Contains(path, ".log") || strings.Contains(path, ".bag") {
				files = append(files, path)
			}
			return nil
		},
	})
	sort.Strings(files)
	files = reverse(files)
	return err, files
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Exists(fileFullPath string) bool {
	_, err := os.Stat(fileFullPath)
	if err == nil {
		// path/to/whatever exists

		return true
	} else if os.IsNotExist(err) {
		log.Printf("file %s is not exist %v \n", fileFullPath, err)
		// path/to/whatever does *not* exist
		return false
	} else {
		log.Printf("stat %s error %v \n", fileFullPath, err)
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		return false
	}
}

func FindExitingFile(fileList ...string) string {
	for _, checkingExitingFile := range fileList {
		if checkingExitingFile != "" && Exists(checkingExitingFile) {
			return checkingExitingFile
		}
	}
	return ""
}

func ExistsFile(fileFullPath string) bool {
	if fsStat, err := os.Stat(fileFullPath); err == nil {
		// path/to/whatever exists
		return !fsStat.IsDir()
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		return false
	}
}

func ExistsDir(fileFullPath string) bool {
	if fsStat, err := os.Stat(fileFullPath); err == nil {
		// path/to/whatever exists
		return fsStat.IsDir()
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		return false
	}
}

func GetOSFile(inputFilepath string) (f os.FileInfo, err error) {
	osFile, openErr := os.Open(inputFilepath)
	if openErr != nil {
		err = openErr
		return
	}
	f, statErr := osFile.Stat()
	if statErr != nil {
		err = statErr
		return
	}
	return
}

func CopyFile(in, out string) (int64, error) {
	i, e := os.Open(in)
	if e != nil {
		return 0, e
	}
	defer i.Close()
	o, e := os.Create(out)
	if e != nil {
		return 0, e
	}
	defer o.Close()
	return o.ReadFrom(i)
}

func CopyIoFile(src, dest string) error {
	buf := make([]byte, 1024)
	fin, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return err
	}
	defer fin.Close()

	fOut, err := os.Create(dest)
	if err != nil {
		log.Println(err)
		return err
	}
	defer fOut.Close()

	n, err := io.CopyBuffer(fOut, fin, buf)
	log.Println("write", n)
	if err != nil {
		log.Println("io copy has err: ", err)
		return err
	}
	return nil
}
