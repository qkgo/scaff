package cfg

import (
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func ReadYamlFile(filename string) map[string]interface{} {
	mapData := make(map[string]interface{})

	fileDataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		Log.Warn("read file error: ", err)
		os.Exit(-10)
		return nil
	}
	err = yaml.Unmarshal([]byte(fileDataByte), &mapData)
	if err != nil {
		log.Fatalf("read yml error: %v", err)
		os.Exit(-12)
		return nil
	}
	log.Printf("--- m:\n%v\n\n", mapData)
	return mapData
}
