package cfg

import (
	"github.com/qkgo/scaff/pkg/util/system"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func ReadYamlFile(filename string) map[string]interface{} {
	mapData := make(map[string]interface{})
	fileDataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		Log.Warn("read file error: ", err)
		system.ExIt(-10)
		return nil
	}
	err = yaml.Unmarshal([]byte(fileDataByte), &mapData)
	if err != nil {
		log.Fatalf("read yml error: %v", err)
		system.ExIt(-12)
		return nil
	}
	log.Printf("--- m:\n%v\n\n", mapData)
	return mapData
}
