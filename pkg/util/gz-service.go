package util

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"github.com/djherbis/buffer"
	"github.com/qkgo/aliyun-oss-go-sdk/oss"
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/smallnest/ringbuffer"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type UploadFileInformationRes struct {
	OssBucketName  string
	OssStorageFile string
	FileMd5        string
	FileSize       int64
	Host           string
	FileName       string
}

type UploadFileInformation struct {
	File           *multipart.FileHeader
	OssBucketName  string
	OssPackageFile string
	FileMd5        string
	FileSize       int64
	OssContentFile string
	NotUseMemory   bool
	FileList       []string
	CheckFile      func(io.Reader) error
	SumMd5         bool
	MaxKeys        int64
}

func GetMemoryWriter() io.Writer {
	rb := ringbuffer.New(9999999)
	return rb
}

func GetMemoryWriter2() io.Writer {
	rb := buffer.New(999999999)
	return rb
}

func GetFileWriter(fileInformation *UploadFileInformation) io.Writer {
	memoryTarWriter, err := os.Create(path.Join(GetCurrentPathWithoutError(), fileInformation.OssPackageFile))
	if err != nil {
		cfg.LogInfo.Info(err)
	}
	return memoryTarWriter
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReplaceTarGz(OssSavedObject *io.Reader, fileInformation *UploadFileInformation) (io.Writer, error) {
	gzipReader, err := gzip.NewReader(*OssSavedObject)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	var memoryTarWriter io.Writer
	if fileInformation.NotUseMemory {
		memoryTarWriter = GetFileWriter(fileInformation)
	} else {
		memoryTarWriter = new(bytes.Buffer)
	}
	gw := gzip.NewWriter(memoryTarWriter)
	defer gw.Close()
	tarWriter := tar.NewWriter(gw)
	defer tarWriter.Close()
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				cfg.LogInfo.Info(err)
				break
			}
		}
		if strings.Contains(header.Name, fileInformation.OssContentFile) {
			hdr := new(tar.Header)
			hdr.Name = header.Name
			hdr.Size = fileInformation.File.Size
			hdr.ModTime = time.Now()
			hdr.AccessTime = time.Now()
			hdr.ChangeTime = time.Now()
			hdr.Mode = 420
			hdr.Typeflag = 48
			err = tarWriter.WriteHeader(hdr)
			if err != nil {
				cfg.LogInfo.Info(err)
				break
			}
			file, err := fileInformation.File.Open()
			if err != nil {
				cfg.LogInfo.Info(err)
				break
			}
			data, err := ioutil.ReadAll(file)
			if err != nil {
				cfg.LogInfo.Info(err)
				break
			}
			ii, err := tarWriter.Write(data)
			if err != nil {
				cfg.LogInfo.Info(err, ii)
				break
			}
		} else {
			err = tarWriter.WriteHeader(header)
			if err != nil {
				cfg.LogInfo.Info(err)
				break
			}
			byteData, err := ioutil.ReadAll(tarReader)
			if err != nil {
				cfg.LogInfo.Info(err)
				break
			}
			ii, err := tarWriter.Write(byteData)
			if err != nil {
				cfg.LogInfo.Info(err, ii)
				break
			}
		}
	}
	if !fileInformation.NotUseMemory {
		tarWriter.Close()
		gw.Close()
		return memoryTarWriter, nil
	}

	return nil, errors.New("gz not implemented")
}

type FileInfo struct {
	Name     string
	Size     int64
	ModTime  time.Time
	TypeFlag byte
	Mode     int64
}

func FileNameListInTarGzContent(OssSavedObject *io.Reader) ([]*FileInfo, error) {
	gzipReader, err := gzip.NewReader(*OssSavedObject)
	if err != nil {
		cfg.LogInfo.Info(err)
		return nil, err
	}
	defer gzipReader.Close()
	tarReader := tar.NewReader(gzipReader)
	var fileList []*FileInfo
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				cfg.LogInfo.Info(err)
				break
			}
		}
		file := &FileInfo{
			Name:     header.Name,
			Size:     header.Size,
			ModTime:  header.ModTime,
			TypeFlag: header.Typeflag,
			Mode:     header.Mode,
		}
		fileList = append(fileList, file)
	}
	return fileList, nil
}

func FileNameListInTarGzContentByte(buf io.Reader) ([]*FileInfo, error) {
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		cfg.LogInfo.Info(err)
		return nil, err
	}
	defer gzipReader.Close()
	tarReader := tar.NewReader(gzipReader)
	var fileList []*FileInfo
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				cfg.LogInfo.Info(err)
				break
			}
		}
		file := &FileInfo{
			Name:     header.Name,
			Size:     header.Size,
			ModTime:  header.ModTime,
			TypeFlag: header.Typeflag,
			Mode:     header.Mode,
		}
		fileList = append(fileList, file)
	}
	return fileList, nil
}

func writeFile3(byte bytes.Buffer, fileInformation *UploadFileInformation) {
	c, _ := ioutil.ReadAll(&byte)
	if ioutil.WriteFile(path.Join(GetCurrentPathWithoutError(), fileInformation.OssPackageFile), c, os.ModePerm) == nil {
		cfg.LogInfo.Info("ok")
	}
}

func writeFile2(memoryTarWriter io.Writer, fileInformation *UploadFileInformation) {
	var b = memoryTarWriter.(io.ReadWriter)
	c, _ := ioutil.ReadAll(b)
	if ioutil.WriteFile(path.Join(GetCurrentPathWithoutError(), fileInformation.OssPackageFile), c, os.ModePerm) == nil {
		cfg.LogInfo.Info("ok")
	}
}

func writeFile(memoryTarWriter io.Writer, fileInformation *UploadFileInformation) {
	var b = memoryTarWriter.(io.ReadWriter)
	c, _ := ioutil.ReadAll(b)
	f, err := os.Create(path.Join(GetCurrentPathWithoutError(), fileInformation.OssPackageFile))
	check(err)
	defer f.Close()
	n2, err := f.Write(c)
	cfg.LogInfo.Info(n2)
}

func UploadToOss(ossConnection *oss.Client, ossBucket *oss.Bucket, writer io.Writer, fileUrl string) (*oss.Response, error) {
	resp, err := ossBucket.PutObjectV2(fileUrl, writer.(io.Reader), nil)
	if err != nil {
		return nil, err
	} else {
		cfg.LogInfo.Info(resp)
		return resp, nil
	}
	//err := ossBucket.PutObject(fileUrl, writer.(io.Reader), nil)
	//return nil, err
}

func ReParkFile(OssSavedObject *io.Reader, fileInformation *UploadFileInformation) (io.Writer, error) {
	gzipReader, err := gzip.NewReader(*OssSavedObject)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()
	tarReader := tar.NewReader(gzipReader)
	var memoryTarWriter io.Writer
	if fileInformation.NotUseMemory {
		memoryTarWriter = GetFileWriter(fileInformation)
	} else {
		memoryTarWriter = new(bytes.Buffer)
	}
	gw := gzip.NewWriter(memoryTarWriter)
	defer gw.Close()
	tarWriter := tar.NewWriter(gw)
	defer tarWriter.Close()
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				cfg.LogInfo.Info(err)
				break
			}
		}
		for _, itemPath := range fileInformation.FileList {
			if header.Name == itemPath {
				err = tarWriter.WriteHeader(header)
				if err != nil {
					cfg.LogInfo.Info(err)
					break
				}
				byteData, err := ioutil.ReadAll(tarReader)
				if err != nil {
					cfg.LogInfo.Info(err)
					break
				}
				ii, err := tarWriter.Write(byteData)
				if err != nil {
					cfg.LogInfo.Info(err, ii)
					break
				}
			}
		}
	}
	if !fileInformation.NotUseMemory {
		tarWriter.Close()
		gw.Close()
		return memoryTarWriter, nil
	}

	return nil, errors.New("FileList file not found")
}

func GetFileFromMultipart(OssSavedObject *io.Reader, fileInformation *UploadFileInformation) (io.Reader, error) {
	gzipReader, err := gzip.NewReader(*OssSavedObject)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				cfg.LogInfo.Info(err)
				break
			}
		}
		if strings.Contains(header.Name, fileInformation.OssContentFile) {
			return tarReader, nil
		}
	}
	return nil, errors.New("OssContentFile file not found")
}
