package main

import (
	"context"
	"fmt"
	"github.com/qkgo/scaff/pkg/util/buffers"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"
)

func main() {
	//src := "words.txt"
	fileName := "file_create_by_cancelable_copy_%d_%d.block"

	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:7777", nil))
	}()

	argsWithoutProgram := os.Args[1:]

	maxThreads, err := strconv.Atoi(argsWithoutProgram[0])
	if err != nil {
		maxThreads = 10
	}

	maxDownloadIdleTime, err := strconv.Atoi(argsWithoutProgram[1])
	if err != nil {
		maxDownloadIdleTime = 10
	}

	log.Printf("current test file set maxThreads to %d ; and max download time is %d   \n", maxThreads, maxDownloadIdleTime)

	var ctxList []context.Context

	for k := 0; ; k++ {
		log.Printf("loop count:  %d \n", k)
		for i := 0; i < maxThreads; i++ {
			log.Printf("starting download by thread num:  %d \n", i)
			go DownloadFileAutoIncrease(ctxList, fileName, k, i)
			time.Sleep(time.Millisecond)
		}

		for i := maxDownloadIdleTime; i > 0; i-- {
			log.Printf("stop context after %d second \n", i)
			time.Sleep(time.Second)
		}

		for i := 0; i < len(ctxList); i++ {
			log.Printf("starting stop context %d \n", i)
			(ctxList)[i].Done()
			time.Sleep(time.Millisecond)
		}

		for i := maxDownloadIdleTime; i > 0; i-- {
			log.Printf("start context after %d second \n", i)
			time.Sleep(time.Second)
		}
	}
}

var remoteUrl = "https://tempo-prod-cn-shanghai-dogfish.oss-cn-shanghai.aliyuncs.com/1970/1/1/0/GS-0038-0001-2411-0001/42a9d562-c4e4-4414-8ca0-c1e2e0c2c80b/2022-3-12-9-17-49_5.bag?Expires=1669142771&OSSAccessKeyId=TMP.3Kg2f8TnKrkNSbw8RPjHKNPyQjKjhHMjfb8Myc3wjLA2waezMNXnwUL4YCVQ88carqY8EbUq47Cm5eHeePJJ1pwCucuo7n&Signature=dLBUaxlAZoolk4eNmmRnPq9XP8Y%3D"

func DownloadFileAutoIncrease(ctxList []context.Context, fileName string, k int, i int) {
	func(loopCount, currentNum int) {
		ctx := context.Background()
		ctxList = append(ctxList, ctx)

		tr := &http.Transport{} // TODO: copy defaults from http.DefaultTransport

		req, networkErr := http.NewRequestWithContext(ctx, "GET", remoteUrl, nil)
		if networkErr != nil {
			log.Println(networkErr)
			return
		}
		client := http.Client{
			Transport: tr,
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		// Put content on file

		resp, networkErr := client.Do(req)
		if networkErr != nil {
			log.Println(networkErr)
			return
		}
		defer resp.Body.Close()

		currentFileName := fmt.Sprintf(fileName, loopCount, currentNum)
		// Create blank file
		file, createFileErr := os.Create(currentFileName)
		if createFileErr != nil {
			log.Println(createFileErr)
			return
		}
		size, copyBufferError := buffers.CancelableCopy(ctx, file, resp.Body, func() {
			log.Printf("stop stream at %d : %d \n", loopCount, currentNum)
			srcCloseErr := file.Close()
			dstCloseErr := resp.Body.Close()
			if srcCloseErr != nil {
				log.Panicln(srcCloseErr)
			}
			if dstCloseErr != nil {
				log.Panicln(dstCloseErr)
			}
		})

		defer file.Close()
		if copyBufferError != nil {
			log.Println(copyBufferError)
			return
		}
		fmt.Printf("Downloaded a file %s with size %d", currentFileName, size)
	}(k, i)
}
