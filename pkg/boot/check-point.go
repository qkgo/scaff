package boot

import (
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/intercepter"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var countHealthCheck atomic.Value

var checkMap sync.Map

type NeedValidationObject struct {
	Search bool
	Change bool
	Remove bool
	Input  bool
}

func calcServerHealthWrapper(url string, runChannel chan int, doValidationAPIList *[]string) bool {
	if url == "" {
		return false
	}
	if strings.Contains(url, "RoleInfo:Remove") {
		cfg.Log.Info(url)
	}
	result := CheckServerHealthStore(url, doValidationAPIList)
	if result {
		cfg.CheckServerStateMap[url] = true
	}
	if countHealthCheck.Load().(int) >= len(*doValidationAPIList)-1 {
		cfg.Log.Info("ok:", countHealthCheck.Load(), " - ", len(*doValidationAPIList), " ->", url, " : ", result)
		runChannel <- countHealthCheck.Load().(int)
	}
	countHealthCheck.Store(countHealthCheck.Load().(int) + 1)
	cfg.Log.Info(countHealthCheck.Load(), " - ", len(*doValidationAPIList), " ->", url, " : ", result)
	return result
}

func CheckServerHealthStore(url string, doValidationAPIList *[]string) bool {
	if url == "" {
		return false
	}
	if intercepter.CheckServerHost == "" {
		return false
	}
	if strings.Contains(url, "RoleInfo:Remove") {
		cfg.Log.Info(url)
	}
	checkMap.Store(url, true)
	resp, err := resty.R().Options(intercepter.CheckServerHost + url + "/check")
	checkMap.Delete(url)
	if err != nil {
		cfg.Log.Info(err)
		return false
	}
	if resp.StatusCode() < 300 {
		return resp.StatusCode() < 300
	}
	return false
}

func CheckServerList(doValidationAPIList *[]string, runChannel chan int) {
	countHealthCheck.Store(0)
	for _, ApiUrl := range *doValidationAPIList {
		cfg.Log.Info(ApiUrl + ": resource has verify api")
		if ApiUrl != "" {
			go calcServerHealthWrapper(ApiUrl, runChannel, doValidationAPIList)
		}
	}
}

func PreRequestCheckPointInitialization(permissionCodeList *[]string) {
	runnerChannel := make(chan int, 1)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 5)
		timeout <- true
	}()

	validationHost := cfg.OzConfig.GetValidationHost()
	cfg.Log.Info("validation host:", validationHost)
	intercepter.SetCheckServerHost(validationHost)
	go CheckServerList(permissionCodeList, runnerChannel)
	select {
	case <-runnerChannel:
		cfg.Log.Info("finish :", runnerChannel)
		return
	case <-timeout:
		cfg.Log.Info("time out , ch > 0")
		id := 0
		Aux := func(key interface{}, value interface{}) bool {
			cfg.Log.Info("error Kv:", "value:", key.(string), value.(bool))
			id++
			return true
		}
		checkMap.Range(Aux)
		if id > 0 {
			cfg.Log.Info("timeout. ch > 0 ")
		} else {
			cfg.Log.Info("timeout. ch < 1")
		}
		return
	}
}

func PostRequestHookPointInitialization() {

}
