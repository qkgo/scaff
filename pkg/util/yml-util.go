package util

import (
	config "github.com/go-ozzo/ozzo-config"
	"github.com/qkgo/scaff/pkg/cfg"
)

func GetItfFromYml(path string) interface{} {
	var ConfigParam = config.New()
	err := ConfigParam.Load(path)
	if err != nil {
		cfg.Log.Info(err)
	}
	return ConfigParam.Data()
}


func GetConfigStringList(c *config.Config, key string) []string {
	var hasIgnoreMatchFieldList []interface{}
	ignoreMatchFieldAssert := c.Get(key)
	if ignoreMatchFieldAssert != nil {
		hasIgnoreMatchFieldList = ignoreMatchFieldAssert.([]interface{})
	}
	var ignoreMatchFieldList []string
	for _, a := range hasIgnoreMatchFieldList {
		ignoreMatchFieldList = append(ignoreMatchFieldList, a.(string))
	}
	return ignoreMatchFieldList
}
