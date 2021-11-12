package util

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/util/system"
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

<<<<<<< Updated upstream
func MustJSONDecode(b []byte, i interface{}) {
	err := json.Unmarshal(b, i)
	if err != nil {
		log.Printf("decode json error: %v \n", err)
	}
}

func Get(inputBytes []byte, expectObject interface{}) interface{} {
	objType := reflect.TypeOf(expectObject).Elem()
	resultObject := reflect.New(objType).Interface()
	MustJSONDecode(inputBytes, &resultObject)
	if system.GO111MODULE() {
		log.Printf("json unmarshal resultObject = %#v \n", resultObject)
	}
	return resultObject
=======
func MustJSONDecode(b []byte, i interface{}) (err error) {
	err = json.Unmarshal(b, i)
	if err != nil {
		log.Printf("decode json error: %v \n", err)
	}
	return
}

func Get(inputBytes []byte, inputType interface{}) (i interface{}, err error) {
	typed := reflect2.TypeOf(inputType)
	newObject := typed.New()
	err = MustJSONDecode(inputBytes, &newObject)
	if system.GO111MODULE() {
		log.Printf("json unmarshal resultObject = %+v \n", newObject)
	}
	return newObject, err
}

func GetWithoutError(inputBytes []byte, inputType interface{}) (i interface{}, err error) {
	typed := reflect2.TypeOf(inputType)
	newObject := typed.New()
	err = MustJSONDecode(inputBytes, &newObject)
	if system.GO111MODULE() {
		log.Printf("json unmarshal resultObject = %+v \n", newObject)
	}
	return newObject, err
>>>>>>> Stashed changes
}
