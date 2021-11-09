package util

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/qkgo/scaff/pkg/cfg"
	"log"
	"reflect"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonParse(input interface{}) []byte {
	var jsonByte []byte
	defer func() []byte {
		if err := recover(); err != nil {
			log.Println(err)
			return []byte(fmt.Sprintf("%#v", input))
		}
		return jsonByte
	}()
	switch reflect.TypeOf(input) {
	case reflect.TypeOf([]string{}):
		jsonByte, err := json.Marshal(input)
		if err != nil {
			cfg.Log.Info("json parse:", err)
		}
		return jsonByte
	case reflect.TypeOf(map[string]string{}):
		jsonByte, err := json.Marshal(input)
		if err != nil {
			cfg.Log.Info("json parse:", err)
		}
		return jsonByte
	default:
		return nil
	}
}

func JsonQuickParse(input interface{}) []byte {
	var jsonByte []byte
	defer func() []byte {
		if err := recover(); err != nil {
			log.Printf("json parse error: ", err)
			return []byte(fmt.Sprintf("%#v", input))
		}
		return jsonByte
	}()
	jsonByte, err := json.Marshal(input)
	if err != nil {
		log.Printf("json parse error: ", err.Error())
		if jsonByte == nil {
			return []byte(fmt.Sprintf("%#v", input))
		}
	}
	return jsonByte
}
