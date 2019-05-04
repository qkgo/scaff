package util

import (
	"github.com/chain-zhang/pinyin"
	"regexp"
	"strconv"
	"time"
)

func GetCode(chinese string) string {
	str, err := pinyin.New(chinese).Split("").Mode(pinyin.InitialsInCapitals).Convert()
	if err != nil {
	} else {
	}
	return str
}

func mergeStringList(stringList []string) string {
	chString := ""
	for _, value := range stringList {
		chString += value
	}
	return chString
}

func mergeChinaStringList(stringList []string) string {
	chString := ""
	for _, value := range stringList {
		chString += GetCode(value)
	}
	return chString
}

func StringToCode(text string, maxLength int) string {
	if maxLength < 20 {
		maxLength = 20
	}
	reg := regexp.MustCompile("[\u4e00-\u9fa5]+")
	chString := mergeChinaStringList(reg.FindAllString(text, -1))
	reg2 := regexp.MustCompile("[A-Za-z]+")
	chString1 := mergeStringList(reg2.FindAllString(text, -1))
	reg3 := regexp.MustCompile("[0-9]+")
	chString2 := mergeStringList(reg3.FindAllString(text, -1))
	if maxLength >= len(text) {
		return chString + chString1 + chString2
	} else {
		return (chString + chString1 + chString2)[:maxLength-1]
	}
}

func StringToDateCode(text string, maxLength int) string {
	if maxLength < 20 {
		maxLength = 20
	}
	reg := regexp.MustCompile("[\u4e00-\u9fa5]+")
	chString := mergeChinaStringList(reg.FindAllString(text, -1))
	reg2 := regexp.MustCompile("[A-Za-z]+")
	chString1 := mergeStringList(reg2.FindAllString(text, -1))
	reg3 := regexp.MustCompile("[0-9]+")
	chString2 := mergeStringList(reg3.FindAllString(text, -1))
	now := time.Now()
	timeString := now.Month().String() + strconv.Itoa(now.Day())
	if maxLength >= len(text) {
		return (chString + chString1 + chString2) + timeString
	} else {
		return (chString + chString1 + chString2)[:maxLength-1-len(timeString)] + timeString
	}

}

func StringToUnixCode(text string, maxLength int) string {
	if maxLength < 20 {
		maxLength = 20
	}
	if maxLength > len(text) {
		maxLength = len(text)
	}
	reg := regexp.MustCompile("[\u4e00-\u9fa5]+")
	chString := mergeChinaStringList(reg.FindAllString(text, -1))
	reg2 := regexp.MustCompile("[A-Za-z]+")
	chString1 := mergeStringList(reg2.FindAllString(text, -1))
	reg3 := regexp.MustCompile("[0-9]+")
	chString2 := mergeStringList(reg3.FindAllString(text, -1))
	now := time.Now()
	timeString := strconv.FormatInt(now.UnixNano(), 10)
	if maxLength >= len(text) {
		return (chString + chString1 + chString2) + timeString
	} else {
		return (chString + chString1 + chString2)[:maxLength-1-len(timeString)] + timeString
	}
}
