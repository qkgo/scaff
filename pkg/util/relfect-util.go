package util

import (
	"encoding/json"
	"reflect"
)

func GetString(str interface{}) string {
	if reflect.TypeOf(str).String() == "null.String" {
		return str.(null.String).String
	} else if reflect.TypeOf(str).String() == "string" {
		return str.(string)
	} else {
		return ""
	}
}

func CompareString(v1 interface{}, v2 interface{}) bool {
	return GetString(v1) == GetString(v2)
}

func MergeProperties(originInterface interface{}, propertiesInterface ...interface{}) interface{} {
	JsonByte, _ := json.Marshal(originInterface)
	var collection map[string]interface{}
	json.Unmarshal(JsonByte, &collection)
	for _, data := range propertiesInterface {
		collection[reflect.TypeOf(data).String()] = data
	}
	return collection
}
