package main

import (
	"context"
	"fmt"
	"github.com/qkgo/scaff/pkg/util/buffers"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"
)

const keyServerAddr = "serverAddr"

var remoteUrl = "https://consumer-tkb.huawei.com/weknow/servlet/download/public?contextNo=W00036307"
var maxThreads = 10
var maxDownloadIdleTime = 2

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my website!\n")
}

func main() {
	//src := "words.txt"
	fileName := "file_create_by_cancelable_copy_%d_%d.block"

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", getRoot)
		//mux.HandleFunc("/hello", getHello)

		log.Println(http.ListenAndServe("0.0.0.0:7777", mux))
	}()

	var err error
	var argsWithoutProgram []string
	if len(os.Args) > 1 {
		argsWithoutProgram = os.Args[1:]
	}
	if len(argsWithoutProgram) > 1 {
		maxThreads, err = strconv.Atoi(argsWithoutProgram[0])
		if err != nil {
			maxThreads = 10
			err = nil
		}
	}

	if len(argsWithoutProgram) > 2 {
		maxDownloadIdleTime, err = strconv.Atoi(argsWithoutProgram[1])
		if err != nil {
			maxDownloadIdleTime = 10
			err = nil
		}
	}

	if len(argsWithoutProgram) > 3 {
		remoteUrl = argsWithoutProgram[2]
	}

	log.Printf("current test file set maxThreads to %d ; and max download time is %d   \n", maxThreads, maxDownloadIdleTime)

	var ctxList []context.Context
	var cancelList []context.CancelFunc

	for k := 0; k < 1; k++ {
		log.Printf("loop count:  %d \n", k)
		for i := 0; i < maxThreads; i++ {
			log.Printf("starting download by thread num:  %d \n", i)
			go DownloadFileAutoIncrease(&ctxList, &cancelList, fileName, k, i)
			time.Sleep(time.Millisecond)
		}

		for i := maxDownloadIdleTime; i > 0; i-- {
			log.Printf("stop context after %d second \n", i)
			time.Sleep(time.Second)
		}

		for i := 0; i < len(cancelList); i++ {
			log.Printf("starting stop context %d \n", i)
			// abort http request only for:  ctx.cancel
			(cancelList)[i]()
			//(ctxList)[i].Done()  // abort http request only for:  ctx.cancel
			time.Sleep(time.Millisecond)
		}

		for i := maxDownloadIdleTime + 2; i > 0; i-- {
			log.Printf("start context after %d second \n", i)
			time.Sleep(time.Second)
		}
	}
	select {}
}

func DownloadFileAutoIncrease(ctxList *[]context.Context, cancelList *[]context.CancelFunc, fileName string, loopCount, currentNum int) {
	//func(loopCount, currentNum int) {
	ctx, cancel := context.WithCancel(context.Background())
	*ctxList = append(*ctxList, ctx)
	*cancelList = append(*cancelList, cancel)

	//tr := &http.Transport{} // TODO: copy defaults from http.DefaultTransport
	//tr.CancelRequest()
	req, networkErr := http.NewRequestWithContext(ctx, "GET", remoteUrl, nil)
	if networkErr != nil {
		log.Println("NewRequestWithContext:", networkErr)
		return
	}
	client := http.Client{
		//Transport: tr,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file

	resp, networkErr := client.Do(req)
	if networkErr != nil {
		log.Println("do.networkErr:", networkErr)
		return
	}
	defer resp.Body.Close()
	currentFileName := fmt.Sprintf(fileName, loopCount, currentNum)
	// Create blank file
	file, createFileErr := os.Create(currentFileName)
	defer file.Close()
	if createFileErr != nil {
		log.Println("createFileErr:", createFileErr)
		return
	}
	size, copyBufferError := buffers.CancelableCopy(ctx, file, resp.Body, func() {
		log.Printf("stop stream at %d : %d \n", loopCount, currentNum)
		dstCloseErr := resp.Body.Close()
		if dstCloseErr != nil {
			log.Printf("close %s dst err: %s \n", currentFileName, dstCloseErr)
		}
		srcCloseErr := file.Close()
		if srcCloseErr != nil {
			log.Printf("close %s src err: %s \n", currentFileName, srcCloseErr)
		}
	})
	//size, copyBufferError := io.Copy(file, resp.Body)

	if copyBufferError != nil {
		log.Println("copyBufferError:", copyBufferError)
		return
	}
	fmt.Printf("Downloaded a file %s with size %d  \n ", currentFileName, size)
	//}(k, i)
}
